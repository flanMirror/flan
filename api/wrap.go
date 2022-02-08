package openapi

import (
	"context"
	"github.com/friendsofgo/errors"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"io"
	"net/http"
	"reflect"
)

/* note: the Bind method in this does NOT use tags

instead, it flips the first character of the field name to lower case and indexes the ugly map
it also doesn't support more complex data structures, variables are directly typecasted,
this might be a problem in the future but for now I'll not worry about it

this is the ugliest piece of code in this entire codebase so far */

// Context carries API call context information
// The first call to any method that accesses the response
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

	UglyMap() (map[string]interface{}, error)
	Bind(obj interface{}) error
	// TODO: auth method

	Abort()
	context.Context
}

type ginContext struct {
	intermediate map[string]interface{}
	// TODO: store auth info

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
		ctx.intermediate = make(map[string]interface{})
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

// ensureUglyMap ensures the ugly map is populated
// NOT safe for concurrent use on first access
func (ctx *ginContext) ensureUglyMap() error {
	if ctx.intermediate == nil {
		return ctx.populateUglyMap()
	}

	return nil
}

func (ctx *ginContext) UglyMap() (map[string]interface{}, error) {
	if err := ctx.ensureUglyMap(); err != nil {
		return nil, err
	}

	return ctx.intermediate, nil
}

func (ctx *ginContext) Bind(obj interface{}) error {
	if err := ctx.ensureUglyMap(); err != nil {
		return err
	}
	return bindMap(obj, ctx.intermediate)
}

// bindMap accepts a struct and unmarshal intermediate map into it
func bindMap(obj interface{}, intermediate map[string]interface{}) error {
	v := reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		name := []byte(t.Field(i).Name)
		name[0] ^= 0x20

		if data, ok := intermediate[string(name)]; ok {
			field := v.Field(i)
			if !field.IsValid() || !field.CanSet() {
				return errors.New("cannot set field " + string(name))
			}
			value := reflect.ValueOf(data)
			if field.Type() == value.Type() {
				field.Set(value)
			} else if reflect.PtrTo(value.Type()) == field.Type() {
				ptr := reflect.New(value.Type())
				ptr.Elem().Set(value)
				field.Set(ptr)
			} else {
				return errors.New("type difference on field " + string(name) +
					", field " + field.Type().String() +
					", value " + value.Type().String())
			}
		}
	}
	return nil
}
