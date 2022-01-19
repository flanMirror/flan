package openapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	route{http.MethodPost, "app/create", appCreate}.register()
	route{http.MethodPost, "app/show", appShow}.register()
}

// appCreate - app/create
func appCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// appShow - app/show
func appShow(c *gin.Context) {
	// TODO
	placeholder(c)
}
