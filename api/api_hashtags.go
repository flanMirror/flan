package openapi

import (
	"net/http"
)

func init() {
	route{http.MethodPost, "hashtags/list", hashtagsList}.register()
	route{http.MethodPost, "hashtags/search", hashtagsSearch}.register()
	route{http.MethodPost, "hashtags/show", hashtagsShow}.register()
	route{http.MethodPost, "hashtags/trend", hashtagsTrend}.register()
	route{http.MethodPost, "hashtags/users", hashtagsUsers}.register()
}

// hashtagsList - hashtags/list
func hashtagsList(ctx Context) {
	// TODO
	placeholder(ctx)
}

// hashtagsSearch - hashtags/search
func hashtagsSearch(ctx Context) {
	// TODO
	placeholder(ctx)
}

// hashtagsShow - hashtags/show
func hashtagsShow(ctx Context) {
	// TODO
	placeholder(ctx)
}

// hashtagsTrend - hashtags/trend
func hashtagsTrend(ctx Context) {
	// TODO
	placeholder(ctx)
}

// hashtagsUsers - hashtags/users
func hashtagsUsers(ctx Context) {
	// TODO
	placeholder(ctx)
}
