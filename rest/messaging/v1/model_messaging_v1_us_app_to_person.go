/*
 * Twilio - Messaging
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

// MessagingV1UsAppToPerson struct for MessagingV1UsAppToPerson
type MessagingV1UsAppToPerson struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// A2P Brand Registration SID
	BrandRegistrationSid *string `json:"brand_registration_sid,omitempty"`
	// The Campaign Registry (TCR) Campaign ID.
	CampaignId *string `json:"campaign_id,omitempty"`
	// Campaign status
	CampaignStatus *string `json:"campaign_status,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// A short description of what this SMS campaign does
	Description *string `json:"description,omitempty"`
	// Indicate that this SMS campaign will send messages that contain links
	HasEmbeddedLinks *bool `json:"has_embedded_links,omitempty"`
	// Indicates that this SMS campaign will send messages that contain phone numbers
	HasEmbeddedPhone *bool `json:"has_embedded_phone,omitempty"`
	// Indicates whether the campaign was registered externally or not
	IsExternallyRegistered *bool `json:"is_externally_registered,omitempty"`
	// Message samples
	MessageSamples *[]string `json:"message_samples,omitempty"`
	// The SID of the Messaging Service the resource is associated with
	MessagingServiceSid *string `json:"messaging_service_sid,omitempty"`
	// A boolean that specifies whether campaign is a mock or not.
	Mock *bool `json:"mock,omitempty"`
	// Rate limit and/or classification set by each carrier
	RateLimits *map[string]interface{} `json:"rate_limits,omitempty"`
	// The unique string that identifies a US A2P Compliance resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the US App to Person resource
	Url *string `json:"url,omitempty"`
	// A2P Campaign Use Case.
	UsAppToPersonUsecase *string `json:"us_app_to_person_usecase,omitempty"`
}
