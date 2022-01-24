package main

import (
	"context"
	"log"
	"random.chars.jp/git/misskey/db"
	"time"
)

func cleanup() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("error shutting down web server: %s", err)
	}

	if err := db.Close(); err != nil {
		log.Printf("error closing database: %s", err)
	}
}
