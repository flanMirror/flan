package openapi

import (
	"github.com/gin-gonic/gin"
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
func followingCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// followingDelete - following/delete
func followingDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}

// followingInvalidate - following/invalidate
func followingInvalidate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// followingRequestsAccept - following/requests/accept
func followingRequestsAccept(c *gin.Context) {
	// TODO
	placeholder(c)
}

// followingRequestsCancel - following/requests/cancel
func followingRequestsCancel(c *gin.Context) {
	// TODO
	placeholder(c)
}

// followingRequestsList - following/requests/list
func followingRequestsList(c *gin.Context) {
	// TODO
	placeholder(c)
}

// followingRequestsReject - following/requests/reject
func followingRequestsReject(c *gin.Context) {
	// TODO
	placeholder(c)
}
