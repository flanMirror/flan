package main

import (
	"context"
	"log"
	"random.chars.jp/git/misskey/api/payload"
	"random.chars.jp/git/misskey/db"
)

// populate is called once database connection is established.
// it populates some payloads that are relatively constant.
func populate() {
	if meta, err := db.FetchMeta(context.Background()); err != nil {
		log.Printf("error fetching meta: %s", err)
	} else {
		if meta.Name.Valid {
			payload.ManifestUpdate(meta.Name.String)
		} else {
			payload.ManifestUpdate("")
		}
	}
}
