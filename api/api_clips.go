package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
func clipsCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// clipsDelete - clips/delete
func clipsDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}

// clipsList - clips/list
func clipsList(c *gin.Context) {
	// TODO
	placeholder(c)
}

// clipsShow - clips/show
func clipsShow(c *gin.Context) {
	// TODO
	placeholder(c)
}

// clipsUpdate - clips/update
func clipsUpdate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesClips - notes/clips
func notesClips(c *gin.Context) {
	// TODO
	placeholder(c)
}
