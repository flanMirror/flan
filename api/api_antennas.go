package openapi

import (
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
func antennasCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// antennasDelete - antennas/delete
func antennasDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}

// antennasList - antennas/list
func antennasList(ctx Context) {
	// TODO
	placeholder(ctx)
}

// antennasNotes - antennas/notes
func antennasNotes(ctx Context) {
	// TODO
	placeholder(ctx)
}

// antennasShow - antennas/show
func antennasShow(ctx Context) {
	// TODO
	placeholder(ctx)
}

// antennasUpdate - antennas/update
func antennasUpdate(ctx Context) {
	// TODO
	placeholder(ctx)
}
