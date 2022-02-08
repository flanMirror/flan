package openapi

import (
	"log"
	"net/http"
	"random.chars.jp/git/misskey/api/payload"
	"random.chars.jp/git/misskey/config"
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
	p := struct {
		Detail *bool
	}{}
	if err := ctx.Bind(&p); err != nil {
		if config.Log.Verbose {
			log.Printf("error parsing meta params: %s", err)
		}
		ctx.BadRequest()
		return
	}

	detail := true
	if p.Detail != nil {
		detail = *p.Detail
	}

	if !detail {
		ctx.RawJSON(http.StatusOK, payload.Meta.Data())
	} else {
		// TODO: admin
		ctx.RawJSON(http.StatusOK, payload.MetaDetail.Data())
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
