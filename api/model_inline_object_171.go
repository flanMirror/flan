/*
 * Misskey API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject171 struct {
	NoteId string `json:"noteId"`

	SinceId string `json:"sinceId,omitempty"`

	UntilId string `json:"untilId,omitempty"`

	Limit float32 `json:"limit,omitempty"`
}
