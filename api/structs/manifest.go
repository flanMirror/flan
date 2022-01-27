package structs

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

func NewManifest(name string) Manifest {
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

	if name != "" {
		m.ShortName = name
		m.Name = name
	}
	return m
}
