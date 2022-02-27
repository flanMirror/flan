package main

import (
	"log"

	"random.chars.jp/git/misskey/data"
)

const (
	apiDocPath = "assets/public/static-assets/redoc.html"
)

var (
	apiDoc []byte
	base   data.P
)

func init() {
	if d, err := assets.ReadFile(apiDocPath); err != nil {
		log.Printf("error caching asset %s: %s", apiDocPath, err)
	} else {
		apiDoc = d
	}
}
