/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// TrusthubV1TrustProductChannelEndpointAssignment struct for TrusthubV1TrustProductChannelEndpointAssignment
type TrusthubV1TrustProductChannelEndpointAssignment struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The sid of an channel endpoint
	ChannelEndpointSid *string `json:"channel_endpoint_sid,omitempty"`
	// The type of channel endpoint
	ChannelEndpointType *string `json:"channel_endpoint_type,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The unique string that identifies the CustomerProfile resource.
	TrustProductSid *string `json:"trust_product_sid,omitempty"`
	// The absolute URL of the Identity resource
	Url *string `json:"url,omitempty"`
}
