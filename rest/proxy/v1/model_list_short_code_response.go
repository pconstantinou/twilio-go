/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListShortCodeResponse struct for ListShortCodeResponse
type ListShortCodeResponse struct {
	Meta       ListServiceResponseMeta `json:"meta,omitempty"`
	ShortCodes []ProxyV1ShortCode      `json:"short_codes,omitempty"`
}
