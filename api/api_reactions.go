package api

import (
	"net/http"
)

func init() {
	route{http.MethodPost, "notes/reactions/create", notesReactionsCreate}.register()
	route{http.MethodPost, "notes/reactions/delete", notesReactionsDelete}.register()
}

// notesReactionsCreate - notes/reactions/create
func notesReactionsCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesReactionsDelete - notes/reactions/delete
func notesReactionsDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}
