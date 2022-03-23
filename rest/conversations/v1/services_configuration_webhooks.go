/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"net/url"
	"strings"
)

// Fetch a specific service webhook configuration.
func (c *ApiService) FetchServiceWebhookConfiguration(ChatServiceSid string) (*ConversationsV1ServiceWebhookConfiguration, error) {
	path := "/v1/Services/{ChatServiceSid}/Configuration/Webhooks"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceWebhookConfiguration{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateServiceWebhookConfiguration'
type UpdateServiceWebhookConfigurationParams struct {
	// The list of events that your configured webhook targets will receive. Events not configured here will not fire. Possible values are `onParticipantAdd`, `onParticipantAdded`, `onDeliveryUpdated`, `onConversationUpdated`, `onConversationRemove`, `onParticipantRemove`, `onConversationUpdate`, `onMessageAdd`, `onMessageRemoved`, `onParticipantUpdated`, `onConversationAdded`, `onMessageAdded`, `onConversationAdd`, `onConversationRemoved`, `onParticipantUpdate`, `onMessageRemove`, `onMessageUpdated`, `onParticipantRemoved`, `onMessageUpdate` or `onConversationStateUpdated`.
	Filters *[]string `json:"Filters,omitempty"`
	// The HTTP method to be used when sending a webhook request. One of `GET` or `POST`.
	Method *string `json:"Method,omitempty"`
	// The absolute url the post-event webhook request should be sent to.
	PostWebhookUrl *string `json:"PostWebhookUrl,omitempty"`
	// The absolute url the pre-event webhook request should be sent to.
	PreWebhookUrl *string `json:"PreWebhookUrl,omitempty"`
}

func (params *UpdateServiceWebhookConfigurationParams) SetFilters(Filters []string) *UpdateServiceWebhookConfigurationParams {
	params.Filters = &Filters
	return params
}
func (params *UpdateServiceWebhookConfigurationParams) SetMethod(Method string) *UpdateServiceWebhookConfigurationParams {
	params.Method = &Method
	return params
}
func (params *UpdateServiceWebhookConfigurationParams) SetPostWebhookUrl(PostWebhookUrl string) *UpdateServiceWebhookConfigurationParams {
	params.PostWebhookUrl = &PostWebhookUrl
	return params
}
func (params *UpdateServiceWebhookConfigurationParams) SetPreWebhookUrl(PreWebhookUrl string) *UpdateServiceWebhookConfigurationParams {
	params.PreWebhookUrl = &PreWebhookUrl
	return params
}

// Update a specific Webhook.
func (c *ApiService) UpdateServiceWebhookConfiguration(ChatServiceSid string, params *UpdateServiceWebhookConfigurationParams) (*ConversationsV1ServiceWebhookConfiguration, error) {
	path := "/v1/Services/{ChatServiceSid}/Configuration/Webhooks"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Filters != nil {
		for _, item := range *params.Filters {
			data.Add("Filters", item)
		}
	}
	if params != nil && params.Method != nil {
		data.Set("Method", *params.Method)
	}
	if params != nil && params.PostWebhookUrl != nil {
		data.Set("PostWebhookUrl", *params.PostWebhookUrl)
	}
	if params != nil && params.PreWebhookUrl != nil {
		data.Set("PreWebhookUrl", *params.PreWebhookUrl)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceWebhookConfiguration{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
