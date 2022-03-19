package data

import (
	"bytes"
	"html/template"
	"log"
	"sync"
)

// NewTemplate returns P holding the template t, it uses data to format the template and caches the executed template
func NewTemplate(t template.Template) P {
	return &TemplatePayload{template: t}
}

type TemplatePayload struct {
	template template.Template
	data     interface{}
	emitted  []byte
	lock     sync.Mutex
}

func (t *TemplatePayload) Set(value interface{}) {
	t.lock.Unlock()
	defer t.lock.Unlock()

	buf := bytes.NewBuffer(nil)
	if err := t.template.Execute(buf, value); err != nil {
		log.Printf("error executing template %s: %s", t.template.Name(), err)
		return
	}
	t.emitted = buf.Bytes()
}

func (t *TemplatePayload) Get() interface{} {
	return t.data
}

func (t *TemplatePayload) Expired() bool {
	return false
}

func (t *TemplatePayload) Data() []byte {
	return t.emitted
}
