package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"random.chars.jp/git/misskey/api/payload"
	"random.chars.jp/git/misskey/spec"
)

func routesSetup() {
	// TODO: HSTS

	router.GET("/manifest.json", setNoFrame, func(context *gin.Context) {
		context.Header("Cache-Control", "max-age=300")
		context.Data(http.StatusOK, gin.MIMEJSON, payload.Manifest.Data())
	})
	router.GET("/api-doc", setNoFrame, func(context *gin.Context) {
		context.Header("Cache-Control", "max-age=0")
		context.Data(http.StatusOK, gin.MIMEHTML, apiDoc)
	})
	router.GET("/api.json", setNoFrame, func(context *gin.Context) {
		context.Data(http.StatusOK, gin.MIMEJSON, spec.JSON())
	})

	router.GET("/streaming", func(context *gin.Context) {
		context.Header("Cache-Control", "private, max-age=0")
		context.Data(http.StatusServiceUnavailable,
			"text/plain; charset=utf-8",
			[]byte("Service Unavailable"))
	})
}
