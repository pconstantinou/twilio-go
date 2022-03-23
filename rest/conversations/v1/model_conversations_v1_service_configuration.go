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

// ConversationsV1ServiceConfiguration struct for ConversationsV1ServiceConfiguration
type ConversationsV1ServiceConfiguration struct {
	// The unique string that identifies the resource
	ChatServiceSid *string `json:"chat_service_sid,omitempty"`
	// The service role assigned to users when they are added to the service
	DefaultChatServiceRoleSid *string `json:"default_chat_service_role_sid,omitempty"`
	// The role assigned to a conversation creator user when they join a new conversation
	DefaultConversationCreatorRoleSid *string `json:"default_conversation_creator_role_sid,omitempty"`
	// The role assigned to users when they are added to a conversation
	DefaultConversationRoleSid *string `json:"default_conversation_role_sid,omitempty"`
	// Absolute URL to access the push notifications configuration of this service.
	Links *map[string]interface{} `json:"links,omitempty"`
	// Whether the Reachability Indicator feature is enabled for this Conversations Service
	ReachabilityEnabled *bool `json:"reachability_enabled,omitempty"`
	// An absolute URL for this service configuration.
	Url *string `json:"url,omitempty"`
}
