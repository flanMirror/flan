package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	route{http.MethodPost, "channels/pin-note", channelsPinNote}.register()
	route{http.MethodPost, "request-reset-password", requestResetPassword}.register()
	route{http.MethodPost, "reset-db", resetDb}.register()
	route{http.MethodPost, "reset-password", resetPassword}.register()
}

// channelsPinNote - channels/pin-note
func channelsPinNote(c *gin.Context) {
	// TODO
	placeholder(c)
}

// requestResetPassword - request-reset-password
func requestResetPassword(c *gin.Context) {
	// TODO
	placeholder(c)
}

// resetDb - reset-db
func resetDb(c *gin.Context) {
	// TODO
	placeholder(c)
}

// resetPassword - reset-password
func resetPassword(c *gin.Context) {
	// TODO
	placeholder(c)
}
