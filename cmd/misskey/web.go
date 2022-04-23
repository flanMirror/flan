package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"random.chars.jp/git/misskey/api"
	"random.chars.jp/git/misskey/config"
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
		router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			if param.Latency > time.Minute {
				param.Latency = param.Latency - param.Latency%time.Second
			}

			prefix := ""
			if param.Keys != nil {
				if _, ok := param.Keys["proxy"]; ok {
					prefix = "(proxy) "
				}
			}

			return fmt.Sprintf("%s %3d | %-5s | %13v | %15s | %s%#v\n%s",
				param.TimeStamp.Format("2006/01/02 15:04:05"),
				param.StatusCode,
				param.Method,
				param.Latency,
				param.ClientIP,
				prefix, param.Path,
				param.ErrorMessage,
			)
		}))
		router.Use(gin.Recovery())
	} else {
		router.Use(recovery())
	}

	if templates, err := template.
		New("all built-in templates").
		Funcs(template.FuncMap{"str": func(str string) template.HTML { return template.HTML(str) }}).
		ParseFS(assets, "assets/template/*.tmpl"); err != nil {
		log.Fatalf("error parsing built-in templates: %s", err)
	} else {
		log.Printf("%d built-in templates present", len(templates.Templates()))
		router.SetHTMLTemplate(templates)
		templateSetup(templates)
	}

	if config.Web.HSTS && config.HTTPS {
		router.Use(func(context *gin.Context) {
			context.Header("strict-transport-security", "max-age=15552000; preload")
		})
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
		return func(ctx *gin.Context) {
			if ctx.Request.Method != http.MethodGet {
				notFound(ctx)
				return
			}

			p := ctx.Request.URL.Path
			p = strings.TrimPrefix(p, "/")
			p = path.Clean(p)

			if s, err := fs.Stat(public, p); err != nil || s.IsDir() {
				serveBase(ctx)
				return
			}

			setNoFrame(ctx)

			// headers for specific paths
			if strings.HasPrefix(p, "/assets/") ||
				strings.HasPrefix(p, "/static-assets/") ||
				strings.HasPrefix(p, "/client-assets/") {
				// 7 days
				ctx.Header("Cache-Control", "max-age=604800")
			}
			if strings.HasPrefix(p, "/twemoji/") {
				// 30 days
				ctx.Header("Content-Security-Policy", "default-src 'none'; style-src 'unsafe-inline'")
				ctx.Header("Cache-Control", "max-age=2592000")
			}

			ctx.FileFromFS(p, http.FS(public))
		}
	} else {
		log.Print("no public bundled, not serving")
		return func(context *gin.Context) {
			context.Data(http.StatusNoContent, gin.MIMEHTML, nil)
		}
	}
}

func serve() {
	if config.Web.FastCGI {
		// serve using fcgi
		if err := fcgi.Serve(listener, server.Handler); err != nil {
			log.Printf("error serve: %s", err)
		}
	} else {
		// serve using http server
		if err := server.Serve(listener); err == http.ErrServerClosed {
			log.Print("web server closed")
		} else {
			log.Printf("error serve: %s", err)
		}
	}
}

func setNoFrame(context *gin.Context) {
	context.Header("X-Frame-Options", "DENY")
}
