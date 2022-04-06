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

// FlexV1InteractionChannelInvite struct for FlexV1InteractionChannelInvite
type FlexV1InteractionChannelInvite struct {
	ChannelSid     *string                 `json:"channel_sid,omitempty"`
	InteractionSid *string                 `json:"interaction_sid,omitempty"`
	Routing        *map[string]interface{} `json:"routing,omitempty"`
	Sid            *string                 `json:"sid,omitempty"`
	Url            *string                 `json:"url,omitempty"`
}
