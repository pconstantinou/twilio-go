/*
 * Twilio - Conversations
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

// ConversationsV1Conversation struct for ConversationsV1Conversation
type ConversationsV1Conversation struct {
	// The unique ID of the Account responsible for this conversation.
	AccountSid *string `json:"account_sid,omitempty"`
	// An optional string metadata field you can use to store any data you wish.
	Attributes *string                 `json:"attributes,omitempty"`
	Bindings   *map[string]interface{} `json:"bindings,omitempty"`
	// The unique ID of the Conversation Service this conversation belongs to.
	ChatServiceSid *string `json:"chat_service_sid,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The human-readable name of this conversation.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// Absolute URLs to access the participants, messages and webhooks of this conversation.
	Links *map[string]interface{} `json:"links,omitempty"`
	// The unique ID of the Messaging Service this conversation belongs to.
	MessagingServiceSid *string `json:"messaging_service_sid,omitempty"`
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"sid,omitempty"`
	// Current state of this conversation.
	State *string `json:"state,omitempty"`
	// Timer date values for this conversation.
	Timers *map[string]interface{} `json:"timers,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"unique_name,omitempty"`
	// An absolute URL for this conversation.
	Url *string `json:"url,omitempty"`
}
