package api

import (
	"net/http"
)

func init() {
	route{http.MethodPost, "channels/pin-note", channelsPinNote}.register()
	route{http.MethodPost, "request-reset-password", requestResetPassword}.register()
	route{http.MethodPost, "reset-db", resetDb}.register()
	route{http.MethodPost, "reset-password", resetPassword}.register()
}

// channelsPinNote - channels/pin-note
func channelsPinNote(ctx Context) {
	// TODO
	placeholder(ctx)
}

// requestResetPassword - request-reset-password
func requestResetPassword(ctx Context) {
	// TODO
	placeholder(ctx)
}

// resetDb - reset-db
func resetDb(ctx Context) {
	// TODO
	placeholder(ctx)
}

// resetPassword - reset-password
func resetPassword(ctx Context) {
	// TODO
	placeholder(ctx)
}
