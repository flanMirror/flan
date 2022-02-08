package openapi

import (
	"net/http"
)

func init() {
	route{http.MethodPost, "email-address/available", emailAddressAvailable}.register()
	route{http.MethodPost, "pinned-users", pinnedUsers}.register()
	route{http.MethodPost, "username/available", usernameAvailable}.register()
	route{http.MethodPost, "users", users}.register()
	route{http.MethodPost, "users/clips", usersClips}.register()
	route{http.MethodPost, "users/followers", usersFollowers}.register()
	route{http.MethodPost, "users/following", usersFollowing}.register()
	route{http.MethodPost, "users/gallery/posts", usersGalleryPosts}.register()
	route{http.MethodPost, "users/get-frequently-replied-users", usersGetFrequentlyRepliedusers}.register()
	route{http.MethodPost, "users/notes", usersNotes}.register()
	route{http.MethodPost, "users/pages", usersPages}.register()
	route{http.MethodPost, "users/reactions", usersReactions}.register()
	route{http.MethodPost, "users/recommendation", usersRecommendation}.register()
	route{http.MethodPost, "users/relation", usersRelation}.register()
	route{http.MethodPost, "users/report-abuse", usersReportAbuse}.register()
	route{http.MethodPost, "users/search", usersSearch}.register()
	route{http.MethodPost, "users/search-by-username-and-host", usersSearchByUsernameAndHost}.register()
	route{http.MethodPost, "users/show", usersShow}.register()
	route{http.MethodPost, "users/stats", usersStats}.register()
}

// emailAddressAvailable - email-address/available
func emailAddressAvailable(ctx Context) {
	// TODO
	placeholder(ctx)
}

// pinnedUsers - pinned-users
func pinnedUsers(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usernameAvailable - username/available
func usernameAvailable(ctx Context) {
	// TODO
	placeholder(ctx)
}

// users - users
func users(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersClips - users/clips
func usersClips(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersFollowers - users/followers
func usersFollowers(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersFollowing - users/following
func usersFollowing(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersGalleryPosts - users/gallery/posts
func usersGalleryPosts(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersGetFrequentlyRepliedusers - users/get-frequently-replied-users
func usersGetFrequentlyRepliedusers(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersNotes - users/notes
func usersNotes(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersPages - users/pages
func usersPages(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersReactions - users/reactions
func usersReactions(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersRecommendation - users/recommendation
func usersRecommendation(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersRelation - users/relation
func usersRelation(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersReportAbuse - users/report-abuse
func usersReportAbuse(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersSearch - users/search
func usersSearch(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersSearchByUsernameAndHost - users/search-by-username-and-host
func usersSearchByUsernameAndHost(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersShow - users/show
func usersShow(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersStats - users/stats
func usersStats(ctx Context) {
	// TODO
	placeholder(ctx)
}
