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

const (
	feedNone = iota
	feedAtom1
	feedRSS2
	feedJSON
)

var feedStateLUT = []int{
	// this should never be accessed
	0,
	// .atom
	5,
	// .rss
	4,
	// .json
	5,
}

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

	router.GET("/@:user_str", setNoFrame, func(ctx *gin.Context) {
		user := ctx.Param("user_str")
		segments := strings.SplitN(user, "/", 2)
		switch len(segments) {
		case 1:
			feedState := feedNone
			if strings.HasSuffix(user, ".json") {
				feedState = feedJSON
			} else if strings.HasSuffix(user, ".rss") {
				feedState = feedRSS2
			} else if strings.HasSuffix(user, ".atom") {
				feedState = feedAtom1
			}

			if feedState != feedNone {
				user = user[:len(user)-feedStateLUT[feedState]]
				emitter := getFeed(ctx, user)
				if emitter == nil {
					notFound(ctx)
					return
				}
				switch feedState {
				case feedJSON:
					ctx.Header("Content-Type", "application/json; charset=utf-8")
					ctx.JSON(http.StatusOK, emitter.JSON())
				case feedRSS2:
					if data, err := emitter.RSS2().XML(); err != nil {
						log.Printf("error generating RSS 2.0 feed for user %s: %s", user, err)
						ctx.String(http.StatusInternalServerError, "Internal Server Error")
						return
					} else {
						ctx.Data(http.StatusOK, "application/rss+xml; charset=utf-8", data)
					}
				case feedAtom1:
					if data, err := emitter.Atom1().XML(); err != nil {
						log.Printf("error generating RSS 2.0 feed for user %s: %s", user, err)
						ctx.String(http.StatusInternalServerError, "Internal Server Error")
						return
					} else {
						ctx.Data(http.StatusOK, "application/atom+xml; charset=utf-8", data)
					}
				}
				return
			}

			// TODO: handle user profile with no sub
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

	var hostMod qm.QueryMod
	if account.Host.IsLocal() {
		hostMod = qm.Where("host is null")
	} else {
		hostMod = qm.Where("host = ?", account.Host)
	}

	if user, err := models.Users(
		qm.Where(`"usernameLower" = ?`, strings.ToLower(account.Username)),
		hostMod,
		qm.Where(`"isSuspended" = false`),
	).OneG(ctx); err != nil || user == nil {
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
