package openapi

import (
	"net/http"
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
func pagesCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// pagesDelete - pages/delete
func pagesDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}

// pagesFeatured - pages/featured
func pagesFeatured(ctx Context) {
	// TODO
	placeholder(ctx)
}

// pagesLike - pages/like
func pagesLike(ctx Context) {
	// TODO
	placeholder(ctx)
}

// pagesShow - pages/show
func pagesShow(ctx Context) {
	// TODO
	placeholder(ctx)
}

// pagesUnlike - pages/unlike
func pagesUnlike(ctx Context) {
	// TODO
	placeholder(ctx)
}

// pagesUpdate - pages/update
func pagesUpdate(ctx Context) {
	// TODO
	placeholder(ctx)
}
