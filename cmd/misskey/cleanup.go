package main

import (
	"context"
	"log"
	"time"

	"random.chars.jp/git/misskey/db"
)

func cleanup() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("error shutting down web server: %s", err)
	}

	if db.Cache != nil {
		if err := db.Cache.Close(); err != nil {
			log.Printf("error closing redis: %s", err)
		}
	}
	if err := db.Close(); err != nil {
		log.Printf("error closing database: %s", err)
	}
}
