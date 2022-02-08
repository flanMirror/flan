package openapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var routes []route

type handlerFunc func(context Context)
type route struct {
	Method  string
	Pattern string
	Handler handlerFunc
}

func (r route) register() {
	routes = append(routes, r)
}

// RegisterRoutes registers API routes to a router group.
func RegisterRoutes(router *gin.RouterGroup) {
	for _, r := range routes {
		switch r.Method {
		case http.MethodGet:
			router.GET(r.Pattern, newWrap(r))
		case http.MethodPost:
			router.POST(r.Pattern, newWrap(r))
		case http.MethodPut:
			router.PUT(r.Pattern, newWrap(r))
		case http.MethodPatch:
			router.PATCH(r.Pattern, newWrap(r))
		case http.MethodDelete:
			router.DELETE(r.Pattern, newWrap(r))
		}
	}
}
