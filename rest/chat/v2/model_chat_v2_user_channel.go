/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ChatV2UserChannel struct for ChatV2UserChannel
type ChatV2UserChannel struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Channel the resource belongs to
	ChannelSid *string `json:"channel_sid,omitempty"`
	// The index of the last Message in the Channel the Member has read
	LastConsumedMessageIndex *int `json:"last_consumed_message_index,omitempty"`
	// Absolute URLs to access the Members, Messages , Invites and, if it exists, the last Message for the Channel
	Links *map[string]interface{} `json:"links,omitempty"`
	// The SID of the User as a Member in the Channel
	MemberSid *string `json:"member_sid,omitempty"`
	// The push notification level of the User for the Channel
	NotificationLevel *string `json:"notification_level,omitempty"`
	// The SID of the Service that the resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The status of the User on the Channel
	Status *string `json:"status,omitempty"`
	// The number of unread Messages in the Channel for the User
	UnreadMessagesCount *int `json:"unread_messages_count,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
	// The SID of the User the User Channel belongs to
	UserSid *string `json:"user_sid,omitempty"`
}
