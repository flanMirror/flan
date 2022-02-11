package api

import (
	"net/http"
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
func usersGroupsCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersGroupsDelete - users/groups/delete
func usersGroupsDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersGroupsInvitationsAccept - users/groups/invitations/accept
func usersGroupsInvitationsAccept(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersGroupsInvitationsReject - users/groups/invitations/reject
func usersGroupsInvitationsReject(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersGroupsInvite - users/groups/invite
func usersGroupsInvite(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersGroupsJoined - users/groups/joined
func usersGroupsJoined(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersGroupsLeave - users/groups/leave
func usersGroupsLeave(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersGroupsOwned - users/groups/owned
func usersGroupsOwned(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersGroupsPull - users/groups/pull
func usersGroupsPull(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersGroupsShow - users/groups/show
func usersGroupsShow(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersGroupsTransfer - users/groups/transfer
func usersGroupsTransfer(ctx Context) {
	// TODO
	placeholder(ctx)
}

// usersGroupsUpdate - users/groups/update
func usersGroupsUpdate(ctx Context) {
	// TODO
	placeholder(ctx)
}
