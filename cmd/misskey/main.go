package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"random.chars.jp/git/misskey/config"
	"syscall"
)

func main() {
	flag.Parse()
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
