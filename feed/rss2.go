package feed

import (
	"encoding/xml"
	"strings"
	"time"
)

const (
	RSS2FeedType         = "application/rss+xml"
	RSS2FeedDocs         = "https://validator.w3.org/feed/docs/rss2.html"
	RSS2FeedDC           = "http://purl.org/dc/elements/1.1/"
	RSS2FeedContent      = "http://purl.org/rss/1.0/modules/content/"
	RSS2FeedAtom         = "http://www.w3.org/2005/Atom"
	RSS2FeedTypeCatchAll = "application/octet-stream"
)

type RSS2Feed struct {
	XMLName bool `xml:"rss"`

	Version string `xml:"version,attr"`
	DC      string `xml:"xmlns:dc,attr,omitempty"`
	Content string `xml:"xmlns:content,attr,omitempty"`
	Atom    string `xml:"xmlns:atom,attr,omitempty"`

	Channel RSS2FeedChannel `xml:"channel"`

	isAtom    bool
	isContent bool
}

func (r *RSS2Feed) XML() ([]byte, error) {
	if data, err := xml.Marshal(r); err != nil {
		return nil, err
	} else {
		return append([]byte(xml.Header), data...), nil
	}
}

func (r *RSS2Feed) XMLIndent() ([]byte, error) {
	if data, err := xml.MarshalIndent(r, "", "    "); err != nil {
		return nil, err
	} else {
		return append([]byte(xml.Header), data...), nil
	}
}

type RSS2FeedChannel struct {
	Title         string  `xml:"title"`
	Link          string  `xml:"link,omitempty"`
	Description   *string `xml:"description,omitempty"`
	LastBuildDate string  `xml:"lastBuildDate"`
	Docs          string  `xml:"docs"`
	Generator     string  `xml:"generator"`

	Language  *string            `xml:"language,omitempty"`
	TTL       *int               `xml:"ttl,omitempty"`
	Image     *RSS2FeedImage     `xml:"image,omitempty"`
	Copyright *string            `xml:"copyright,omitempty"`
	Category  []string           `xml:"category,omitempty"`
	AtomLink  []RSS2FeedAtomLink `xml:"atom:link,omitempty"`

	Item []RSS2FeedItem `xml:"item,omitempty"`
}

type RSS2FeedImage struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Link  string `xml:"link,omitempty"`
}

type RSS2FeedAtomLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr,omitempty"`
}

type RSS2FeedItem struct {
	Title       *cdata_t           `xml:"title,omitempty"`
	Link        *string            `xml:"link,omitempty"`
	GUID        *string            `xml:"guid,omitempty"`
	PubDate     *string            `xml:"pubDate,omitempty"`
	Description *cdata_t           `xml:"description,omitempty"`
	Content     *cdata_t           `xml:"content:encoded,omitempty"`
	Author      []string           `xml:"author,omitempty"`
	Category    []RSS2FeedCategory `xml:"category,omitempty"`
	Enclosure   *RSS2FeedEnclosure `xml:"enclosure,omitempty"`
}

type RSS2FeedCategory struct {
	Domain *string `xml:"domain,attr,omitempty"`
	*string
}

type RSS2FeedEnclosure struct {
	URL    *string `xml:"url,attr,omitempty"`
	Length int     `xml:"length,attr"`
	Type   string  `xml:"type,attr"`

	FlattenedURL      *string `json:"url,omitempty"`
	FlattenedType     *string `json:"type,omitempty"`
	FlattenedLength   *int    `json:"length,omitempty"`
	FlattenedTitle    *string `json:"title,omitempty"`
	FlattenedDuration *int    `json:"duration,omitempty"`
}

