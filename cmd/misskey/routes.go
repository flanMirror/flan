package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"random.chars.jp/git/misskey/data/payload"
	"random.chars.jp/git/misskey/db"
	"random.chars.jp/git/misskey/db/orm"
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

	router.GET("/@:user_str", setNoFrame, func(ctx *gin.Context) {
		var (
			atRoot bool
		)

		userStr := ctx.Param("user_str")
		segments := strings.SplitN(userStr, "/", 2)
		switch len(segments) {
		case 1:
			atRoot = true
			if feedState := getFeedState(userStr); feedState != feedNone {
				doFeed(ctx, userStr, feedState)
				return
			}
		case 2:
			userStr = segments[0]
			atRoot = false
		}

		// get user and handle user root and sub
		user := getUser(ctx, userStr)
		if user == nil {
			serveBase(ctx)
			return
		}

		var (
			profile *orm.UserProfile
		)

		metum := db.Meta.Get()
		if metum == nil {
			log.Println("meta cache is nil")
			serveBase(ctx)
			return
		}

		if p, err := orm.FindUserProfileG(ctx, user.ID); err != nil {
			log.Printf("error getting user profile for user %s: %s", user.ID, err)
			serveBase(ctx)
			return
		} else {
			profile = p
		}
		ctx.HTML(http.StatusOK, "user.tmpl", getUserParameters(user, profile, atRoot))
	})

	router.GET("/users/:user", func(ctx *gin.Context) {
		if user, err := orm.Users(
			qm.Where(`"id" = ?`, ctx.Param("user")),
			qm.Where(`host is null`),
			qm.Where(`"isSuspended" = false`),
		).OneG(ctx); err != nil || user == nil {
			notFound(ctx)
		} else {
			ctx.Redirect(http.StatusFound, "/@"+user.Username)
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

func serveBase(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", baseTemplate.Data())
}

func notFound(ctx *gin.Context) {
	ctx.Data(http.StatusNotFound, "text/plain", []byte("Not Found"))
}

func getUser(ctx context.Context, str string) *orm.User {
	account := acct.Parse(str)

	var hostMod qm.QueryMod
	if account.Host.IsLocal() {
		hostMod = qm.Where("host is null")
	} else {
		hostMod = qm.Where("host = ?", account.Host)
	}

	if user, err := orm.Users(
		qm.Where(`"usernameLower" = ?`, strings.ToLower(account.Username)),
		hostMod,
		qm.Where(`"isSuspended" = false`),
	).OneG(ctx); err != nil || user == nil {
		return nil
	} else {
		return user
	}
}
