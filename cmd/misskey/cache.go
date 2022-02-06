package main

import (
	"log"
)

const (
	apiDocPath = "assets/public/static-assets/redoc.html"
)

var (
	apiDoc []byte
)

func init() {
	if d, err := assets.ReadFile(apiDocPath); err != nil {
		log.Printf("error caching asset %s: %s", apiDocPath, err)
	} else {
		apiDoc = d
	}
}
