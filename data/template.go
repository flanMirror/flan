package data

import (
	"bytes"
	"html/template"
	"log"
	"sync"
)

// NewTemplate returns P holding the template t, it uses value to format the template and caches the executed template
func NewTemplate[T any](t *template.Template) *TemplatePayload[T] {
	return &TemplatePayload[T]{template: t}
}

type TemplatePayload[T any] struct {
	template *template.Template
	value    T
	data     []byte
	lock     sync.Mutex
}

func (t *TemplatePayload[T]) Set(value T) {
	t.lock.Lock()
	defer t.lock.Unlock()

	buf := bytes.NewBuffer(nil)
	if err := t.template.Execute(buf, value); err != nil {
		log.Printf("error executing template %s: %s", t.template.Name(), err)
		return
	}
	t.data = buf.Bytes()
}

func (t *TemplatePayload[T]) Get() T {
	return t.value
}

func (t *TemplatePayload[T]) Expired() bool {
	return false
}

func (t *TemplatePayload[T]) Data() []byte {
	return t.data
}
