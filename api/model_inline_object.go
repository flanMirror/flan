/*
 * Misskey API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject struct {
	Limit float32 `json:"limit,omitempty"`

	SinceId string `json:"sinceId,omitempty"`

	UntilId string `json:"untilId,omitempty"`

	State *string `json:"state,omitempty"`

	ReporterOrigin string `json:"reporterOrigin,omitempty"`

	TargetUserOrigin string `json:"targetUserOrigin,omitempty"`
}
