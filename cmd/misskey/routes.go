package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"random.chars.jp/git/misskey/config"
	"random.chars.jp/git/misskey/data/payload"
	"random.chars.jp/git/misskey/db/models"
	"random.chars.jp/git/misskey/db/pack"
	"random.chars.jp/git/misskey/feed"
	"random.chars.jp/git/misskey/sdk/acct"
	"random.chars.jp/git/misskey/spec"
)

func routesSetup() {
	router.GET("/manifest.json", setNoFrame, func(ctx *gin.Context) {
		ctx.Header("Cache-Control", "max-age=300")
		ctx.Data(http.StatusOK, gin.MIMEJSON, payload.Manifest.Data())
	})
	router.GET("/api-doc", setNoFrame, func(ctx *gin.Context) {
		ctx.Header("Cache-Control", "max-age=0")
		ctx.Data(http.StatusOK, gin.MIMEHTML, apiDoc)
	})
	router.GET("/api.json", setNoFrame, func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, gin.MIMEJSON, spec.JSON())
	})

	// TODO: move these to the monolithic GET handler below
	//router.GET("/@:user.atom", setNoFrame, func(ctx *gin.Context) {
	//	emitter := getFeed(ctx, ctx.Param("user"))
	//	if emitter == nil {
	//		notFound(ctx)
	//		return
	//	}
	//	ctx.Data(http.StatusOK, "application/atom+xml; charset=utf-8", emitter.Atom1())
	//})
	//router.GET("/@:user.rss", setNoFrame, func(ctx *gin.Context) {
	//	emitter := getFeed(ctx, ctx.Param("user"))
	//	if emitter == nil {
	//		notFound(ctx)
	//		return
	//	}
	//	ctx.Data(http.StatusOK, "application/rss+xml; charset=utf-8", emitter.RSS2())
	//})
	router.GET("/@:user_str", setNoFrame, func(ctx *gin.Context) {
		user := ctx.Param("user_str")
		segments := strings.SplitN(user, "/", 2)
		switch len(segments) {
		case 1:
			if strings.HasSuffix(user, ".json") {
				user = user[:len(user)-5]
				emitter := getFeed(ctx, user)
				if emitter == nil {
					notFound(ctx)
					return
				}
				ctx.JSON(http.StatusOK, emitter.JSON())
				return
			}
		case 2:
			user = segments[0]
			//sub := segments[1]
			// TODO: handle user profile and sub
		}
	})

	// FIXME: uncomment this after proper websocket middleware is implemented
	//router.GET("/streaming", func(ctx *gin.Context) {
	//	ctx.Header("Cache-Control", "private, max-age=0")
	//	ctx.Data(http.StatusServiceUnavailable,
	//		"text/plain; charset=utf-8",
	//		[]byte("Service Unavailable"))
	//})
}

func notFound(ctx *gin.Context) {
	ctx.Data(http.StatusNotFound, "text/plain", []byte("Not Found"))
}

func getFeed(ctx context.Context, str string) *feed.Emitter {
	account := acct.Parse(str)

	query := models.Users()
	if account.Host.IsLocal() {
		query = models.Users(
			qm.Where(`"usernameLower" = ?`, strings.ToLower(account.Username)),
			// no host here
			qm.Where(`"isSuspended" = false`),
		)
	} else {
		query = models.Users(
			qm.Where(`"usernameLower" = ?`, strings.ToLower(account.Username)),
			qm.Where("host = ?", account.Host),
			qm.Where(`"isSuspended" = false`),
		)
	}

	if user, err := query.OneG(ctx); err != nil || user == nil {
		return nil
	} else {
		var emitter *feed.Emitter
		if emitter, err = pack.Feed(ctx, user); err != nil {
			if config.Log.Verbose {
				log.Printf("error packing feed: %s", err)
			}
			return nil
		}
		return emitter
	}
}
