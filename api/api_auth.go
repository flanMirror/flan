package openapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	route{http.MethodPost, "auth/session/generate", authSessionGenerate}.register()
	route{http.MethodPost, "auth/session/show", authSessionShow}.register()
	route{http.MethodPost, "auth/session/userkey", authSessionUserkey}.register()
}

// authSessionGenerate - auth/session/generate
func authSessionGenerate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// authSessionShow - auth/session/show
func authSessionShow(c *gin.Context) {
	// TODO
	placeholder(c)
}

// authSessionUserkey - auth/session/userkey
func authSessionUserkey(c *gin.Context) {
	// TODO
	placeholder(c)
}