// RSS2 emits an RSS 2.0 feed
func (e *Emitter) RSS2() *RSS2Feed {
	var feed = &RSS2Feed{
		Version: "2.0",

		Channel: RSS2FeedChannel{
			Title:       e.Options.Title,
			Link:        sanitize(e.Options.Link),
			Description: e.Options.Description,
		},

		isAtom:    false,
		isContent: false,
	}

	if e.Options.Updated != nil {
		feed.Channel.LastBuildDate = e.Options.Updated.UTC().Format(time.RFC1123)
	} else {
		feed.Channel.LastBuildDate = time.Now().UTC().Format(time.RFC1123)
	}
	if e.Options.Docs != nil {
		feed.Channel.Docs = *e.Options.Docs
	} else {
		feed.Channel.Docs = RSS2FeedDocs
	}
	if e.Options.Generator != nil {
		feed.Channel.Generator = *e.Options.Generator
	} else {
		feed.Channel.Generator = DefaultGenerator
	}

	feed.Channel.Language = e.Options.Language
	feed.Channel.TTL = e.Options.TTL
	if e.Options.Image != nil {
		feed.Channel.Image = &RSS2FeedImage{
			Title: e.Options.Title,
			URL:   *e.Options.Image,
			Link:  feed.Channel.Link,
		}
	}
	feed.Channel.Copyright = e.Options.Copyright
	feed.Channel.Category = e.Categories
	var atomLink *string
	if e.Options.Feed != nil {
		atomLink = e.Options.Feed
	} else if e.Options.FeedLinks != nil {
		if rss, ok := e.Options.FeedLinks["rss"]; ok {
			atomLink = &rss
		}
	}
	if atomLink != nil {
		feed.isAtom = true
		feed.Channel.AtomLink = append(feed.Channel.AtomLink, RSS2FeedAtomLink{
			Href: sanitize(atomLink),
			Rel:  "self",
			Type: RSS2FeedType,
		})
	}
	if e.Options.Hub != nil {
		feed.isAtom = true
		feed.Channel.AtomLink = append(feed.Channel.AtomLink, RSS2FeedAtomLink{
			Href: sanitize(e.Options.Hub),
			Rel:  "hub",
		})
	}
	if n := len(e.Items); n > 0 {
		feed.Channel.Item = make([]RSS2FeedItem, n)
		for i, item := range e.Items {
			feed.Channel.Item[i].Title = cdata(&item.Title)
			feed.Channel.Item[i].Link = copy_str_ptr(&item.Link)
			if item.GUID != nil {
				feed.Channel.Item[i].GUID = item.GUID
			} else if item.ID != nil {
				feed.Channel.Item[i].GUID = item.ID
			} else if item.Link != "" {
				str := sanitize(&item.Link)
				feed.Channel.Item[i].GUID = &str
			}
			if item.Published != nil {
				str := item.Published.UTC().Format(time.RFC1123)
				feed.Channel.Item[i].PubDate = &str
			} else if item.Date != nil {
				str := item.Date.UTC().Format(time.RFC1123)
				feed.Channel.Item[i].PubDate = &str
			}
			feed.Channel.Item[i].Description = cdata(item.Description)
			if item.Content != nil {
				feed.isContent = true
				feed.Channel.Item[i].Content = cdata(item.Content)
			}
			for _, author := range item.Author {
				if author.Email != nil && author.Name != nil {
					feed.Channel.Item[i].Author = append(feed.Channel.Item[i].Author, *author.Email+" ("+*author.Name+")")
				}
			}
			if m := len(item.Category); m > 0 {
				feed.Channel.Item[i].Category = make([]RSS2FeedCategory, m)
				for j, category := range item.Category {
					feed.Channel.Item[i].Category[j] = RSS2FeedCategory{
						Domain: category.Domain,
						string: category.Name,
					}
				}
			}
			if item.Video != nil {
				feed.Channel.Item[i].Enclosure = &RSS2FeedEnclosure{
					URL:    item.Video,
					Length: 0,
					Type:   guessType("video", item.Video),
				}
			} else if item.Audio != nil {
				feed.Channel.Item[i].Enclosure = &RSS2FeedEnclosure{
					URL:    item.Audio,
					Length: 0,
					Type:   guessType("audio", item.Audio),
				}
			} else if item.Image != nil {
				feed.Channel.Item[i].Enclosure = &RSS2FeedEnclosure{
					URL:    item.Image,
					Length: 0,
					Type:   guessType("image", item.Image),
				}
			} else if item.Enclosure != nil {
				feed.Channel.Item[i].Enclosure = &RSS2FeedEnclosure{
					Length: 0,
					// original library had zero value here which is "image"
					Type: guessType("image", copy_str_ptr(&item.Enclosure.URL)),

					FlattenedURL:      copy_str_ptr(&item.Enclosure.URL),
					FlattenedType:     item.Enclosure.Type,
					FlattenedLength:   item.Enclosure.Length,
					FlattenedTitle:    item.Enclosure.Title,
					FlattenedDuration: item.Enclosure.Duration,
				}
			}
		}
	}

	if feed.isContent {
		feed.DC = RSS2FeedDC
		feed.Content = RSS2FeedContent
	}
	if feed.isAtom {
		feed.Atom = RSS2FeedAtom
	}

	return feed
}

func guessType(cat string, str *string) string {
	s := strings.Split(*str, ".")
	if len(s) > 0 {
		return cat + s[len(s)-1]
	}
	return RSS2FeedTypeCatchAll
}
