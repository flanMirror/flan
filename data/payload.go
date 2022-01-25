package data

import (
	"encoding/json"
	"log"
	"random.chars.jp/git/misskey/config"
	"sync"
)

// New returns pointer to a new default Payload.
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

func (p *Payload) JSON() []byte {
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
