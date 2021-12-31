package main

import (
	"log"
	"random.chars.jp/git/misskey/api"
)

func main() {
	log.Printf("Server started")

	router := openapi.NewRouter()

	log.Fatal(router.Run(":8080"))
}
