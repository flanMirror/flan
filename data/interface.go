package data

type P interface {
	Set(value interface{})
	Get() interface{}
	Expired() bool
	Data() []byte
}
