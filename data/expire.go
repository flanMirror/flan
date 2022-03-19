package data

import "time"

// NewExpire returns a P that expires every d duration, using the return value of the refresh function
// as the new value it holds
func NewExpire(d time.Duration, refresh func() interface{}) P {
	return &ExpiringPayload{
		duration: d,
		refresh:  refresh,
		Payload:  &Payload{get: refresh},
	}
}

// ExpiringPayload implements P in a periodically expiring way
type ExpiringPayload struct {
	current  time.Time
	duration time.Duration
	refresh  func() interface{}
	*Payload
}

func (e *ExpiringPayload) Get() interface{} {
	e.lock.RLock()
	defer e.lock.RUnlock()

	return e.get()
}

func (e *ExpiringPayload) Expired() bool {
	e.lock.Lock()
	defer e.lock.Unlock()

	n := time.Now()
	if e.current.Add(e.duration).Before(n) {
		e.current = n
		return true
	}
	return false
}
