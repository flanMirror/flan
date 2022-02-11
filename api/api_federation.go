package api

import (
	"net/http"
)

func init() {
	route{http.MethodPost, "ap/get", apGet}.register()
	route{http.MethodPost, "ap/show", apShow}.register()
	route{http.MethodPost, "federation/dns", federationDns}.register()
	route{http.MethodPost, "federation/followers", federationFollowers}.register()
	route{http.MethodPost, "federation/following", federationFollowing}.register()
	route{http.MethodPost, "federation/instances", federationInstances}.register()
	route{http.MethodPost, "federation/show-instance", federationShowInstance}.register()
	route{http.MethodPost, "federation/update-remote-user", federationUpdateRemoteUser}.register()
	route{http.MethodPost, "federation/users", federationUsers}.register()
}

// apGet - ap/get
func apGet(ctx Context) {
	// TODO
	placeholder(ctx)
}

// apShow - ap/show
func apShow(ctx Context) {
	// TODO
	placeholder(ctx)
}

// federationDns - federation/dns
func federationDns(ctx Context) {
	// TODO
	placeholder(ctx)
}

// federationFollowers - federation/followers
func federationFollowers(ctx Context) {
	// TODO
	placeholder(ctx)
}

// federationFollowing - federation/following
func federationFollowing(ctx Context) {
	// TODO
	placeholder(ctx)
}

// federationInstances - federation/instances
func federationInstances(ctx Context) {
	// TODO
	placeholder(ctx)
}

// federationShowInstance - federation/show-instance
func federationShowInstance(ctx Context) {
	// TODO
	placeholder(ctx)
}

// federationUpdateRemoteUser - federation/update-remote-user
func federationUpdateRemoteUser(ctx Context) {
	// TODO
	placeholder(ctx)
}

// federationUsers - federation/users
func federationUsers(ctx Context) {
	// TODO
	placeholder(ctx)
}
