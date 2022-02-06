package data

import (
	"log"
	"random.chars.jp/git/misskey/config"
	"runtime"
	"sync"
)

type Eager struct {
	fetch    func() interface{}
	handlers []func(data interface{})
	data     interface{}
	lock     sync.Mutex
}

func (e *Eager) Expire() {
	e.lock.Lock()
	defer e.lock.Unlock()

	if data := e.fetch(); data == nil {
		if config.Log.Verbose {
			pc := make([]uintptr, 15)
			n := runtime.Callers(2, pc)
			frames := runtime.CallersFrames(pc[:n])
			frame, _ := frames.Next()

			log.Printf("got nil in expiry eager load at %s:%d %s",
				frame.File, frame.Line, frame.Function)
		}
		return
	} else {
		e.data = e.fetch()
	}

	w := sync.WaitGroup{}
	w.Add(len(e.handlers))
	for _, handler := range e.handlers {
		go func(handler func(interface{})) {
			handler(e.data)
			w.Done()
		}(handler)
	}
	w.Wait()
}

func (e *Eager) Get() interface{} {
	if e.data != nil {
		e.lock.Lock()
		defer e.lock.Unlock()
		e.data = e.fetch()
	}
	return e.data
}

func (e *Eager) Register(handler func(data interface{})) {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.handlers = append(e.handlers, handler)
}

func NewEager(fetch func() interface{}) EE {
	return &Eager{
		fetch: fetch,
		lock:  sync.Mutex{},
	}
}
