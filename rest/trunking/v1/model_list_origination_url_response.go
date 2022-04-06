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

// ListOriginationUrlResponse struct for ListOriginationUrlResponse
type ListOriginationUrlResponse struct {
	Meta            ListTrunkResponseMeta      `json:"meta,omitempty"`
	OriginationUrls []TrunkingV1OriginationUrl `json:"origination_urls,omitempty"`
}
