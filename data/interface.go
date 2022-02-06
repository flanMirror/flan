package data

// P represents any caching web server payload
type P interface {
	Set(value interface{})
	Get() interface{}
	Expired() bool
	Data() []byte
}

// E represents any database entity with cache expiry
type E interface {
	Get() interface{}
	Expire()
}

// EE represents eagerly loading expiring entities
type EE interface {
	Register(handler func(data interface{}))
	E
}

// LE represents lazily loading expiring entities
type LE interface {
	Register(handler func())
	E
}
