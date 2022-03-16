package feed

import (
	"encoding/xml"
	"time"
)

const (
	Atom1FeedNamespace = "http://www.w3.org/2005/Atom"
)

type Atom1Feed struct {
	XMLName   bool   `xml:"feed"`
	Namespace string `xml:"xmlns,attr"`
	ID        string `xml:"id"`
	Title     string `xml:"title"`
	Updated   string `xml:"updated"`
	Generator string `xml:"generator"`

	Author      *Atom1FeedAuthor       `xml:"author,omitempty"`
	Link        []Atom1FeedLink        `xml:"link,omitempty"`
	Subtitle    *string                `xml:"subtitle,omitempty"`
	Logo        *string                `xml:"logo,omitempty"`
	Icon        *string                `xml:"icon,omitempty"`
	Rights      *string                `xml:"rights,omitempty"`
	Category    []Atom1FeedCategory    `xml:"category,omitempty"`
	Contributor []Atom1FeedContributor `xml:"contributor,omitempty"`

	Entry []Atom1FeedEntry `xml:"entry,omitempty"`
}

func (a *Atom1Feed) XML() ([]byte, error) {
	if data, err := xml.Marshal(a); err != nil {
		return nil, err
	} else {
		return append([]byte(xml.Header), data...), nil
	}
}

func (a *Atom1Feed) XMLIndent() ([]byte, error) {
	if data, err := xml.MarshalIndent(a, "", "    "); err != nil {
		return nil, err
	} else {
		return append([]byte(xml.Header), data...), nil
	}
}

type Atom1FeedAuthor struct {
	Name  *string `xml:"name,omitempty"`
	Email *string `xml:"email,omitempty"`
	URI   string  `xml:"uri,omitempty"`
}

type Atom1FeedLink struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

type Atom1FeedCategory struct {
	Label  *string `xml:"label,attr,omitempty"`
	Scheme *string `xml:"scheme,attr,omitempty"`
	Term   *string `xml:"term,attr,omitempty"`
}

type Atom1FeedContributor struct {
	Name  *string `xml:"name,omitempty"`
	Email *string `xml:"email,omitempty"`
	URI   string  `xml:"uri,omitempty"`
}

type Atom1FeedEntry struct {
	Title   Atom1FeedEntryCDATAWithType `xml:"title"`
	ID      string                      `xml:"id"`
	Link    []Atom1FeedEntryLink        `xml:"link"`
	Updated string                      `xml:"updated,omitempty"`

	Summary     *Atom1FeedEntryCDATAWithType `xml:"summary,omitempty"`
	Content     *Atom1FeedEntryCDATAWithType `xml:"content,omitempty"`
	Author      []Atom1FeedAuthor            `xml:"author,omitempty"`
	Contributor []Atom1FeedAuthor            `xml:"contributor,omitempty"`
	Category    []Atom1FeedCategory          `xml:"category,omitempty"`
	Published   *string                      `xml:"published,omitempty"`
	Rights      *string                      `xml:"rights,omitempty"`
}

type Atom1FeedEntryCDATAWithType struct {
	Type string `xml:"type,attr"`

	cdata_t
}

type Atom1FeedEntryLink struct {
	Href string `xml:"href,attr"`
}

