package api

import (
	"net/http"
)

func init() {
	route{http.MethodPost, "messaging/history", messagingHistory}.register()
	route{http.MethodPost, "messaging/messages", messagingMessages}.register()
	route{http.MethodPost, "messaging/messages/create", messagingMessagesCreate}.register()
	route{http.MethodPost, "messaging/messages/delete", messagingMessagesDelete}.register()
	route{http.MethodPost, "messaging/messages/read", messagingMessagesRead}.register()
}

// messagingHistory - messaging/history
func messagingHistory(ctx Context) {
	// TODO
	placeholder(ctx)
}

// messagingMessages - messaging/messages
func messagingMessages(ctx Context) {
	// TODO
	placeholder(ctx)
}

// messagingMessagesCreate - messaging/messages/create
func messagingMessagesCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// messagingMessagesDelete - messaging/messages/delete
func messagingMessagesDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}

// messagingMessagesRead - messaging/messages/read
func messagingMessagesRead(ctx Context) {
	// TODO
	placeholder(ctx)
}
