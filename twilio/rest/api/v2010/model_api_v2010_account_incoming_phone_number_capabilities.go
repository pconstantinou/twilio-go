/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountIncomingPhoneNumberCapabilities Indicate if a phone can receive calls or messages
type ApiV2010AccountIncomingPhoneNumberCapabilities struct {
	Fax   bool `json:"Fax,omitempty"`
	Mms   bool `json:"Mms,omitempty"`
	Sms   bool `json:"Sms,omitempty"`
	Voice bool `json:"Voice,omitempty"`
}