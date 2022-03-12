package feed

import "time"

// TODO: xml tags

type Item struct {
	Title string    `json:"title"`
	ID    *string   `json:"id,omitempty"`
	Link  string    `json:"link"`
	Date  time.Time `json:"date"`

	Description *string    `json:"description,omitempty"`
	Content     *string    `json:"content,omitempty"`
	Category    []Category `json:"category,omitempty"`

	GUID *string `json:"guid,omitempty"`

	Image     *string    `json:"image,omitempty"`
	Audio     *string    `json:"audio,omitempty"`
	Video     *string    `json:"video,omitempty"`
	Enclosure *Enclosure `json:"enclosure,omitempty"`

	Author      []Author `json:"author,omitempty"`
	Contributor []Author `json:"contributor,omitempty"`

	Published time.Time `json:"published"`
	Copyright *string   `json:"copyright,omitempty"`

	Extensions []Extension `json:"extensions,omitempty"`
}

type Enclosure struct {
	URL      string  `json:"url"`
	Type     *string `json:"type,omitempty"`
	Length   *int    `json:"length,omitempty"`
	Title    *string `json:"title,omitempty"`
	Duration *int    `json:"duration,omitempty"`
}

type Author struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	Link  *string `json:"link,omitempty"`
}

type Category struct {
	Name   *string `json:"name,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Scheme *string `json:"scheme,omitempty"`
	Term   *string `json:"term,omitempty"`
}

type Options struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Updated   *time.Time `json:"updated,omitempty"`
	Generator *string    `json:"generator,omitempty"`
	Language  *string    `json:"language,omitempty"`
	TTL       *int       `json:"ttl,omitempty"`

	Feed      *string           `json:"feed,omitempty"`
	FeedLinks map[string]string `json:"feedLinks,omitempty"`
	Hub       *string           `json:"hub,omitempty"`
	Docs      *string           `json:"docs,omitempty"`

	Author      *Author `json:"author,omitempty"`
	Link        *string `json:"link,omitempty"`
	Description *string `json:"description,omitempty"`
	Image       *string `json:"image,omitempty"`
	Favicon     *string `json:"favicon,omitempty"`
	Copyright   string  `json:"copyright"`
}

type Extension struct {
	Name    string      `json:"name"`
	Objects interface{} `json:"objects"`
}
