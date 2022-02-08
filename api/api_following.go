package openapi

import (
	"net/http"
)

func init() {
	route{http.MethodPost, "following/create", followingCreate}.register()
	route{http.MethodPost, "following/delete", followingDelete}.register()
	route{http.MethodPost, "following/invalidate", followingInvalidate}.register()
	route{http.MethodPost, "following/requests/accept", followingRequestsAccept}.register()
	route{http.MethodPost, "following/requests/cancel", followingRequestsCancel}.register()
	route{http.MethodPost, "following/requests/list", followingRequestsList}.register()
	route{http.MethodPost, "following/requests/reject", followingRequestsReject}.register()
}

// followingCreate - following/create
func followingCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// followingDelete - following/delete
func followingDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}

// followingInvalidate - following/invalidate
func followingInvalidate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// followingRequestsAccept - following/requests/accept
func followingRequestsAccept(ctx Context) {
	// TODO
	placeholder(ctx)
}

// followingRequestsCancel - following/requests/cancel
func followingRequestsCancel(ctx Context) {
	// TODO
	placeholder(ctx)
}

// followingRequestsList - following/requests/list
func followingRequestsList(ctx Context) {
	// TODO
	placeholder(ctx)
}

// followingRequestsReject - following/requests/reject
func followingRequestsReject(ctx Context) {
	// TODO
	placeholder(ctx)
}
