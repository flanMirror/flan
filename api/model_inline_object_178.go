/*
 * Misskey API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject178 struct {
	Limit float32 `json:"limit,omitempty"`

	SinceId string `json:"sinceId,omitempty"`

	UntilId string `json:"untilId,omitempty"`

	SinceDate float32 `json:"sinceDate,omitempty"`

	UntilDate float32 `json:"untilDate,omitempty"`

	IncludeMyRenotes bool `json:"includeMyRenotes,omitempty"`

	IncludeRenotedMyNotes bool `json:"includeRenotedMyNotes,omitempty"`

	IncludeLocalRenotes bool `json:"includeLocalRenotes,omitempty"`

	WithFiles bool `json:"withFiles,omitempty"`
}
