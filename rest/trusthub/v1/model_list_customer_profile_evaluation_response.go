/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListCustomerProfileEvaluationResponse struct for ListCustomerProfileEvaluationResponse
type ListCustomerProfileEvaluationResponse struct {
	Meta    ListCustomerProfileResponseMeta       `json:"meta,omitempty"`
	Results []TrusthubV1CustomerProfileEvaluation `json:"results,omitempty"`
}
