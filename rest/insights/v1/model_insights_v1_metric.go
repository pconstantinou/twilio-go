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

// InsightsV1Metric struct for InsightsV1Metric
type InsightsV1Metric struct {
	AccountSid  *string                 `json:"account_sid,omitempty"`
	CallSid     *string                 `json:"call_sid,omitempty"`
	CarrierEdge *map[string]interface{} `json:"carrier_edge,omitempty"`
	ClientEdge  *map[string]interface{} `json:"client_edge,omitempty"`
	Direction   *string                 `json:"direction,omitempty"`
	Edge        *string                 `json:"edge,omitempty"`
	SdkEdge     *map[string]interface{} `json:"sdk_edge,omitempty"`
	SipEdge     *map[string]interface{} `json:"sip_edge,omitempty"`
	Timestamp   *string                 `json:"timestamp,omitempty"`
}
