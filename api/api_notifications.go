package api

import (
	"net/http"
)

func init() {
	route{http.MethodPost, "notifications/create", notificationsCreate}.register()
	route{http.MethodPost, "notifications/mark-all-as-read", notificationsMarkAllAsRead}.register()
	route{http.MethodPost, "notifications/read", notificationsRead}.register()
}

// notificationsCreate - notifications/create
func notificationsCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notificationsMarkAllAsRead - notifications/mark-all-as-read
func notificationsMarkAllAsRead(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notificationsRead - notifications/read
func notificationsRead(ctx Context) {
	// TODO
	placeholder(ctx)
}
