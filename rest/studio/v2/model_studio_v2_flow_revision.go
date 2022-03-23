/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// StudioV2FlowRevision struct for StudioV2FlowRevision
type StudioV2FlowRevision struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// Description of change made in the revision
	CommitMessage *string `json:"commit_message,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// JSON representation of flow definition
	Definition *map[string]interface{} `json:"definition,omitempty"`
	// List of error in the flow definition
	Errors *[]map[string]interface{} `json:"errors,omitempty"`
	// The string that you assigned to describe the Flow
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The latest revision number of the Flow's definition
	Revision *int `json:"revision,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The status of the Flow
	Status *string `json:"status,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
	// Boolean if the flow definition is valid
	Valid *bool `json:"valid,omitempty"`
}
