package openapi

import (
	"context"
	"database/sql"
	"github.com/friendsofgo/errors"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"io"
	"log"
	"net/http"
	"random.chars.jp/git/misskey/api/payload"
	"random.chars.jp/git/misskey/api/structs"
	"random.chars.jp/git/misskey/data"
	"random.chars.jp/git/misskey/db/models"
	"strings"
	"time"
)

var authenticationFailure = data.New()

func init() {
	authenticationFailure.Set(structs.APIError{
		Message: "Authentication failed. Please ensure your token is correct.",
		Code:    "AUTHENTICATION_FAILED",
		ID:      "b0a7f5f8-dc2f-4171-b91f-de88ad238e14",
		Kind:    "client",
	})
}

/* the only reason this is the way it is, is due to authentication data being sent as part of the request body alongside
all the parameters, which makes it almost impossible to have well-defined data structures and handle them reasonably
transparently without wasting a parse one way or another, and since upstream basically accesses the fields on-demand
and the models in documentation are mostly bullshit, this is actually not so bad

if you want to get rid of this ugliness, the only way is to go through syuilo, either make him accept my API design or
make him learn and do proper API design, and from what I've seen, that's not going to happen any time soon, and I bet
it won't happen before Go gets some form of generics, and maybe that'll make this code less ugly

until that happens, or someone figures out a better way of doing this, this is the way it will be

good luck */

// UglyMap is the intermediate data structure for response body handling
type UglyMap map[string]json.RawMessage

// Context carries API call context information
// The first call to any method that accesses the response body
// (Bind, User, etc.) is not safe for concurrent use
type Context interface {
	// Request returns pointer to the API call's http.Request
	Request() *http.Request
	// Writer returns response writer of the API call
	Writer() http.ResponseWriter

	// FullPath returns full path of matched route or empty string if not found
	FullPath() string
	// GetHeader gets the value of a header
	GetHeader(key string) string
	// BindHeader binds the headers to a struct
	BindHeader(obj interface{}) error
	// JSON sends marshalled JSON payload from obj
	JSON(code int, obj interface{})
	// RawJSON sends bytes with JSON mime type
	RawJSON(code int, data []byte)
	// String sends string as plaintext
	String(code int, str string)
	// BadRequest sends bad request and aborts
	BadRequest()
	// InternalServerError sends zero value of structs.APIError and aborts
	InternalServerError()

	// UglyMap returns intermediate ugly map for response body parsing
	UglyMap() (UglyMap, error)
	// BindKey binds to key of response body
	BindKey(key string, obj interface{}) (bool, error)
	// Authenticate authenticates using response body
	// the boolean is true when key does not exist or authentication is not successful
	// and false when key exists but does not match
	Authenticate() (*models.User, *models.AccessToken, bool, error)
	// MustAuthenticate ensures authentication for required credential endpoints
	MustAuthenticate() (*models.User, *models.AccessToken, bool, error)
	// RequireCredential wraps a function that requires credential
	RequireCredential(handler func(ctx Context, user *models.User))

	// Abort prevents pending middleware from being called
	Abort()
	context.Context
}

type ginContext struct {
	intermediate UglyMap

	route route
	*gin.Context
}

func newWrap(r route) gin.HandlerFunc {
	return func(context *gin.Context) {
		r.Handler(&ginContext{route: r, Context: context})
	}
}

func (ctx *ginContext) Request() *http.Request {
	return ctx.Context.Request
}

func (ctx *ginContext) Writer() http.ResponseWriter {
	return ctx.Context.Writer
}

func (ctx *ginContext) populateUglyMap() error {
	d := json.NewDecoder(ctx.Context.Request.Body)
	d.UseNumber()
	err := d.Decode(&ctx.intermediate)
	if err == io.EOF {
		ctx.intermediate = make(UglyMap)
		return nil
	}
	return err
}

func (ctx *ginContext) RawJSON(code int, data []byte) {
	ctx.Data(code, gin.MIMEJSON, data)
}

