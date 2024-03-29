package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"random.chars.jp/git/misskey/config"
	"random.chars.jp/git/misskey/data/payload"
	"random.chars.jp/git/misskey/db"
	"random.chars.jp/git/misskey/db/redis"
	"random.chars.jp/git/misskey/spec"
)

const banner = `
  _____ _         _           
 |     |_|___ ___| |_ ___ _ _ 
 | | | | |_ -|_ -| '_| -_| | |
 |_|_|_|_|___|___|_,_|___|_  |
             ...but fast |___|

 Misskey-fast is a Go implementation of Misskey %s.
 If you like Misskey, please donate to support development. https://www.patreon.com/syuilo

--- %s (PID: %d) ---
`

var noBanner bool

func init() {
	flag.BoolVar(&noBanner, "nobanner", false, "disable displaying of the banner")
}

func main() {
	flag.Parse()

	if !noBanner {
		doBanner()
	}

	if config.Parse() {
		log.Print("configuration load complete")
	} else {
		log.Print("legacy configuration file loaded, please consider converting to new format")
	}

	redis.Instance = redis.New()
	if info, err := redis.Instance.Info(context.Background(), "server").Result(); err != nil {
		log.Fatalf("error getting redis info: %s", err)
	} else {
		if f := strings.SplitN(info, "redis_version:", 2); len(f) == 2 {
			if f = strings.SplitN(f[1], "\r\n", 2); len(f) == 2 {
				payload.Redis = f[0]
				log.Printf("connected to Redis version %s", f[0])
			} else {
				log.Fatal("error getting Redis line ending")
			}
		} else {
			log.Fatal("error getting redis_version field")
		}
	}
	if err := db.Open(); err != nil {
		log.Fatalf("error opening database: %s", err)
	} else {
		var version string
		if err = db.DB.QueryRow("SHOW server_version").Scan(&version); err != nil {
			log.Fatalf("error getting database version: %s", err)
		} else {
			payload.PSQL = version
			log.Printf("connected to PostgreSQL version %s", version)
		}
	}
	webSetup()
	populateDBCache()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		defer func() { cleanup() }()
		for {
			s := <-sig
			switch s {
			case os.Interrupt:
				fmt.Println()
				fallthrough
			default:
				log.Print("shutting down")
				return
			}
		}
	}()

	serve()
}

func doBanner() {
	if stat, err := os.Stdin.Stat(); err == nil {
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			var hostname string
			if hostname, err = os.Hostname(); err != nil {
				return
			}
			fmt.Printf(banner, spec.Target, hostname, os.Getpid())
		}
	}
}
