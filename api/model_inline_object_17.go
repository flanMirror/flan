/*
 * Misskey API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject17 struct {
	Query *string `json:"query,omitempty"`

	Limit float32 `json:"limit,omitempty"`

	SinceId string `json:"sinceId,omitempty"`

	UntilId string `json:"untilId,omitempty"`
}
