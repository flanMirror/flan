package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const referenceBackendURLString = "http://localhost:3001"

var (
	referenceBackendURL          *url.URL
	referenceBackendReverseProxy *httputil.ReverseProxy
)

func init() {
	upstream := referenceBackendURLString
	if u := os.Getenv("UPSTREAM"); u != "" {
		upstream = u
	}
	if u, err := url.Parse(upstream); err != nil {
		panic("error parsing reference URL: " + err.Error())
	} else {
		referenceBackendURL = u
	}
	referenceBackendReverseProxy = httputil.NewSingleHostReverseProxy(referenceBackendURL)
	log.Printf("proxying unimplemented routes to %s", upstream)
}

// placeholder is the function called by unimplemented API endpoints
func placeholder(ctx Context) {
	if ctx.GetHeader("X-Forwarded-For") != "" {
		ctx.String(http.StatusBadRequest, "X-Forwarded-For is set")
		return
	}

	ctx.(*ginContext).Context.Set("proxy", true)
	referenceBackendReverseProxy.ServeHTTP(ctx.Writer(), ctx.Request())
	ctx.Abort()
}

// Placeholder is the function called by unimplemented web routes
func Placeholder(ctx *gin.Context) {
	placeholder(&ginContext{Context: ctx})
}
