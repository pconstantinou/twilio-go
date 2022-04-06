/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// FlexV1InteractionChannel struct for FlexV1InteractionChannel
type FlexV1InteractionChannel struct {
	// The Interaction Sid for this channel.
	InteractionSid *string                 `json:"interaction_sid,omitempty"`
	Links          *map[string]interface{} `json:"links,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The Interaction Channel's type.
	Type *string `json:"type,omitempty"`
	Url  *string `json:"url,omitempty"`
}
