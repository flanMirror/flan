package openapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	route{http.MethodPost, "room/show", roomShow}.register()
	route{http.MethodPost, "room/update", roomUpdate}.register()
}

// roomShow - room/show
func roomShow(c *gin.Context) {
	// TODO
	placeholder(c)
}

// roomUpdate - room/update
func roomUpdate(c *gin.Context) {
	// TODO
	placeholder(c)
}
