package data

// P represents any caching web server payload
type P interface {
	// Set sets value held by P to value
	Set(value interface{})
	// Get returns the value held by E
	Get() interface{}
	// Expired returns whether P is between expiry and next fetch
	Expired() bool
	// Data returns the (cached) payload generated from the value of P
	Data() []byte
}

// E represents any database entity with cache expiry
type E interface {
	// Get returns the value held by E
	Get() interface{}
	// Expire expires the value held by E, requiring fetch on next access for an LE and immediately for an EE
	Expire()
}

// EE represents eagerly loading expiring entities
type EE interface {
	// Register registers a handler function that will be called on the expiry of EE. Must be safe for concurrent use
	Register(handler func(data interface{}))
	E
}

// LE represents lazily loading expiring entities
type LE interface {
	// Register registers a handler function that will be called on the expiry of LE. Must be safe for concurrent use
	Register(handler func())
	E
}

// no, you looked hard enough, there is really no LE implementation yet because it's
// not needed anywhere for the time being, if you need to use it for anything that would
//work better than EE, please email me to make me implement it, or maybe do it yourself
// since it's not that complicated
