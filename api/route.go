package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var routes []route

type route struct {
	Method  string
	Pattern string
	Handler gin.HandlerFunc
}

func (r route) register() {
	routes = append(routes, r)
}

func init() {
	route{
		Method:  http.MethodGet,
		Pattern: "",
		Handler: func(context *gin.Context) {
			context.String(http.StatusOK, "Hello from Go :3")
		},
	}.register()
}

// RegisterRoutes registers API routes to a router group.
func RegisterRoutes(router *gin.RouterGroup) {
	for _, r := range routes {
		switch r.Method {
		case http.MethodGet:
			router.GET(r.Pattern, r.Handler)
		case http.MethodPost:
			router.POST(r.Pattern, r.Handler)
		case http.MethodPut:
			router.PUT(r.Pattern, r.Handler)
		case http.MethodPatch:
			router.PATCH(r.Pattern, r.Handler)
		case http.MethodDelete:
			router.DELETE(r.Pattern, r.Handler)
		}
	}
}

func placeholder(context *gin.Context) {
	context.String(http.StatusNotImplemented, "TODO")
}
