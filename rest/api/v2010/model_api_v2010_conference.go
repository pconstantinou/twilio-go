/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010Conference struct for ApiV2010Conference
type ApiV2010Conference struct {
	// The SID of the Account that created this resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The API version used to create this conference
	ApiVersion *string `json:"api_version,omitempty"`
	// The call SID that caused the conference to end
	CallSidEndingConference *string `json:"call_sid_ending_conference,omitempty"`
	// The RFC 2822 date and time in GMT that this resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that this resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// A string that you assigned to describe this conference room
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The reason why a conference ended.
	ReasonConferenceEnded *string `json:"reason_conference_ended,omitempty"`
	// A string that represents the Twilio Region where the conference was mixed
	Region *string `json:"region,omitempty"`
	// The unique string that identifies this resource
	Sid *string `json:"sid,omitempty"`
	// The status of this conference
	Status *string `json:"status,omitempty"`
	// A list of related resources identified by their relative URIs
	SubresourceUris *map[string]interface{} `json:"subresource_uris,omitempty"`
	// The URI of this resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}
