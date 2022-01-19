package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	route{http.MethodPost, "notifications/create", notificationsCreate}.register()
	route{http.MethodPost, "notifications/mark-all-as-read", notificationsMarkAllAsRead}.register()
	route{http.MethodPost, "notifications/read", notificationsRead}.register()
}

// notificationsCreate - notifications/create
func notificationsCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notificationsMarkAllAsRead - notifications/mark-all-as-read
func notificationsMarkAllAsRead(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notificationsRead - notifications/read
func notificationsRead(c *gin.Context) {
	// TODO
	placeholder(c)
}
