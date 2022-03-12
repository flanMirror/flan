package feed

func New(opt Options) *Emitter {
	return &Emitter{Options: opt}
}

type Emitter struct {
	Options      Options     `json:"options"`
	Items        []Item      `json:"items"`
	Categories   []Category  `json:"categories"`
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
func (e *Emitter) AddCategories(categories []Category) {
	offset := len(e.Categories)
	newCategories := make([]Category, offset+len(categories))
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
