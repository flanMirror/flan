package api

import (
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
func galleryFeatured(ctx Context) {
	// TODO
	placeholder(ctx)
}

// galleryPopular - gallery/popular
func galleryPopular(ctx Context) {
	// TODO
	placeholder(ctx)
}

// galleryPosts - gallery/posts
func galleryPosts(ctx Context) {
	// TODO
	placeholder(ctx)
}

// galleryPostsCreate - gallery/posts/create
func galleryPostsCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// galleryPostsDelete - gallery/posts/delete
func galleryPostsDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}

// galleryPostsLike - gallery/posts/like
func galleryPostsLike(ctx Context) {
	// TODO
	placeholder(ctx)
}

// galleryPostsShow - gallery/posts/show
func galleryPostsShow(ctx Context) {
	// TODO
	placeholder(ctx)
}

// galleryPostsUnlike - gallery/posts/unlike
func galleryPostsUnlike(ctx Context) {
	// TODO
	placeholder(ctx)
}

// galleryPostsUpdate - gallery/posts/update
func galleryPostsUpdate(ctx Context) {
	// TODO
	placeholder(ctx)
}
