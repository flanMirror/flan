package main

import (
	"context"
	"log"
	"time"
)

func cleanup() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("error shutting down web server: %s", err)
	}

	// TODO: other cleanup stuff like database
}
