/*
 * Twilio - Trunking
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListIpAccessControlListResponse struct for ListIpAccessControlListResponse
type ListIpAccessControlListResponse struct {
	IpAccessControlLists []TrunkingV1IpAccessControlList `json:"ip_access_control_lists,omitempty"`
	Meta                 ListTrunkResponseMeta           `json:"meta,omitempty"`
}
