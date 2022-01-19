package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	route{http.MethodPost, "messaging/history", messagingHistory}.register()
	route{http.MethodPost, "messaging/messages", messagingMessages}.register()
	route{http.MethodPost, "messaging/messages/create", messagingMessagesCreate}.register()
	route{http.MethodPost, "messaging/messages/delete", messagingMessagesDelete}.register()
	route{http.MethodPost, "messaging/messages/read", messagingMessagesRead}.register()
}

// messagingHistory - messaging/history
func messagingHistory(c *gin.Context) {
	// TODO
	placeholder(c)
}

// messagingMessages - messaging/messages
func messagingMessages(c *gin.Context) {
	// TODO
	placeholder(c)
}

// messagingMessagesCreate - messaging/messages/create
func messagingMessagesCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// messagingMessagesDelete - messaging/messages/delete
func messagingMessagesDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}

// messagingMessagesRead - messaging/messages/read
func messagingMessagesRead(c *gin.Context) {
	// TODO
	placeholder(c)
}
