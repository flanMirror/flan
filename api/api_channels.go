package api

import (
	"net/http"
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
func channelsCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// channelsFeatured - channels/featured
func channelsFeatured(ctx Context) {
	// TODO
	placeholder(ctx)
}

// channelsFollow - channels/follow
func channelsFollow(ctx Context) {
	// TODO
	placeholder(ctx)
}

// channelsFollowed - channels/followed
func channelsFollowed(ctx Context) {
	// TODO
	placeholder(ctx)
}

// channelsOwned - channels/owned
func channelsOwned(ctx Context) {
	// TODO
	placeholder(ctx)
}

// channelsShow - channels/show
func channelsShow(ctx Context) {
	// TODO
	placeholder(ctx)
}

// channelsUnfollow - channels/unfollow
func channelsUnfollow(ctx Context) {
	// TODO
	placeholder(ctx)
}

// channelsUpdate - channels/update
func channelsUpdate(ctx Context) {
	// TODO
	placeholder(ctx)
}
