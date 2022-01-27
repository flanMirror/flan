package payload

import (
	"random.chars.jp/git/misskey/api/structs"
	"random.chars.jp/git/misskey/data"
)

var Manifest = data.New()

// ManifestUpdate updates Manifest payload with new name.
// Must be called after instance name change.
func ManifestUpdate(name string) {
	Manifest.Set(structs.NewManifest(name))
}
