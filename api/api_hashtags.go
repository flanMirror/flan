package openapi

import (
	"github.com/gin-gonic/gin"
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
func hashtagsList(c *gin.Context) {
	// TODO
	placeholder(c)
}

// hashtagsSearch - hashtags/search
func hashtagsSearch(c *gin.Context) {
	// TODO
	placeholder(c)
}

// hashtagsShow - hashtags/show
func hashtagsShow(c *gin.Context) {
	// TODO
	placeholder(c)
}

// hashtagsTrend - hashtags/trend
func hashtagsTrend(c *gin.Context) {
	// TODO
	placeholder(c)
}

// hashtagsUsers - hashtags/users
func hashtagsUsers(c *gin.Context) {
	// TODO
	placeholder(c)
}
