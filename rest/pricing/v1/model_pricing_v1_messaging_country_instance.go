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

// PricingV1MessagingCountryInstance struct for PricingV1MessagingCountryInstance
type PricingV1MessagingCountryInstance struct {
	// The name of the country
	Country *string `json:"country,omitempty"`
	// The list of InboundPrice records
	InboundSmsPrices *[]PricingV1MessagingMessagingCountryInstanceInboundSmsPrices `json:"inbound_sms_prices,omitempty"`
	// The ISO country code
	IsoCountry *string `json:"iso_country,omitempty"`
	// The list of OutboundSMSPrice records
	OutboundSmsPrices *[]PricingV1MessagingMessagingCountryInstanceOutboundSmsPrices `json:"outbound_sms_prices,omitempty"`
	// The currency in which prices are measured, in ISO 4127 format (e.g. usd, eur, jpy)
	PriceUnit *string `json:"price_unit,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
