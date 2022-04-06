/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListBucketResponse struct for ListBucketResponse
type ListBucketResponse struct {
	Buckets []VerifyV2Bucket                    `json:"buckets,omitempty"`
	Meta    ListVerificationAttemptResponseMeta `json:"meta,omitempty"`
}
