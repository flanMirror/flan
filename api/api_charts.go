package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
func chartsActiveUsers(c *gin.Context) {
	// TODO
	placeholder(c)
}

// chartsDrive - charts/drive
func chartsDrive(c *gin.Context) {
	// TODO
	placeholder(c)
}

// chartsFederation - charts/federation
func chartsFederation(c *gin.Context) {
	// TODO
	placeholder(c)
}

// chartsHashtag - charts/hashtag
func chartsHashtag(c *gin.Context) {
	// TODO
	placeholder(c)
}

// chartsInstance - charts/instance
func chartsInstance(c *gin.Context) {
	// TODO
	placeholder(c)
}

// chartsNetwork - charts/network
func chartsNetwork(c *gin.Context) {
	// TODO
	placeholder(c)
}

// chartsNotes - charts/notes
func chartsNotes(c *gin.Context) {
	// TODO
	placeholder(c)
}

// chartsUserDrive - charts/user/drive
func chartsUserDrive(c *gin.Context) {
	// TODO
	placeholder(c)
}

// chartsUserFollowing - charts/user/following
func chartsUserFollowing(c *gin.Context) {
	// TODO
	placeholder(c)
}

// chartsUserNotes - charts/user/notes
func chartsUserNotes(c *gin.Context) {
	// TODO
	placeholder(c)
}

// chartsUserReactions - charts/user/reactions
func chartsUserReactions(c *gin.Context) {
	// TODO
	placeholder(c)
}

// chartsUsers - charts/users
func chartsUsers(c *gin.Context) {
	// TODO
	placeholder(c)
}
