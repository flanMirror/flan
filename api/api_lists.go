package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
func usersListsCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersListsDelete - users/lists/delete
func usersListsDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersListsList - users/lists/list
func usersListsList(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersListsPull - users/lists/pull
func usersListsPull(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersListsPush - users/lists/push
func usersListsPush(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersListsShow - users/lists/show
func usersListsShow(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersListsUpdate - users/lists/update
func usersListsUpdate(c *gin.Context) {
	// TODO
	placeholder(c)
}
