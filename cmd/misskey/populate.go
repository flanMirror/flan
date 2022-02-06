package main

import (
	"random.chars.jp/git/misskey/db"
)

// populate is called once database connection is established.
// it populates some payloads that are relatively constant.
func populate() {
	// expire some eager loading expiring items to populate structs behind them

	// order is important for these initial expirations
	// see db/meta.go for further explanation
	db.LocalUserCount.Expire()
	db.Meta.Expire()
	db.Ads.Expire()
	db.Emojis.Expire()
	db.ProxyAccount.Expire()
}
