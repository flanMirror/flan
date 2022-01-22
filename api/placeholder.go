package openapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const referenceBackendURLString = "http://localhost:3001"

var (
	referenceBackendURL          *url.URL
	referenceBackendReverseProxy *httputil.ReverseProxy
)

func init() {
	if u, err := url.Parse(referenceBackendURLString); err != nil {
		panic("error parsing reference URL: " + err.Error())
	} else {
		referenceBackendURL = u
	}
	referenceBackendReverseProxy = httputil.NewSingleHostReverseProxy(referenceBackendURL)
}

// placeholder is the function called by unimplemented API endpoints
func placeholder(context *gin.Context) {
	if context.GetHeader("X-Forwarded-For") != "" {
		context.String(http.StatusBadRequest, "X-Forwarded-For is set")
		return
	}
	//if f, ok := context.Request.Header["X-Forwarded-For"]; ok {
	//	delete(context.Request.Header, "X-Forwarded-For")
	//	log.Printf("deleted X-Forwarded-For in placeholder proxying from %s with content %s",
	//		context.ClientIP(), f)
	//}
	referenceBackendReverseProxy.ServeHTTP(context.Writer, context.Request)
	context.Abort()
}
