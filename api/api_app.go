package api

import (
	"net/http"
)

func init() {
	route{http.MethodPost, "app/create", appCreate}.register()
	route{http.MethodPost, "app/show", appShow}.register()
}

// appCreate - app/create
func appCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// appShow - app/show
func appShow(ctx Context) {
	// TODO
	placeholder(ctx)
}
