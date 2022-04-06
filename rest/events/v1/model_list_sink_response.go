/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSinkResponse struct for ListSinkResponse
type ListSinkResponse struct {
	Meta  ListSchemaVersionResponseMeta `json:"meta,omitempty"`
	Sinks []EventsV1Sink                `json:"sinks,omitempty"`
}
