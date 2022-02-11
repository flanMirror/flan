package api

import (
	"net/http"
)

func init() {
	route{http.MethodPost, "room/show", roomShow}.register()
	route{http.MethodPost, "room/update", roomUpdate}.register()
}

// roomShow - room/show
func roomShow(ctx Context) {
	// TODO
	placeholder(ctx)
}

// roomUpdate - room/update
func roomUpdate(ctx Context) {
	// TODO
	placeholder(ctx)
}
