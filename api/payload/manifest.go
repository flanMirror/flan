package payload

import (
	"random.chars.jp/git/misskey/api/structs"
	"random.chars.jp/git/misskey/data"
	"random.chars.jp/git/misskey/db"
	"random.chars.jp/git/misskey/db/models"
)

var Manifest = data.New()

func init() {
	db.Meta.Register(func(data interface{}) {
		ManifestUpdate(data.(*models.Metum).Name.String)
	})
}

// ManifestUpdate updates Manifest payload with new name.
// Must be called after instance name change.
func ManifestUpdate(name string) {
	Manifest.Set(structs.NewManifest(name))
}
