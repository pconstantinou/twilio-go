/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListUserResponse struct for ListUserResponse
type ListUserResponse struct {
	Meta  ListCredentialResponseMeta `json:"meta,omitempty"`
	Users []ChatV1User               `json:"users,omitempty"`
}
