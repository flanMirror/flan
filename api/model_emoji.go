/*
 * Misskey API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Emoji struct {
	Id string `json:"id"`

	Aliases []string `json:"aliases"`

	Name string `json:"name"`

	Category *string `json:"category"`

	Host *string `json:"host"`

	Url string `json:"url"`
}
