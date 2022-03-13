package feed

import (
	json "github.com/json-iterator/go"
)

const RFC3339JavaScript = "2006-01-02T15:04:05.000Z07:00"

type JSONFeed struct {
	Version string `json:"version"`
	Title   string `json:"title"`

	HomePageURL *string         `json:"home_page_url,omitempty"`
	FeedURL     *string         `json:"feed_url,omitempty"`
	Description *string         `json:"description,omitempty"`
	Icon        *string         `json:"icon,omitempty"`
	Author      *JSONFeedAuthor `json:"author,omitempty"`
	Items       []JSONFeedItem  `json:"items,omitempty"`

	// xattr accommodates extensions, will be handled in MarshalJSON method
	xattr map[string]interface{}
}

type jsonFeed JSONFeed

func (f *JSONFeed) MarshalJSON() ([]byte, error) {
	if len(f.xattr) == 0 {
		return json.Marshal((*jsonFeed)(f))
	}

	// ugly!!

	// only allocate non-omitempty fields
	// because we don't know how many omitempty fields present or how many overlap in xattr
	m := make(map[string]interface{}, 2)
	m["version"] = f.Version
	m["title"] = f.Title
	if f.HomePageURL != nil {
		m["home_page_url"] = f.HomePageURL
	}
	if f.FeedURL != nil {
		m["feed_url"] = f.FeedURL
	}
	if f.Description != nil {
		m["description"] = f.Description
	}
	if f.Icon != nil {
		m["icon"] = f.Icon
	}
	if f.Author != nil {
		m["author"] = f.Author
	}
	if f.Items != nil {
		m["items"] = f.Items
	}
	for key, value := range f.xattr {
		m[key] = value
	}
	return json.Marshal(m)
}

type JSONFeedAuthor struct {
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}

type JSONFeedItem struct {
	ID          *string `json:"id,omitempty"`
	ContentHTML *string `json:"content_html,omitempty"`

	URL           string          `json:"url"`
	Title         string          `json:"title"`
	Summary       *string         `json:"summary,omitempty"`
	Image         *string         `json:"image,omitempty"`
	DateModified  string          `json:"date_modified,omitempty"`
	DatePublished string          `json:"date_published,omitempty"`
	Author        *JSONFeedAuthor `json:"author,omitempty"`
	Tags          []string        `json:"tags,omitempty"`

	// xattr accommodates extensions, will be handled in MarshalJSON method
	xattr map[string]interface{}
}

type jsonFeedItem JSONFeedItem

func (t *JSONFeedItem) MarshalJSON() ([]byte, error) {
	if len(t.xattr) == 0 {
		return json.Marshal((*jsonFeedItem)(t))
	}

	// ugly!!

	// only allocate non-omitempty fields
	// because we don't know how many omitempty fields present or how many overlap in xattr
	m := make(map[string]interface{}, 2)
	if t.ID != nil {
		m["id"] = t.ID
	}
	if t.ContentHTML != nil {
		m["content_html"] = t.ContentHTML
	}
	m["url"] = t.URL
	m["title"] = t.Title
	if t.Summary != nil {
		m["summary"] = t.Summary
	}
	if t.Image != nil {
		m["image"] = t.Image
	}
	if t.DateModified != "" {
		m["date_modified"] = t.DateModified
	}
	if t.DatePublished != "" {
		m["date_published"] = t.DatePublished
	}
	if t.Author != nil {
		m["author"] = t.Author
	}
	if t.Tags != nil {
		m["tags"] = t.Tags
	}
	for key, value := range t.xattr {
		m[key] = value
	}
	return json.Marshal(m)
}

const JSONFeedVersion = "https://jsonfeed.org/version/1"

// JSON emits a JSON feed
func (e *Emitter) JSON() *JSONFeed {
	var feed = &JSONFeed{
		Version: JSONFeedVersion,
		Title:   e.Options.Title,
	}

	feed.HomePageURL = e.Options.Link
	if e.Options.FeedLinks != nil {
		if j, ok := e.Options.FeedLinks["json"]; ok {
			feed.FeedURL = &j
		}
	}
	feed.Description = e.Options.Description
	feed.Icon = e.Options.Image
	if e.Options.Author != nil {
		feed.Author = &JSONFeedAuthor{
			Name: e.Options.Author.Name,
			URL:  e.Options.Author.Link,
		}
	}

	if n := len(e.Extensions); n > 0 {
		feed.xattr = make(map[string]interface{}, n)

		for _, extension := range e.Extensions {
			feed.xattr[extension.Name] = extension.Objects
		}
	}

	if n := len(e.Items); n > 0 {
		feed.Items = make([]JSONFeedItem, n)

		for i, item := range e.Items {
			feed.Items[i].ID = item.ID
			feed.Items[i].ContentHTML = item.Content
			feed.Items[i].URL = item.Link
			feed.Items[i].Title = item.Title
			feed.Items[i].Summary = item.Description
			feed.Items[i].Image = item.Image
			if item.Date != nil {
				feed.Items[i].DateModified = item.Date.Format(RFC3339JavaScript)
			}
			if item.Published != nil {
				feed.Items[i].DatePublished = item.Published.Format(RFC3339JavaScript)
			}
			if len(item.Author) > 0 {
				feed.Items[i].Author = &JSONFeedAuthor{
					Name: item.Author[0].Name,
					URL:  item.Author[0].Link,
				}
			}
			// begin category
			var tags []string
			for _, category := range item.Category {
				if category.Name != nil {
					tags = append(tags, *category.Name)
				}
			}
			// this is because category name could be nil
			// we don't have a way to know the amount of categories with non-nil entries before iteration
			// which means we couldn't allocate the array in advance
			if len(tags) > 0 {
				feed.Items[i].Tags = tags
			}
			// end category

			if m := len(item.Extensions); m > 0 {
				feed.Items[i].xattr = make(map[string]interface{}, m)

				for _, extension := range item.Extensions {
					feed.Items[i].xattr[extension.Name] = extension.Objects
				}
			}
		}
	}

	return feed
}
