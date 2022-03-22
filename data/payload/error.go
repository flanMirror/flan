package payload

import (
	"random.chars.jp/git/misskey/api/response"
	"random.chars.jp/git/misskey/data"
)

const kindZero = "client"

// see api/NOTES for detailed usage of these error payloads

var (
	// InternalServerError should be sent with http.StatusInternalServerError
	InternalServerError = data.New[response.APIError]()
	// FileRequired should be sent with http.StatusBadRequest
	FileRequired = data.New[response.APIError]()
	// AccessDenied should be sent with http.StatusForbidden
	AccessDenied = data.New[response.APIError]()
	// AccessDeniedNotAdmin should be sent with http.StatusForbidden
	AccessDeniedNotAdmin = data.New[response.APIError]()
	// AccessDeniedNotModerator should be sent with http.StatusForbidden
	AccessDeniedNotModerator = data.New[response.APIError]()
	// NoSuchEndpoint should be sent with http.StatusNotFound
	NoSuchEndpoint = data.New[response.APIError]()
	// CredentialRequired should be sent with http.StatusUnauthorized
	CredentialRequired = data.New[response.APIError]()
	// AccountSuspended should be sent with http.StatusForbidden
	AccountSuspended = data.New[response.APIError]()
	// AppLackingPermission should be sent with http.StatusForbidden
	AppLackingPermission = data.New[response.APIError]()
	// RateLimitExceeded should be sent with http.StatusTooManyRequests
	RateLimitExceeded = data.New[response.APIError]()
)

func init() {
	InternalServerError.Set(response.APIError{
		Message: "Internal error occurred. Please contact us if the error persists.",
		Code:    "INTERNAL_ERROR",
		ID:      "5d37dbcb-891e-41ca-a3d6-e690c97775ac",
		Kind:    "server",
	})
	InternalServerError.SetImmutable()

	FileRequired.Set(response.APIError{
		Message: "File required.",
		Code:    "FILE_REQUIRED",
		ID:      "4267801e-70d1-416a-b011-4ee502885d8b",
		Kind:    kindZero,
	})
	FileRequired.SetImmutable()

	AccessDenied.Set(newAccessDenied(nil))
	AccessDenied.SetImmutable()

	AccessDeniedNotAdmin.Set(NewAccessDenied("You are not the admin."))
	AccessDeniedNotAdmin.SetImmutable()

	AccessDeniedNotModerator.Set(NewAccessDenied("You are not a moderator."))
	AccessDeniedNotModerator.SetImmutable()

	// I don't know what this does, but it appears in code
	NoSuchEndpoint.Set(response.APIError{
		Message: "No such endpoint.",
		Code:    "NO_SUCH_ENDPOINT",
		ID:      "f8080b67-5f9c-4eb7-8c18-7f1eeae8f709",
		Kind:    kindZero,
	})
	NoSuchEndpoint.SetImmutable()

	CredentialRequired.Set(response.APIError{
		Message: "Credential required.",
		Code:    "CREDENTIAL_REQUIRED",
		ID:      "1384574d-a912-4b81-8601-c7b1c4085df1",
		Kind:    kindZero,
	})
	CredentialRequired.SetImmutable()

	AccountSuspended.Set(response.APIError{
		Message: "Your account has been suspended.",
		Code:    "YOUR_ACCOUNT_SUSPENDED",
		ID:      "a8c724b3-6e9c-4b46-b1a8-bc3ed6258370",
		Kind:    kindZero,
	})
	AccountSuspended.SetImmutable()

	AppLackingPermission.Set(response.APIError{
		Message: "Your app does not have the necessary permissions to use this endpoint.",
		Code:    "PERMISSION_DENIED",
		ID:      "1370e5b7-d4eb-4566-bb1d-7748ee6a1838",
		Kind:    kindZero,
	})
	AppLackingPermission.SetImmutable()

	RateLimitExceeded.Set(response.APIError{
		Message: "Rate limit exceeded. Please try again later.",
		Code:    "RATE_LIMIT_EXCEEDED",
		ID:      "d5826d14-3982-4d2e-8011-b9e9f02499ef",
		Kind:    kindZero,
	})
	RateLimitExceeded.SetImmutable()
}

func NewInvalidParam(key, message string) response.APIError {
	return response.APIError{
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

func NewAccessDenied(reason string) response.APIError {
	return newAccessDenied(struct {
		Reason string `json:"reason"`
	}{reason})
}

func newAccessDenied(info interface{}) response.APIError {
	return response.APIError{
		Message: "Access denied.",
		Code:    "ACCESS_DENIED",
		ID:      "56f35758-7dd5-468b-8439-5d6fb8ec9b8e",
		Kind:    kindZero,
		Info:    info,
	}
}
