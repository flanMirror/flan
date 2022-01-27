package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"random.chars.jp/git/misskey/api/payload"
)

func routesSetup() {
	router.GET("/manifest.json", setNoFrame, func(context *gin.Context) {
		context.Header("Cache-Control", "max-age=300")
		context.Data(http.StatusOK, gin.MIMEJSON, payload.Manifest.Data())
	})
}
