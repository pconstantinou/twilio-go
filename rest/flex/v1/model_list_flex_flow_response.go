/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListFlexFlowResponse struct for ListFlexFlowResponse
type ListFlexFlowResponse struct {
	FlexFlows []FlexV1FlexFlow        `json:"flex_flows,omitempty"`
	Meta      ListChannelResponseMeta `json:"meta,omitempty"`
}
