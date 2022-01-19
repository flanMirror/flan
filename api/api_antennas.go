package openapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	route{http.MethodPost, "antennas/create", antennasCreate}.register()
	route{http.MethodPost, "antennas/delete", antennasDelete}.register()
	route{http.MethodPost, "antennas/list", antennasList}.register()
	route{http.MethodPost, "antennas/notes", antennasNotes}.register()
	route{http.MethodPost, "antennas/show", antennasShow}.register()
	route{http.MethodPost, "antennas/update", antennasUpdate}.register()
}

// antennasCreate - antennas/create
func antennasCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// antennasDelete - antennas/delete
func antennasDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}

// antennasList - antennas/list
func antennasList(c *gin.Context) {
	// TODO
	placeholder(c)
}

// antennasNotes - antennas/notes
func antennasNotes(c *gin.Context) {
	// TODO
	placeholder(c)
}

// antennasShow - antennas/show
func antennasShow(c *gin.Context) {
	// TODO
	placeholder(c)
}

// antennasUpdate - antennas/update
func antennasUpdate(c *gin.Context) {
	// TODO
	placeholder(c)
}
