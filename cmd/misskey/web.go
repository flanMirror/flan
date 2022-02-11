package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"random.chars.jp/git/misskey/api"
	"random.chars.jp/git/misskey/config"
	"strconv"
	"strings"
)

//go:embed assets
var assets embed.FS

var (
	router   *gin.Engine
	listener net.Listener
	server   = http.Server{}
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func webSetup() {
	if d, ok := os.LookupEnv("GIN_DEBUG"); ok {
		if ginDebug, err := strconv.ParseBool(d); err == nil {
			if ginDebug {
				gin.SetMode(gin.DebugMode)
			}
		}
	}

	router = gin.New()
	router.ForwardedByClientIP = config.Web.ForwardedByClientIP
	if err := router.SetTrustedProxies(config.Web.TrustedProxies); err != nil {
		log.Fatalf("error setting trusted proxies: %s", err)
	}
	if config.Log.Verbose {
		router.Use(gin.Logger())
		router.Use(gin.Recovery())
	} else {
		router.Use(recovery())
	}
	router.NoRoute(serveStatic())
	routesSetup()
	apiGroup := router.Group("/api/")
	api.RegisterRoutes(apiGroup)

	if config.Web.Port < 0 {
		if l, err := net.Listen("unix", config.Web.Host); err != nil {
			log.Fatalf("error listening on socket: %s", err)
		} else {
			listener = l
		}

		if err := os.Chmod(config.Web.Host, 0777); err != nil {
			log.Printf("error chmod: %s", err)
		}
	} else {
		if l, err := net.Listen("tcp", config.Web.Host+":"+strconv.Itoa(config.Web.Port)); err != nil {
			log.Fatalf("error listening on TCP port: %s", err)
		} else {
			listener = l
		}

		log.Printf("web server listening on %s:%d", config.Web.Host, config.Web.Port)
	}
	server.Handler = router
}

func prepareAssets() fs.FS {
	if public, err := fs.Sub(assets, "assets/public"); err != nil {
		log.Print("compiled without static assets, not serving")
		return nil
	} else {
		return public
	}
}

func serveStatic() func(context *gin.Context) {
	if public := prepareAssets(); public != nil {
		return func(context *gin.Context) {
			if context.Request.Method != http.MethodGet {
				context.String(http.StatusNotFound, "Not Found")
				return
			}

			// TODO: set cache

			p := context.Request.URL.Path
			if !strings.HasPrefix(p, "/") {
				p = "/" + p
			}
			p = path.Clean(p)

			if s, err := fs.Stat(public, p); err != nil || s.IsDir() {
				api.Placeholder(context)
				return
			}

			setNoFrame(context)
			context.FileFromFS(p, http.FS(public))
		}
	} else {
		//return func(context *gin.Context) {
		//	context.Data(http.StatusNoContent, gin.MIMEHTML, nil)
		//}
		log.Print("got nil public, starting reverse proxy catch-all")
		return func(context *gin.Context) {
			// FIXME
			api.Placeholder(context)
		}
	}
}

func serve() {
	if err := server.Serve(listener); err == http.ErrServerClosed {
		log.Printf("web server closed")
	} else {
		log.Printf("error serve: %s", err)
	}
}

func setNoFrame(context *gin.Context) {
	context.Header("X-Frame-Options", "DENY")
}
