package data

import (
	"encoding/json"
	"log"
	"sync"

	"random.chars.jp/git/misskey/config"
)

// New returns a P implementation with no special features and lazily caches JSON response
func New() P {
	p := &Payload{}
	p.get = p.Get
	return p
}

type Payload struct {
	get   func() interface{}
	value interface{}
	json  []byte
	lock  sync.RWMutex
}

func (p *Payload) Set(v interface{}) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.value = v
	p.json = nil
}

func (p *Payload) Get() interface{} {
	return p.value
}

func (p *Payload) Expired() bool {
	return p.json == nil
}

func (p *Payload) Data() []byte {
	if p.Expired() {
		p.lock.Lock()
		defer p.lock.Unlock()
		if data, err := json.Marshal(p.get()); err != nil {
			if config.Log.Verbose {
				log.Printf("error marshalling payload: %s", err)
			}
			return nil
		} else {
			p.json = data
		}
	}

	return p.json
}
