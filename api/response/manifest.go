//go:build !bugforbug

package response

import "random.chars.jp/git/misskey/db/orm"

// this fixes the nonsensical behaviour where icon does not change

type Manifest struct {
	ShortName       string              `json:"short_name"`
	Name            string              `json:"name"`
	StartURL        string              `json:"start_url"`
	Display         string              `json:"display"`
	BackgroundColor string              `json:"background_color"`
	ThemeColor      string              `json:"theme_color"`
	Icons           []ManifestIcon      `json:"icons"`
	ShareTarget     ManifestShareTarget `json:"share_target"`
}

type ManifestIcon struct {
	Src   string `json:"src"`
	Sizes string `json:"sizes"`
	Type  string `json:"type"`
}

type ManifestShareTarget struct {
	Action string                    `json:"action"`
	Params ManifestShareTargetParams `json:"params"`
}

type ManifestShareTargetParams struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	URL   string `json:"url"`
}

func NewManifest(metum *orm.Metum) Manifest {
	m := Manifest{
		ShortName:       "Misskey",
		Name:            "Misskey",
		StartURL:        "/",
		Display:         "standalone",
		BackgroundColor: "#313a42",
		ThemeColor:      "#86b300",
		Icons: []ManifestIcon{
			{
				Src:   "/static-assets/icons/192.png",
				Sizes: "192x192",
				Type:  "image/png",
			},
			{
				Src:   "/static-assets/icons/512.png",
				Sizes: "512x512",
				Type:  "image/png",
			},
		},
		ShareTarget: ManifestShareTarget{
			Action: "/share/",
			Params: ManifestShareTargetParams{
				Title: "title",
				Text:  "text",
				URL:   "url",
			},
		},
	}

	// FIXME: new version added colours, please remember that when catching up!

	if metum.Name.Valid {
		m.ShortName = metum.Name.String
		m.Name = metum.Name.String
	}
	if metum.IconUrl.Valid {
		// FIXME: scale the images
		m.Icons[0].Src = metum.IconUrl.String
		m.Icons[1].Src = metum.IconUrl.String
	}
	return m
}
