package api

import (
	"net/http"
)

func init() {
	route{http.MethodPost, "clips/create", clipsCreate}.register()
	route{http.MethodPost, "clips/delete", clipsDelete}.register()
	route{http.MethodPost, "clips/list", clipsList}.register()
	route{http.MethodPost, "clips/show", clipsShow}.register()
	route{http.MethodPost, "clips/update", clipsUpdate}.register()
	route{http.MethodPost, "notes/clips", notesClips}.register()
}

// clipsCreate - clips/create
func clipsCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// clipsDelete - clips/delete
func clipsDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}

// clipsList - clips/list
func clipsList(ctx Context) {
	// TODO
	placeholder(ctx)
}

// clipsShow - clips/show
func clipsShow(ctx Context) {
	// TODO
	placeholder(ctx)
}

// clipsUpdate - clips/update
func clipsUpdate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesClips - notes/clips
func notesClips(ctx Context) {
	// TODO
	placeholder(ctx)
}
