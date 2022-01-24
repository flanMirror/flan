package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	openapi "random.chars.jp/git/misskey/api"
	"random.chars.jp/git/misskey/config"
	"syscall"
)

const banner = `
  _____ _         _           
 |     |_|___ ___| |_ ___ _ _ 
 | | | | |_ -|_ -| '_| -_| | |
 |_|_|_|_|___|___|_,_|___|_  |
             ...but fast |___|

 Misskey-fast is a go implementation of Misskey %s.
 If you like Misskey, please donate to support development. https://www.patreon.com/syuilo

--- %s (PID: %d) ---
`

var target bool

func init() {
	flag.BoolVar(&target, "target", false, "display targeted misskey version")
}

func main() {
	flag.Parse()
	if target {
		fmt.Print(openapi.Target)
		os.Exit(0)
	}

	doBanner()

	if config.Parse() {
		log.Print("configuration load complete")
	} else {
		log.Print("legacy configuration file loaded, please consider converting to new format")
	}
	webSetup()
	// TODO: database and stuff

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
			fmt.Printf(banner, openapi.Target, hostname, os.Getpid())
		}
	}
}
