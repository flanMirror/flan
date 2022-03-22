package main

import (
	"random.chars.jp/git/misskey/db/cache"
)

// populateDBCache is called once database connection is established
// it populates some payload caches that are unchanging for the most part
func populateDBCache() {
	// expire some eager loading expiring items to populate structs behind them

	// order is important for these initial expirations
	// see db/meta.go for further explanation
	cache.LocalUserCount.Expire()
	cache.Meta.Expire()
	cache.Ads.Expire()
	cache.Emojis.Expire()
	cache.ProxyAccount.Expire()
}
