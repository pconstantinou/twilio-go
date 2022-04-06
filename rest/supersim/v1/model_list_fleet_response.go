/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListFleetResponse struct for ListFleetResponse
type ListFleetResponse struct {
	Fleets []SupersimV1Fleet           `json:"fleets,omitempty"`
	Meta   ListEsimProfileResponseMeta `json:"meta,omitempty"`
}
