/*
 * Misskey API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject152 struct {
	Local bool `json:"local,omitempty"`

	Reply bool `json:"reply,omitempty"`

	Renote bool `json:"renote,omitempty"`

	WithFiles bool `json:"withFiles,omitempty"`

	Poll bool `json:"poll,omitempty"`

	Limit float32 `json:"limit,omitempty"`

	SinceId string `json:"sinceId,omitempty"`

	UntilId string `json:"untilId,omitempty"`
}
