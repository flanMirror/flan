package payload

import (
	"random.chars.jp/git/misskey/api/response"
	"random.chars.jp/git/misskey/data"
	"random.chars.jp/git/misskey/db"
	"random.chars.jp/git/misskey/db/orm"
)

var Manifest = data.New[response.Manifest]()

func init() {
	db.Meta.Register(func(metum *orm.Metum) {
		ManifestUpdate(metum.Name.String)
	})
}

// ManifestUpdate updates Manifest payload with new name.
// Must be called after instance name change.
func ManifestUpdate(name string) {
	Manifest.Set(response.NewManifest(name))
}
