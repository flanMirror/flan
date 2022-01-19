package openapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	route{http.MethodPost, "gallery/featured", galleryFeatured}.register()
	route{http.MethodPost, "gallery/popular", galleryPopular}.register()
	route{http.MethodPost, "gallery/posts", galleryPosts}.register()
	route{http.MethodPost, "gallery/posts/create", galleryPostsCreate}.register()
	route{http.MethodPost, "gallery/posts/delete", galleryPostsDelete}.register()
	route{http.MethodPost, "gallery/posts/like", galleryPostsLike}.register()
	route{http.MethodPost, "gallery/posts/show", galleryPostsShow}.register()
	route{http.MethodPost, "gallery/posts/unlike", galleryPostsUnlike}.register()
	route{http.MethodPost, "gallery/posts/update", galleryPostsUpdate}.register()
}

// galleryFeatured - gallery/featured
func galleryFeatured(c *gin.Context) {
	// TODO
	placeholder(c)
}

// galleryPopular - gallery/popular
func galleryPopular(c *gin.Context) {
	// TODO
	placeholder(c)
}

// galleryPosts - gallery/posts
func galleryPosts(c *gin.Context) {
	// TODO
	placeholder(c)
}

// galleryPostsCreate - gallery/posts/create
func galleryPostsCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// galleryPostsDelete - gallery/posts/delete
func galleryPostsDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}

// galleryPostsLike - gallery/posts/like
func galleryPostsLike(c *gin.Context) {
	// TODO
	placeholder(c)
}

// galleryPostsShow - gallery/posts/show
func galleryPostsShow(c *gin.Context) {
	// TODO
	placeholder(c)
}

// galleryPostsUnlike - gallery/posts/unlike
func galleryPostsUnlike(c *gin.Context) {
	// TODO
	placeholder(c)
}

// galleryPostsUpdate - gallery/posts/update
func galleryPostsUpdate(c *gin.Context) {
	// TODO
	placeholder(c)
}
