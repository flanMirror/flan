package openapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"random.chars.jp/git/misskey/api/payload"
)

func init() {
	route{http.MethodPost, "announcements", announcements}.register()
	route{http.MethodPost, "endpoint", endpoint}.register()
	route{http.MethodPost, "endpoints", endpoints}.register()
	route{http.MethodPost, "get-online-users-count", getOnlineUsersCount}.register()
	route{http.MethodPost, "meta", meta}.register()
	route{http.MethodPost, "ping", ping}.register()
	route{http.MethodPost, "server-info", serverInfo}.register()
	route{http.MethodPost, "stats", stats}.register()
}

// announcements - announcements
func announcements(c *gin.Context) {
	// TODO
	placeholder(c)
}

// endpoint - endpoint
func endpoint(c *gin.Context) {
	// TODO
	placeholder(c)
}

// endpoints - endpoints
func endpoints(c *gin.Context) {
	// TODO
	placeholder(c)
}

// getOnlineUsersCount - get-online-users-count
func getOnlineUsersCount(c *gin.Context) {
	// TODO
	placeholder(c)
}

// meta - meta
func meta(c *gin.Context) {
	// TODO
	placeholder(c)
}

// ping - ping
func ping(c *gin.Context) {
	// TODO
	placeholder(c)
}

// serverInfo - server-info
func serverInfo(c *gin.Context) {
	c.Data(http.StatusOK, gin.MIMEJSON, payload.ServerInfo.Data())
}

// stats - stats
func stats(c *gin.Context) {
	// TODO
	placeholder(c)
}
