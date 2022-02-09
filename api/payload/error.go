package payload

import (
	"random.chars.jp/git/misskey/api/structs"
	"random.chars.jp/git/misskey/data"
)

var InternalServerError = data.New()

func init() {
	InternalServerError.Set(structs.APIError{
		Message:        "Internal error occurred. Please contact us if the error persists.",
		Code:           "INTERNAL_ERROR",
		ID:             "5d37dbcb-891e-41ca-a3d6-e690c97775ac",
		Kind:           "server",
		HttpStatusCode: 500,
	})
}
