/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// PricingV2TrunkingNumberOriginatingCallPrice The OriginatingCallPrice record
type PricingV2TrunkingNumberOriginatingCallPrice struct {
	BasePrice    float32 `json:"base_price,omitempty"`
	CurrentPrice float32 `json:"current_price,omitempty"`
	NumberType   string  `json:"number_type,omitempty"`
}
