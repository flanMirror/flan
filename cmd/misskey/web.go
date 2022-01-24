package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	openapi "random.chars.jp/git/misskey/api"
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
	apiGroup := router.Group("/api/")
	openapi.RegisterRoutes(apiGroup)

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

func prepareAssets() (fs.FS, []byte) {
	if public, err := fs.Sub(assets, "assets/public"); err != nil {
		log.Print("compiled without static assets, not serving")
		return nil, nil
	} else {
		var (
			indexHTML []byte
			file      fs.File
		)
		if file, err = public.Open("index.html"); err != nil {
			log.Print("bundled assets do not have index.html")
			return public, nil
		} else {
			if indexHTML, err = io.ReadAll(file); err != nil {
				log.Fatalf("error reading bundled index.html: %s", err)
			}
			if err = file.Close(); err != nil {
				log.Printf("error closing bundled index.html: %s", err)
			}
			return public, indexHTML
		}
	}
}

func serveStatic() func(context *gin.Context) {
	public, indexHTML := prepareAssets()
	if public != nil && indexHTML != nil {
		return func(context *gin.Context) {
			if context.Request.Method != http.MethodPost {
				context.String(http.StatusNotFound, "Not Found")
				return
			}
			p := context.Request.URL.Path
			if !strings.HasPrefix(p, "/") {
				p = "/" + p
			}
			p = path.Clean(p)

			if s, err := fs.Stat(public, p); err != nil || s.IsDir() {
				context.Data(http.StatusOK, gin.MIMEHTML, indexHTML)
				return
			}

			context.FileFromFS(p, http.FS(public))
		}
	} else {
		//return func(context *gin.Context) {
		//	context.Data(http.StatusNoContent, gin.MIMEHTML, nil)
		//}
		return func(context *gin.Context) {
			// FIXME
			openapi.Placeholder(context)
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
