package openapi

import (
	"net/http"
)

func init() {
	route{http.MethodPost, "charts/active-users", chartsActiveUsers}.register()
	route{http.MethodPost, "charts/drive", chartsDrive}.register()
	route{http.MethodPost, "charts/federation", chartsFederation}.register()
	route{http.MethodPost, "charts/hashtag", chartsHashtag}.register()
	route{http.MethodPost, "charts/instance", chartsInstance}.register()
	route{http.MethodPost, "charts/network", chartsNetwork}.register()
	route{http.MethodPost, "charts/notes", chartsNotes}.register()
	route{http.MethodPost, "charts/user/drive", chartsUserDrive}.register()
	route{http.MethodPost, "charts/user/following", chartsUserFollowing}.register()
	route{http.MethodPost, "charts/user/notes", chartsUserNotes}.register()
	route{http.MethodPost, "charts/user/reactions", chartsUserReactions}.register()
	route{http.MethodPost, "charts/users", chartsUsers}.register()
}

// chartsActiveUsers - charts/active-users
func chartsActiveUsers(ctx Context) {
	// TODO
	placeholder(ctx)
}

// chartsDrive - charts/drive
func chartsDrive(ctx Context) {
	// TODO
	placeholder(ctx)
}

// chartsFederation - charts/federation
func chartsFederation(ctx Context) {
	// TODO
	placeholder(ctx)
}

// chartsHashtag - charts/hashtag
func chartsHashtag(ctx Context) {
	// TODO
	placeholder(ctx)
}

// chartsInstance - charts/instance
func chartsInstance(ctx Context) {
	// TODO
	placeholder(ctx)
}

// chartsNetwork - charts/network
func chartsNetwork(ctx Context) {
	// TODO
	placeholder(ctx)
}

// chartsNotes - charts/notes
func chartsNotes(ctx Context) {
	// TODO
	placeholder(ctx)
}

// chartsUserDrive - charts/user/drive
func chartsUserDrive(ctx Context) {
	// TODO
	placeholder(ctx)
}

// chartsUserFollowing - charts/user/following
func chartsUserFollowing(ctx Context) {
	// TODO
	placeholder(ctx)
}

// chartsUserNotes - charts/user/notes
func chartsUserNotes(ctx Context) {
	// TODO
	placeholder(ctx)
}

// chartsUserReactions - charts/user/reactions
func chartsUserReactions(ctx Context) {
	// TODO
	placeholder(ctx)
}

// chartsUsers - charts/users
func chartsUsers(ctx Context) {
	// TODO
	placeholder(ctx)
}
