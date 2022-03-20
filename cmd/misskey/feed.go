package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"random.chars.jp/git/misskey/config"
	"random.chars.jp/git/misskey/db/pack"
	"random.chars.jp/git/misskey/feed"
)

const (
	feedNone = iota
	feedAtom1
	feedRSS2
	feedJSON
)

var feedStateLUT = []int{
	// feedNone  - this should never be accessed
	0,
	// feedAtom1 - .atom
	5,
	// feedRSS2  - .rss
	4,
	// feedJSON  - .json
	5,
}

func getFeedState(userStr string) int {
	if strings.HasSuffix(userStr, ".json") {
		return feedJSON
	} else if strings.HasSuffix(userStr, ".rss") {
		return feedRSS2
	} else if strings.HasSuffix(userStr, ".atom") {
		return feedAtom1
	}
	return feedNone
}

func getFeed(ctx context.Context, str string) *feed.Emitter {
	user := getUser(ctx, str)
	if user == nil {
		return nil
	}
	if emitter, err := pack.Feed(ctx, user); err != nil {
		if config.Log.Verbose {
			log.Printf("error packing feed: %s", err)
		}
		return nil
	} else {
		return emitter
	}
}
func doFeed(ctx *gin.Context, userStr string, feedState int) {
	userStr = userStr[:len(userStr)-feedStateLUT[feedState]]
	emitter := getFeed(ctx, userStr)
	if emitter == nil {
		notFound(ctx)
		return
	}
	switch feedState {
	case feedJSON:
		ctx.Header("Content-Type", "application/json; charset=utf-8")
		ctx.JSON(http.StatusOK, emitter.JSON())
	case feedRSS2:
		if d, err := emitter.RSS2().XML(); err != nil {
			log.Printf("error generating RSS 2.0 feed for user %s: %s", userStr, err)
			ctx.String(http.StatusInternalServerError, "Internal Server Error")
			return
		} else {
			ctx.Data(http.StatusOK, "application/rss+xml; charset=utf-8", d)
		}
	case feedAtom1:
		if d, err := emitter.Atom1().XML(); err != nil {
			log.Printf("error generating RSS 2.0 feed for user %s: %s", userStr, err)
			ctx.String(http.StatusInternalServerError, "Internal Server Error")
			return
		} else {
			ctx.Data(http.StatusOK, "application/atom+xml; charset=utf-8", d)
		}
	}
}
