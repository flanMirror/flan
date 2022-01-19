package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	route{http.MethodPost, "pages/create", pagesCreate}.register()
	route{http.MethodPost, "pages/delete", pagesDelete}.register()
	route{http.MethodPost, "pages/featured", pagesFeatured}.register()
	route{http.MethodPost, "pages/like", pagesLike}.register()
	route{http.MethodPost, "pages/show", pagesShow}.register()
	route{http.MethodPost, "pages/unlike", pagesUnlike}.register()
	route{http.MethodPost, "pages/update", pagesUpdate}.register()
}

// pagesCreate - pages/create
func pagesCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// pagesDelete - pages/delete
func pagesDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}

// pagesFeatured - pages/featured
func pagesFeatured(c *gin.Context) {
	// TODO
	placeholder(c)
}

// pagesLike - pages/like
func pagesLike(c *gin.Context) {
	// TODO
	placeholder(c)
}

// pagesShow - pages/show
func pagesShow(c *gin.Context) {
	// TODO
	placeholder(c)
}

// pagesUnlike - pages/unlike
func pagesUnlike(c *gin.Context) {
	// TODO
	placeholder(c)
}

// pagesUpdate - pages/update
func pagesUpdate(c *gin.Context) {
	// TODO
	placeholder(c)
}
