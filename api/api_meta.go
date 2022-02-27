package api

import (
	"log"
	"net/http"

	"random.chars.jp/git/misskey/config"
	"random.chars.jp/git/misskey/data/payload"
)

func init() {
	route{http.MethodPost, "announcements", announcements}.register()
	route{http.MethodPost, "endpoint", endpoint}.register()
	route{http.MethodPost, "endpoints", endpoints}.register()
	route{http.MethodPost, "get-online-users-count", getOnlineUsersCount}.register()
	route{http.MethodPost, "meta", meta}.register()
	route{http.MethodPost, "ping", ping}.register()
	route{http.MethodPost, "server-info", serverInfo}.register()
	route{http.MethodPost, "stats", stats}.register()
}

// announcements - announcements
func announcements(ctx Context) {
	// TODO
	placeholder(ctx)
}

// endpoint - endpoint
func endpoint(ctx Context) {
	// TODO
	placeholder(ctx)
}

// endpoints - endpoints
func endpoints(ctx Context) {
	// TODO
	placeholder(ctx)
}

// getOnlineUsersCount - get-online-users-count
func getOnlineUsersCount(ctx Context) {
	// TODO
	placeholder(ctx)
}

// meta - meta
func meta(ctx Context) {
	var detail bool

	if ok, err := ctx.BindKey("detail", &detail); err != nil {
		if config.Log.Verbose {
			log.Printf("error binding to detail field: %s", err)
		}
		ctx.BadRequest()
		return
	} else {
		if !ok {
			detail = true
		}
	}

	if !detail {
		ctx.RawJSON(http.StatusOK, payload.Meta.Data())
	} else {
		if user, _, ok, err := ctx.Authenticate(); err != nil {
			log.Printf("error authenticating: %s", err)
			ctx.InternalServerError()
			return
		} else if !ok {
			return
		} else if user == nil || !user.IsAdmin {
			ctx.RawJSON(http.StatusOK, payload.MetaDetail.Data())
		} else {
			ctx.RawJSON(http.StatusOK, payload.MetaAdmin.Data())
		}
	}
}

// ping - ping
func ping(ctx Context) {
	// TODO
	placeholder(ctx)
}

// serverInfo - server-info
func serverInfo(ctx Context) {
	ctx.RawJSON(http.StatusOK, payload.ServerInfo.Data())
}

// stats - stats
func stats(ctx Context) {
	// TODO
	placeholder(ctx)
}