// Atom1 emits an Atom 1.0 feed
func (e *Emitter) Atom1() *Atom1Feed {
	feed := &Atom1Feed{
		Namespace: Atom1FeedNamespace,
		ID:        e.Options.ID,
		Title:     e.Options.Title,
	}
	if e.Options.Updated != nil {
		feed.Updated = e.Options.Updated.UTC().Format(RFC3339JavaScript)
	} else {
		feed.Updated = time.Now().Format(RFC3339JavaScript)
	}
	if e.Options.Generator != nil {
		feed.Generator = sanitize(e.Options.Generator)
	} else {
		feed.Generator = DefaultGenerator
	}

	if e.Options.Author != nil {
		feed.Author = &Atom1FeedAuthor{
			Name:  e.Options.Author.Name,
			Email: e.Options.Author.Email,
			URI:   sanitize(e.Options.Author.Link),
		}
	}

	if e.Options.Link != nil {
		feed.Link = append(feed.Link, Atom1FeedLink{
			Rel:  "alternate",
			Href: sanitize(e.Options.Link),
		})
	}
	if e.Options.Feed != nil {
		feed.Link = append(feed.Link, Atom1FeedLink{
			Rel:  "self",
			Href: sanitize(e.Options.Feed),
		})
	} else if e.Options.FeedLinks != nil {
		if atom, ok := e.Options.FeedLinks["atom"]; ok {
			feed.Link = append(feed.Link, Atom1FeedLink{
				Rel:  "self",
				Href: sanitize(&atom),
			})
		}
	} else if e.Options.Hub != nil {
		feed.Link = append(feed.Link, Atom1FeedLink{
			Rel:  "hub",
			Href: sanitize(e.Options.Hub),
		})
	}

	feed.Subtitle = e.Options.Description
	feed.Logo = e.Options.Image
	feed.Icon = e.Options.Favicon
	feed.Rights = e.Options.Copyright

	if n := len(e.Categories); n > 0 {
		feed.Category = make([]Atom1FeedCategory, n)
		for i, category := range e.Categories {
			feed.Category[i] = Atom1FeedCategory{
				Term: copy_str_ptr(&category),
			}
		}
	}

	if n := len(e.Contributors); n > 0 {
		feed.Contributor = make([]Atom1FeedContributor, n)
		for i, contributor := range e.Contributors {
			feed.Contributor[i] = Atom1FeedContributor{
				Name:  contributor.Name,
				Email: contributor.Email,
				URI:   sanitize(contributor.Link),
			}
		}
	}

	if n := len(e.Items); n > 0 {
		feed.Entry = make([]Atom1FeedEntry, n)
		for i, item := range e.Items {
			feed.Entry[i] = Atom1FeedEntry{
				Title: Atom1FeedEntryCDATAWithType{
					Type: "html",
					cdata_t: cdata_t{
						Text: item.Title,
					},
				},
				Link: []Atom1FeedEntryLink{
					{
						Href: sanitize(&item.Link),
					},
				},
			}

			if item.ID != nil {
				feed.Entry[i].ID = sanitize(item.ID)
			} else {
				feed.Entry[i].ID = sanitize(&item.Link)
			}

			if item.Date != nil {
				feed.Entry[i].Updated = item.Date.Format(RFC3339JavaScript)
			}

			if item.Description != nil {
				feed.Entry[i].Summary = &Atom1FeedEntryCDATAWithType{
					Type: "html",
					cdata_t: cdata_t{
						Text: *item.Description,
					},
				}
			}

			if item.Content != nil {
				feed.Entry[i].Content = &Atom1FeedEntryCDATAWithType{
					Type: "html",
					cdata_t: cdata_t{
						Text: *item.Content,
					},
				}
			}

			if m := len(item.Author); m > 0 {
				feed.Entry[i].Author = make([]Atom1FeedAuthor, m)
				for j, author := range item.Author {
					feed.Entry[i].Author[j] = Atom1FeedAuthor{
						Name:  author.Name,
						Email: author.Email,
						URI:   sanitize(author.Link),
					}
				}
			}

			if m := len(item.Contributor); m > 0 {
				feed.Entry[i].Contributor = make([]Atom1FeedAuthor, m)
				for j, contributor := range item.Contributor {
					feed.Entry[i].Contributor[j] = Atom1FeedAuthor{
						Name:  contributor.Name,
						Email: contributor.Email,
						URI:   sanitize(contributor.Link),
					}
				}
			}

			if m := len(item.Category); m > 0 {
				feed.Entry[i].Category = make([]Atom1FeedCategory, m)
				for j, category := range item.Category {
					feed.Entry[i].Category[j] = Atom1FeedCategory{
						Label:  category.Name,
						Scheme: category.Scheme,
						Term:   category.Term,
					}
				}
			}

			if item.Published != nil {
				str := item.Published.Format(RFC3339JavaScript)
				feed.Entry[i].Published = &str
			}

			feed.Entry[i].Rights = item.Copyright
		}
	}

	return feed
}
