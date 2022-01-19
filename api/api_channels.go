package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	route{http.MethodPost, "channels/create", channelsCreate}.register()
	route{http.MethodPost, "channels/featured", channelsFeatured}.register()
	route{http.MethodPost, "channels/follow", channelsFollow}.register()
	route{http.MethodPost, "channels/followed", channelsFollowed}.register()
	route{http.MethodPost, "channels/owned", channelsOwned}.register()
	route{http.MethodPost, "channels/show", channelsShow}.register()
	route{http.MethodPost, "channels/unfollow", channelsUnfollow}.register()
	route{http.MethodPost, "channels/update", channelsUpdate}.register()
}

// channelsCreate - channels/create
func channelsCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// channelsFeatured - channels/featured
func channelsFeatured(c *gin.Context) {
	// TODO
	placeholder(c)
}

// channelsFollow - channels/follow
func channelsFollow(c *gin.Context) {
	// TODO
	placeholder(c)
}

// channelsFollowed - channels/followed
func channelsFollowed(c *gin.Context) {
	// TODO
	placeholder(c)
}

// channelsOwned - channels/owned
func channelsOwned(c *gin.Context) {
	// TODO
	placeholder(c)
}

// channelsShow - channels/show
func channelsShow(c *gin.Context) {
	// TODO
	placeholder(c)
}

// channelsUnfollow - channels/unfollow
func channelsUnfollow(c *gin.Context) {
	// TODO
	placeholder(c)
}

// channelsUpdate - channels/update
func channelsUpdate(c *gin.Context) {
	// TODO
	placeholder(c)
}
