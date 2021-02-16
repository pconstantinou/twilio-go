/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ConversationsV1Configuration struct for ConversationsV1Configuration
type ConversationsV1Configuration struct {
	AccountSid string `json:"AccountSid,omitempty"`
	DefaultChatServiceSid string `json:"DefaultChatServiceSid,omitempty"`
	DefaultClosedTimer string `json:"DefaultClosedTimer,omitempty"`
	DefaultInactiveTimer string `json:"DefaultInactiveTimer,omitempty"`
	DefaultMessagingServiceSid string `json:"DefaultMessagingServiceSid,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	Url string `json:"Url,omitempty"`
}