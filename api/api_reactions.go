package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	route{http.MethodPost, "notes/reactions/create", notesReactionsCreate}.register()
	route{http.MethodPost, "notes/reactions/delete", notesReactionsDelete}.register()
}

// notesReactionsCreate - notes/reactions/create
func notesReactionsCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesReactionsDelete - notes/reactions/delete
func notesReactionsDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}
