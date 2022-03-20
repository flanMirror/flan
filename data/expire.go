package data

import "time"

// NewExpire returns a P that expires every d duration, using the return value of the refresh function
// as the new value it holds
func NewExpire[T any](d time.Duration, refresh func() T) *ExpiringPayload[T] {
	return &ExpiringPayload[T]{
		duration: d,
		refresh:  refresh,
		Payload:  &Payload[T]{get: refresh},
	}
}

// ExpiringPayload implements P in a periodically expiring way
type ExpiringPayload[T any] struct {
	current  time.Time
	duration time.Duration
	refresh  func() T
	*Payload[T]
}

func (e *ExpiringPayload[T]) Get() interface{} {
	e.lock.RLock()
	defer e.lock.RUnlock()

	return e.get()
}

func (e *ExpiringPayload[T]) Expired() bool {
	e.lock.Lock()
	defer e.lock.Unlock()

	n := time.Now()
	if e.current.Add(e.duration).Before(n) {
		e.current = n
		return true
	}
	return false
}
