/*
 * Twilio - Accounts
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

// AccountsV1CredentialPublicKey struct for AccountsV1CredentialPublicKey
type AccountsV1CredentialPublicKey struct {
	// The SID of the Account that created the Credential that the PublicKey resource belongs to
	AccountSid *string `json:"account_sid,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The URI for this resource, relative to `https://accounts.twilio.com`
	Url *string `json:"url,omitempty"`
}
