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
	config.Parse()
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
