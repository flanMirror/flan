package openapi

import (
	"github.com/gin-gonic/gin"
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
func emailAddressAvailable(c *gin.Context) {
	// TODO
	placeholder(c)
}

// pinnedUsers - pinned-users
func pinnedUsers(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usernameAvailable - username/available
func usernameAvailable(c *gin.Context) {
	// TODO
	placeholder(c)
}

// users - users
func users(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersClips - users/clips
func usersClips(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersFollowers - users/followers
func usersFollowers(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersFollowing - users/following
func usersFollowing(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersGalleryPosts - users/gallery/posts
func usersGalleryPosts(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersGetFrequentlyRepliedusers - users/get-frequently-replied-users
func usersGetFrequentlyRepliedusers(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersNotes - users/notes
func usersNotes(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersPages - users/pages
func usersPages(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersReactions - users/reactions
func usersReactions(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersRecommendation - users/recommendation
func usersRecommendation(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersRelation - users/relation
func usersRelation(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersReportAbuse - users/report-abuse
func usersReportAbuse(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersSearch - users/search
func usersSearch(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersSearchByUsernameAndHost - users/search-by-username-and-host
func usersSearchByUsernameAndHost(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersShow - users/show
func usersShow(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersStats - users/stats
func usersStats(c *gin.Context) {
	// TODO
	placeholder(c)
}
