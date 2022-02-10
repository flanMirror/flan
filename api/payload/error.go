package payload

import (
	"random.chars.jp/git/misskey/api/structs"
	"random.chars.jp/git/misskey/data"
)

const kindZero = "client"

// see api/NOTES for detailed usage of these error payloads

var (
	// InternalServerError should be sent with http.StatusInternalServerError
	InternalServerError = data.New()
	// FileRequired should be sent with http.StatusBadRequest
	FileRequired = data.New()
	// AccessDenied should be sent with http.StatusForbidden
	AccessDenied = data.New()
	// AccessDeniedNotAdmin should be sent with http.StatusForbidden
	AccessDeniedNotAdmin = data.New()
	// AccessDeniedNotModerator should be sent with http.StatusForbidden
	AccessDeniedNotModerator = data.New()
	// NoSuchEndpoint should be sent with http.StatusNotFound
	NoSuchEndpoint = data.New()
	// CredentialRequired should be sent with http.StatusUnauthorized
	CredentialRequired = data.New()
	// AccountSuspended should be sent with http.StatusForbidden
	AccountSuspended = data.New()
	// AppLackingPermission should be sent with http.StatusForbidden
	AppLackingPermission = data.New()
	// RateLimitExceeded should be sent with http.StatusTooManyRequests
	RateLimitExceeded = data.New()
)

func init() {
	InternalServerError.Set(structs.APIError{
		Message: "Internal error occurred. Please contact us if the error persists.",
		Code:    "INTERNAL_ERROR",
		ID:      "5d37dbcb-891e-41ca-a3d6-e690c97775ac",
		Kind:    "server",
	})
	FileRequired.Set(structs.APIError{
		Message: "File required.",
		Code:    "FILE_REQUIRED",
		ID:      "4267801e-70d1-416a-b011-4ee502885d8b",
		Kind:    kindZero,
	})
	AccessDenied.Set(newAccessDenied(nil))
	AccessDeniedNotAdmin.Set(NewAccessDenied("You are not the admin."))
	AccessDeniedNotModerator.Set(NewAccessDenied("You are not a moderator."))
	// I don't know what this does
	NoSuchEndpoint.Set(structs.APIError{
		Message: "No such endpoint.",
		Code:    "NO_SUCH_ENDPOINT",
		ID:      "f8080b67-5f9c-4eb7-8c18-7f1eeae8f709",
		Kind:    kindZero,
	})
	CredentialRequired.Set(structs.APIError{
		Message: "Credential required.",
		Code:    "CREDENTIAL_REQUIRED",
		ID:      "1384574d-a912-4b81-8601-c7b1c4085df1",
		Kind:    kindZero,
	})
	AccountSuspended.Set(structs.APIError{
		Message: "Your account has been suspended.",
		Code:    "YOUR_ACCOUNT_SUSPENDED",
		ID:      "a8c724b3-6e9c-4b46-b1a8-bc3ed6258370",
		Kind:    kindZero,
	})
	AppLackingPermission.Set(structs.APIError{
		Message: "Your app does not have the necessary permissions to use this endpoint.",
		Code:    "PERMISSION_DENIED",
		ID:      "1370e5b7-d4eb-4566-bb1d-7748ee6a1838",
		Kind:    kindZero,
	})
	RateLimitExceeded.Set(structs.APIError{
		Message: "Rate limit exceeded. Please try again later.",
		Code:    "RATE_LIMIT_EXCEEDED",
		ID:      "d5826d14-3982-4d2e-8011-b9e9f02499ef",
		Kind:    kindZero,
	})
}

func NewInvalidParam(key, message string) structs.APIError {
	return structs.APIError{
		Message: "Invalid param.",
		Code:    "INVALID_PARAM",
		ID:      "3d81ceae-475f-4600-b2a8-2bc116157532",
		Kind:    kindZero,
		Info: struct {
			Param  string `json:"param"`
			Reason string `json:"reason"`
		}{key, message},
	}
}

func NewAccessDenied(reason string) structs.APIError {
	return newAccessDenied(struct {
		Reason string `json:"reason"`
	}{reason})
}

func newAccessDenied(info interface{}) structs.APIError {
	return structs.APIError{
		Message: "Access denied.",
		Code:    "ACCESS_DENIED",
		ID:      "56f35758-7dd5-468b-8439-5d6fb8ec9b8e",
		Kind:    kindZero,
		Info:    info,
	}
}