func (ctx *ginContext) String(code int, str string) {
	ctx.Data(code, gin.MIMEPlain, []byte(str))
}

func (ctx *ginContext) BadRequest() {
	ctx.String(http.StatusBadRequest, "Bad Request")
	ctx.Abort()
}

func (ctx *ginContext) InternalServerError() {
	ctx.RawJSON(http.StatusInternalServerError, payload.InternalServerError.Data())
	ctx.Abort()
}

// ensureUglyMap ensures the ugly map is populated
// NOT safe for concurrent use on first access
func (ctx *ginContext) ensureUglyMap() error {
	if ctx.intermediate == nil {
		return ctx.populateUglyMap()
	}

	return nil
}

func (ctx *ginContext) UglyMap() (UglyMap, error) {
	if err := ctx.ensureUglyMap(); err != nil {
		return nil, err
	}

	return ctx.intermediate, nil
}

func (ctx *ginContext) BindKey(key string, obj interface{}) (bool, error) {
	if err := ctx.ensureUglyMap(); err != nil {
		return false, err
	}

	if msg, ok := ctx.intermediate[key]; !ok {
		return false, nil
	} else {
		return true, json.Unmarshal(msg, obj)
	}
}

func (ctx *ginContext) Authenticate() (*models.User, *models.AccessToken, bool, error) {
	var token string
	if ok, err := ctx.BindKey("i", &token); err != nil || !ok {
		return nil, nil, true, err
	}

	// upstream has is-native-token.ts with a function that checks token length equals to 16
	if len(token) == 16 {
		if user, err := models.Users(qm.Where("token = ?", token)).OneG(ctx); err != nil {
			if err == sql.ErrNoRows {
				ctx.RawJSON(http.StatusForbidden, authenticationFailure.Data())
				ctx.Abort()
				return nil, nil, false, nil
			}
			return nil, nil, false, err
		} else {
			return user, nil, true, nil
		}
	} else {
		if accessToken, err := models.AccessTokens(
			qm.Where("(hash = ?) or (token = ?)", strings.ToLower(token), token),
		).OneG(ctx); err != nil {
			if err == sql.ErrNoRows {
				ctx.RawJSON(http.StatusForbidden, authenticationFailure.Data())
				ctx.Abort()
				return nil, nil, false, nil
			}
			return nil, nil, false, err
		} else {
			accessToken.LastUsedAt = null.TimeFrom(time.Now())
			if _, err = accessToken.UpdateG(ctx, boil.Infer()); err != nil {
				return nil, nil, false, err
			}

			var (
				user *models.User
				app  *models.App
			)
			if user, err = accessToken.UserIdUser().OneG(ctx); err != nil {
				return nil, nil, false, err
			}

			if accessToken.AppId.Valid {
				if app, err = accessToken.AppIdApp().OneG(ctx); err != nil {
					return nil, nil, false, err
				} else if app == nil {
					return nil, nil, false, errors.New("app not found")
				}
				return user, &models.AccessToken{ID: accessToken.ID, Permission: app.Permission}, true, nil
			} else {
				return user, accessToken, true, nil
			}
		}
	}
}

func (ctx *ginContext) MustAuthenticate() (*models.User, *models.AccessToken, bool, error) {
	// see api/NOTES for why this is done the way it is done
	user, accessToken, ok, err := ctx.Authenticate()
	if ok && err == nil && user == nil {
		ctx.RawJSON(http.StatusUnauthorized, payload.CredentialRequired.Data())
		ctx.Abort()
	}
	return user, accessToken, ok, err
}

func (ctx *ginContext) RequireCredential(handler func(ctx Context, user *models.User)) {
	if user, _, ok, err := ctx.MustAuthenticate(); err != nil {
		log.Printf("error authenticating: %s", err)
		ctx.InternalServerError()
		return
	} else if !ok || user == nil {
		return
	} else {
		handler(ctx, user)
	}
}
