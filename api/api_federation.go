package openapi

import (
	"github.com/gin-gonic/gin"
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
func apGet(c *gin.Context) {
	// TODO
	placeholder(c)
}

// apShow - ap/show
func apShow(c *gin.Context) {
	// TODO
	placeholder(c)
}

// federationDns - federation/dns
func federationDns(c *gin.Context) {
	// TODO
	placeholder(c)
}

// federationFollowers - federation/followers
func federationFollowers(c *gin.Context) {
	// TODO
	placeholder(c)
}

// federationFollowing - federation/following
func federationFollowing(c *gin.Context) {
	// TODO
	placeholder(c)
}

// federationInstances - federation/instances
func federationInstances(c *gin.Context) {
	// TODO
	placeholder(c)
}

// federationShowInstance - federation/show-instance
func federationShowInstance(c *gin.Context) {
	// TODO
	placeholder(c)
}

// federationUpdateRemoteUser - federation/update-remote-user
func federationUpdateRemoteUser(c *gin.Context) {
	// TODO
	placeholder(c)
}

// federationUsers - federation/users
func federationUsers(c *gin.Context) {
	// TODO
	placeholder(c)
}
