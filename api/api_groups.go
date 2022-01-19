package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	route{http.MethodPost, "users/groups/create", usersGroupsCreate}.register()
	route{http.MethodPost, "users/groups/delete", usersGroupsDelete}.register()
	route{http.MethodPost, "users/groups/invitations/accept", usersGroupsInvitationsAccept}.register()
	route{http.MethodPost, "users/groups/invitations/reject", usersGroupsInvitationsReject}.register()
	route{http.MethodPost, "users/groups/invite", usersGroupsInvite}.register()
	route{http.MethodPost, "users/groups/joined", usersGroupsJoined}.register()
	route{http.MethodPost, "users/groups/leave", usersGroupsLeave}.register()
	route{http.MethodPost, "users/groups/owned", usersGroupsOwned}.register()
	route{http.MethodPost, "users/groups/pull", usersGroupsPull}.register()
	route{http.MethodPost, "users/groups/show", usersGroupsShow}.register()
	route{http.MethodPost, "users/groups/transfer", usersGroupsTransfer}.register()
	route{http.MethodPost, "users/groups/update", usersGroupsUpdate}.register()
}

// usersGroupsCreate - users/groups/create
func usersGroupsCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersGroupsDelete - users/groups/delete
func usersGroupsDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersGroupsInvitationsAccept - users/groups/invitations/accept
func usersGroupsInvitationsAccept(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersGroupsInvitationsReject - users/groups/invitations/reject
func usersGroupsInvitationsReject(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersGroupsInvite - users/groups/invite
func usersGroupsInvite(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersGroupsJoined - users/groups/joined
func usersGroupsJoined(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersGroupsLeave - users/groups/leave
func usersGroupsLeave(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersGroupsOwned - users/groups/owned
func usersGroupsOwned(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersGroupsPull - users/groups/pull
func usersGroupsPull(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersGroupsShow - users/groups/show
func usersGroupsShow(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersGroupsTransfer - users/groups/transfer
func usersGroupsTransfer(c *gin.Context) {
	// TODO
	placeholder(c)
}

// usersGroupsUpdate - users/groups/update
func usersGroupsUpdate(c *gin.Context) {
	// TODO
	placeholder(c)
}
