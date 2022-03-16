package feed

import (
	"strings"
)

/*
I am not going to deal with the UTC/GMT inconsistency, I don't think it matters, and I have sunk way too much time
porting this library to Go. I hope some future programmer that has the misfortune of having to simulate the behaviour
of the npm "feed" package find this useful, although I don't think anyone would think of looking in the reimplementation
 of a random social media software... maybe I'll move this package to its own repository, or I move all the ported npm
stuff to their own repository, I really don't know. writing this is driving me insane, please make it end soon
*/

const (
	RFC3339JavaScript = "2006-01-02T15:04:05.000Z07:00"
	DefaultGenerator  = "misskey-fast"
)

func New(opt Options) *Emitter {
	return &Emitter{Options: opt}
}

type Emitter struct {
	Options      Options     `json:"options"`
	Items        []Item      `json:"items"`
	Categories   []string    `json:"categories"`
	Contributors []Author    `json:"contributors"`
	Extensions   []Extension `json:"extensions"`
}

// AddItems adds an array of Item to the emitter
func (e *Emitter) AddItems(items []Item) {
	offset := len(e.Items)
	newItems := make([]Item, offset+len(items))
	for i := 0; i < len(e.Items); i++ {
		newItems[i] = e.Items[i]
	}
	for i := 0; i < len(items); i++ {
		newItems[i+offset] = items[i]
	}
	e.Items = newItems
}

// AddCategories adds an array of Category to the emitter
func (e *Emitter) AddCategories(categories []string) {
	offset := len(e.Categories)
	newCategories := make([]string, offset+len(categories))
	for i := 0; i < len(e.Categories); i++ {
		newCategories[i] = e.Categories[i]
	}
	for i := 0; i < len(categories); i++ {
		newCategories[i+offset] = categories[i]
	}
	e.Categories = newCategories
}

// AddContributors adds an array of Author to the emitter
func (e *Emitter) AddContributors(contributors []Author) {
	offset := len(e.Contributors)
	newContributors := make([]Author, offset+len(contributors))
	for i := 0; i < len(e.Contributors); i++ {
		newContributors[i] = e.Contributors[i]
	}
	for i := 0; i < len(contributors); i++ {
		newContributors[i+offset] = contributors[i]
	}
	e.Contributors = newContributors
}

// AddExtensions adds an array of Extension to the emitter
func (e *Emitter) AddExtensions(extensions []Extension) {
	offset := len(e.Extensions)
	newExtensions := make([]Extension, offset+len(extensions))
	for i := 0; i < len(e.Extensions); i++ {
		newExtensions[i] = e.Extensions[i]
	}
	for i := 0; i < len(extensions); i++ {
		newExtensions[i+offset] = extensions[i]
	}
	e.Extensions = newExtensions
}

func sanitize(str *string) string {
	if str == nil {
		return ""
	}
	return strings.ReplaceAll(*str, "&", "&amp;")
}
