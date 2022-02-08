package openapi

import (
	"net/http"
)

func init() {
	route{http.MethodPost, "auth/session/generate", authSessionGenerate}.register()
	route{http.MethodPost, "auth/session/show", authSessionShow}.register()
	route{http.MethodPost, "auth/session/userkey", authSessionUserkey}.register()
}

// authSessionGenerate - auth/session/generate
func authSessionGenerate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// authSessionShow - auth/session/show
func authSessionShow(ctx Context) {
	// TODO
	placeholder(ctx)
}

// authSessionUserkey - auth/session/userkey
func authSessionUserkey(ctx Context) {
	// TODO
	placeholder(ctx)
}
