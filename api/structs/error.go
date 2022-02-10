package structs

// APIError is the struct returned on API error,
// zero value should be payload.InternalServerError,
// Kind should default to "client"
type APIError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	ID      string `json:"id"`
	Kind    string `json:"kind"`
	//HttpStatusCode int `json:"-"`
	Info interface{} `json:"info,omitempty"`
}
