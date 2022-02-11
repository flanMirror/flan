package api

import (
	"net/http"
)

func init() {
	route{http.MethodPost, "users/lists/create", usersListsCreate}.register()
	route{http.MethodPost, "users/lists/delete", usersListsDelete}.register()
	route{http.MethodPost, "users/lists/list", usersListsList}.register()
	route{http.MethodPost, "users/lists/pull", usersListsPull}.register()
	route{http.MethodPost, "users/lists/push", usersListsPush}.register()
	route{http.MethodPost, "users/lists/show", usersListsShow}.register()
	route{http.MethodPost, "users/lists/update", usersListsUpdate}.register()
}

// usersListsCreate - users/lists/create
func usersListsCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersListsDelete - users/lists/delete
func usersListsDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersListsList - users/lists/list
func usersListsList(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersListsPull - users/lists/pull
func usersListsPull(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersListsPush - users/lists/push
func usersListsPush(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersListsShow - users/lists/show
func usersListsShow(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersListsUpdate - users/lists/update
func usersListsUpdate(ctx Context) {
	// TODO
	placeholder(ctx)
}
