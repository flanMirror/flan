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

// UglyMap is the intermediate data structure for response body handling
type UglyMap map[string]json.RawMessage

// Context carries API call context information
// The first call to any method that accesses the response body
// (Bind, User, etc.) is not safe for concurrent use
type Context interface {
	Request() *http.Request
	Writer() http.ResponseWriter

	FullPath() string
	GetHeader(key string) string
	BindHeader(obj interface{}) error
	JSON(code int, obj interface{})
	RawJSON(code int, data []byte)
	String(code int, str string)
	BadRequest()
	InternalServerError()

	UglyMap() (UglyMap, error)
	BindKey(key string, obj interface{}) (bool, error)
	Authenticate() (*models.User, *models.App, bool, error)

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

func (ctx *ginContext) Authenticate() (*models.User, *models.App, bool, error) {
	var token string
	if ok, err := ctx.BindKey("i", &token); err != nil || !ok {
		return nil, nil, true, err
	}

	// upstream has is-native-token.ts with a function that checks token length equals to 16
	if len(token) == 16 {
		if user, err := models.Users(qm.Where("token = ?", token)).OneG(ctx); err != nil {
			if err == sql.ErrNoRows {
				ctx.RawJSON(http.StatusForbidden, authenticationFailure.Data())
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
				return user, &models.App{ID: accessToken.ID, Permission: app.Permission}, true, nil
			} else {
				return user, &models.App{
					ID:          accessToken.ID,
					CreatedAt:   accessToken.CreatedAt,
					UserId:      null.StringFrom(accessToken.UserId),
					Name:        accessToken.Name.String,
					Description: accessToken.Description.String,
					Permission:  accessToken.Permission,
				}, true, nil
			}
		}
	}
}
