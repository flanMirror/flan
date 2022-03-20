package data

import (
	"log"
	"runtime"
	"sync"

	"random.chars.jp/git/misskey/config"
)

// Eager represents eagerly loading expiring entities
type Eager[T any] struct {
	fetch    func() (T, bool)
	handlers []func(value T)
	value    T
	lock     sync.Mutex
}

// Expire expires the value held by Eager, requiring fetch immediately and invocation of all handlers
func (e *Eager[T]) Expire() {
	e.lock.Lock()
	defer e.lock.Unlock()

	if value, ok := e.fetch(); !ok {
		if config.Log.Verbose {
			pc := make([]uintptr, 15)
			n := runtime.Callers(2, pc)
			frames := runtime.CallersFrames(pc[:n])
			frame, _ := frames.Next()

			log.Printf("got zero in expiry eager load at %s:%d %s",
				frame.File, frame.Line, frame.Function)
		}
		return
	} else {
		e.value = value
	}

	w := sync.WaitGroup{}
	w.Add(len(e.handlers))
	for _, handler := range e.handlers {
		go func(handler func(value T)) {
			handler(e.value)
			w.Done()
		}(handler)
	}
	w.Wait()
}

// Get returns the value held by Eager
func (e *Eager[T]) Get() T {
	return e.value
}

// Register registers a handler function that will be called on the expiry of Eager
func (e *Eager[T]) Register(handler func(value T)) {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.handlers = append(e.handlers, handler)
}

// NewEager returns a new eagerly loading EE with fetch as the function called to eagerly load.
// The return value of fetch would be set as the value of the EE whenever a load happens.
func NewEager[T any](fetch func() (T, bool)) *Eager[T] {
	return &Eager[T]{
		fetch: fetch,
		lock:  sync.Mutex{},
	}
}
