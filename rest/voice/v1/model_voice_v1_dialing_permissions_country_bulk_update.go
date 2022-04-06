/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// VoiceV1DialingPermissionsCountryBulkUpdate struct for VoiceV1DialingPermissionsCountryBulkUpdate
type VoiceV1DialingPermissionsCountryBulkUpdate struct {
	// The number of countries updated
	UpdateCount *int `json:"update_count,omitempty"`
	// A URL encoded JSON array of update objects
	UpdateRequest *string `json:"update_request,omitempty"`
}
