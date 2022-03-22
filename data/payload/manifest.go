package payload

import (
	"random.chars.jp/git/misskey/api/response"
	"random.chars.jp/git/misskey/data"
	"random.chars.jp/git/misskey/db/cache"
	"random.chars.jp/git/misskey/db/orm"
)

// Manifest caches the response of /manifest.json
// and is updated when
// db.Meta is invalidated, using the ManifestUpdate function
var Manifest = data.New[response.Manifest]()

func init() {
	cache.Meta.Register(ManifestUpdate)
}

// ManifestUpdate updates Manifest payload with new name.
// Must be called after instance name change.
func ManifestUpdate(metum *orm.Metum) {
	Manifest.Set(response.NewManifest(metum))
}
