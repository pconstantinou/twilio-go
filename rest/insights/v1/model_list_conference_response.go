/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListConferenceResponse struct for ListConferenceResponse
type ListConferenceResponse struct {
	Conferences []InsightsV1Conference     `json:"conferences,omitempty"`
	Meta        ListConferenceResponseMeta `json:"meta,omitempty"`
}
