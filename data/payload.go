package data

import (
	"encoding/json"
	"log"
	"sync"

	"random.chars.jp/git/misskey/config"
)

// New returns a pointer to Payload with no special features and lazily caches JSON response
func New[T any]() *Payload[T] {
	p := &Payload[T]{immutable: false}
	p.get = p.Get
	return p
}

// Payload represents any caching web server payload
type Payload[T any] struct {
	immutable bool
	get       func() T
	value     T
	json      []byte
	lock      sync.RWMutex
}

// checkImmutable panics if p is immutable
func (p *Payload[T]) checkImmutable() {
	if p.immutable {
		panic("attempting to mutate an immutable payload")
	}
}

// SetImmutable sets p as immutable
func (p *Payload[T]) SetImmutable() {
	p.checkImmutable()
	p.immutable = true
}

// Set sets value held by Payload, cannot be used on am immutable payload
func (p *Payload[T]) Set(v T) {
	p.checkImmutable()

	p.lock.Lock()
	defer p.lock.Unlock()

	p.value = v
	p.json = nil
}

// Get returns the value held by Payload
func (p *Payload[T]) Get() T {
	return p.value
}

// Expired returns whether Payload is between expiry and next Set
func (p *Payload[T]) Expired() bool {
	return p.json == nil
}

// Data returns the (cached) data generated from the value of Payload
func (p *Payload[T]) Data() []byte {
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
