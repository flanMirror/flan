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
	"random.chars.jp/git/misskey/db/models"
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
			ctx.Next()
			return
		}

		var (
			profile *models.UserProfile
		)

		metum := db.Meta.Get()
		if metum == nil {
			log.Println("meta cache is nil")
			ctx.Next()
			return
		}

		if p, err := models.FindUserProfileG(ctx, user.ID); err != nil {
			log.Printf("error getting user profile for user %s: %s", user.ID, err)
			ctx.Next()
			return
		} else {
			profile = p
		}
		ctx.HTML(http.StatusOK, "user.tmpl", getUserParameters(user, profile, atRoot))
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

func getUser(ctx context.Context, str string) *models.User {
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
		return user
	}
}
