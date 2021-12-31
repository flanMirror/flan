/*
 * Misskey API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

type InlineResponse20032 struct {
	Id string `json:"id"`

	CreatedAt time.Time `json:"createdAt"`

	StartedAt time.Time `json:"startedAt"`

	IsStarted bool `json:"isStarted"`

	IsEnded bool `json:"isEnded"`

	Form1 *Any `json:"form1"`

	Form2 *Any `json:"form2"`

	User1Accepted bool `json:"user1Accepted"`

	User2Accepted bool `json:"user2Accepted"`

	User1Id string `json:"user1Id"`

	User2Id string `json:"user2Id"`

	User1 User `json:"user1"`

	User2 User `json:"user2"`

	WinnerId *string `json:"winnerId"`

	Winner User `json:"winner"`

	Surrendered *string `json:"surrendered"`

	Black *float32 `json:"black"`

	Bw string `json:"bw"`

	IsLlotheo bool `json:"isLlotheo"`

	CanPutEverywhere bool `json:"canPutEverywhere"`

	LoopedBoard bool `json:"loopedBoard"`

	Board []Any `json:"board"`

	Turn Any `json:"turn"`
}
