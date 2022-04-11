package pack

import (
	json "github.com/json-iterator/go"
)

// valueHolder holds a value as an unexported field
type valueHolder[T any] struct {
	v T
}

// newValueHolder shortcuts the creation of a value holder
func newValueHolder[T any](value T) valueHolder[T] {
	return valueHolder[T]{v: value}
}

// finalPackable is an internal value containing all fields, and can be packed
type finalPackable interface {
	// packAll ensures finalPackable is fully packed
	packAll() error
}

// Packed is a lazily packed database entity
type Packed[T any] struct {
	final finalPackable
	T
}

func (p *Packed) MarshalJSON() ([]byte, error) {
	if err := p.final.packAll(); err != nil {
		return nil, err
	}
	return json.Marshal(p.final)
}
