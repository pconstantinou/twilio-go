/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateConfigurationAddress'
type CreateConfigurationAddressParams struct {
	// The unique address to be configured. The address can be a whatsapp address or phone number
	Address *string `json:"Address,omitempty"`
	// Conversation Service for the auto-created conversation. If not set, the conversation is created in the default service.
	AutoCreationConversationServiceSid *string `json:"AutoCreation.ConversationServiceSid,omitempty"`
	// Enable/Disable auto-creating conversations for messages to this address
	AutoCreationEnabled *bool `json:"AutoCreation.Enabled,omitempty"`
	// For type `studio`, the studio flow SID where the webhook should be sent to.
	AutoCreationStudioFlowSid *string `json:"AutoCreation.StudioFlowSid,omitempty"`
	// For type `studio`, number of times to retry the webhook request
	AutoCreationStudioRetryCount *int `json:"AutoCreation.StudioRetryCount,omitempty"`
	// Type of Auto Creation. Value can be one of `webhook`, `studio` or `default`.
	AutoCreationType *string `json:"AutoCreation.Type,omitempty"`
	// The list of events, firing webhook event for this Conversation. Values can be any of the following: `onMessageAdded`, `onMessageUpdated`, `onMessageRemoved`, `onConversationUpdated`, `onConversationStateUpdated`, `onConversationRemoved`, `onParticipantAdded`, `onParticipantUpdated`, `onParticipantRemoved`, `onDeliveryUpdated`
	AutoCreationWebhookFilters *[]string `json:"AutoCreation.WebhookFilters,omitempty"`
	// For type `webhook`, the HTTP method to be used when sending a webhook request.
	AutoCreationWebhookMethod *string `json:"AutoCreation.WebhookMethod,omitempty"`
	// For type `webhook`, the url for the webhook request.
	AutoCreationWebhookUrl *string `json:"AutoCreation.WebhookUrl,omitempty"`
	// The human-readable name of this configuration, limited to 256 characters. Optional.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Type of Address. Value can be `whatsapp` or `sms`.
	Type *string `json:"Type,omitempty"`
}

func (params *CreateConfigurationAddressParams) SetAddress(Address string) *CreateConfigurationAddressParams {
	params.Address = &Address
	return params
}
func (params *CreateConfigurationAddressParams) SetAutoCreationConversationServiceSid(AutoCreationConversationServiceSid string) *CreateConfigurationAddressParams {
	params.AutoCreationConversationServiceSid = &AutoCreationConversationServiceSid
	return params
}
func (params *CreateConfigurationAddressParams) SetAutoCreationEnabled(AutoCreationEnabled bool) *CreateConfigurationAddressParams {
	params.AutoCreationEnabled = &AutoCreationEnabled
	return params
}
func (params *CreateConfigurationAddressParams) SetAutoCreationStudioFlowSid(AutoCreationStudioFlowSid string) *CreateConfigurationAddressParams {
	params.AutoCreationStudioFlowSid = &AutoCreationStudioFlowSid
	return params
}
func (params *CreateConfigurationAddressParams) SetAutoCreationStudioRetryCount(AutoCreationStudioRetryCount int) *CreateConfigurationAddressParams {
	params.AutoCreationStudioRetryCount = &AutoCreationStudioRetryCount
	return params
}
func (params *CreateConfigurationAddressParams) SetAutoCreationType(AutoCreationType string) *CreateConfigurationAddressParams {
	params.AutoCreationType = &AutoCreationType
	return params
}
func (params *CreateConfigurationAddressParams) SetAutoCreationWebhookFilters(AutoCreationWebhookFilters []string) *CreateConfigurationAddressParams {
	params.AutoCreationWebhookFilters = &AutoCreationWebhookFilters
	return params
}
func (params *CreateConfigurationAddressParams) SetAutoCreationWebhookMethod(AutoCreationWebhookMethod string) *CreateConfigurationAddressParams {
	params.AutoCreationWebhookMethod = &AutoCreationWebhookMethod
	return params
}
func (params *CreateConfigurationAddressParams) SetAutoCreationWebhookUrl(AutoCreationWebhookUrl string) *CreateConfigurationAddressParams {
	params.AutoCreationWebhookUrl = &AutoCreationWebhookUrl
	return params
}
func (params *CreateConfigurationAddressParams) SetFriendlyName(FriendlyName string) *CreateConfigurationAddressParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateConfigurationAddressParams) SetType(Type string) *CreateConfigurationAddressParams {
	params.Type = &Type
	return params
}

// Create a new address configuration
func (c *ApiService) CreateConfigurationAddress(params *CreateConfigurationAddressParams) (*ConversationsV1ConfigurationAddress, error) {
	path := "/v1/Configuration/Addresses"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Address != nil {
		data.Set("Address", *params.Address)
	}
	if params != nil && params.AutoCreationConversationServiceSid != nil {
		data.Set("AutoCreation.ConversationServiceSid", *params.AutoCreationConversationServiceSid)
	}
	if params != nil && params.AutoCreationEnabled != nil {
		data.Set("AutoCreation.Enabled", fmt.Sprint(*params.AutoCreationEnabled))
	}
	if params != nil && params.AutoCreationStudioFlowSid != nil {
		data.Set("AutoCreation.StudioFlowSid", *params.AutoCreationStudioFlowSid)
	}
	if params != nil && params.AutoCreationStudioRetryCount != nil {
		data.Set("AutoCreation.StudioRetryCount", fmt.Sprint(*params.AutoCreationStudioRetryCount))
	}
	if params != nil && params.AutoCreationType != nil {
		data.Set("AutoCreation.Type", *params.AutoCreationType)
	}
	if params != nil && params.AutoCreationWebhookFilters != nil {
		for _, item := range *params.AutoCreationWebhookFilters {
			data.Add("AutoCreation.WebhookFilters", item)
		}
	}
	if params != nil && params.AutoCreationWebhookMethod != nil {
		data.Set("AutoCreation.WebhookMethod", *params.AutoCreationWebhookMethod)
	}
	if params != nil && params.AutoCreationWebhookUrl != nil {
		data.Set("AutoCreation.WebhookUrl", *params.AutoCreationWebhookUrl)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Type != nil {
		data.Set("Type", *params.Type)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConfigurationAddress{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an existing address configuration
func (c *ApiService) DeleteConfigurationAddress(Sid string) error {
	path := "/v1/Configuration/Addresses/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch an address configuration
func (c *ApiService) FetchConfigurationAddress(Sid string) (*ConversationsV1ConfigurationAddress, error) {
	path := "/v1/Configuration/Addresses/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConfigurationAddress{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListConfigurationAddress'
type ListConfigurationAddressParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListConfigurationAddressParams) SetPageSize(PageSize int) *ListConfigurationAddressParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListConfigurationAddressParams) SetLimit(Limit int) *ListConfigurationAddressParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ConfigurationAddress records from the API. Request is executed immediately.
func (c *ApiService) PageConfigurationAddress(params *ListConfigurationAddressParams, pageToken, pageNumber string) (*ListConfigurationAddressResponse, error) {
	path := "/v1/Configuration/Addresses"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConfigurationAddressResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ConfigurationAddress records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListConfigurationAddress(params *ListConfigurationAddressParams) ([]ConversationsV1ConfigurationAddress, error) {
	response, err := c.StreamConfigurationAddress(params)
	if err != nil {
		return nil, err
	}

	records := make([]ConversationsV1ConfigurationAddress, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams ConfigurationAddress records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamConfigurationAddress(params *ListConfigurationAddressParams) (chan ConversationsV1ConfigurationAddress, error) {
	if params == nil {
		params = &ListConfigurationAddressParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageConfigurationAddress(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan ConversationsV1ConfigurationAddress, 1)

	go func() {
		for response != nil {
			responseRecords := response.AddressConfigurations
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListConfigurationAddressResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListConfigurationAddressResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListConfigurationAddressResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConfigurationAddressResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateConfigurationAddress'
type UpdateConfigurationAddressParams struct {
	// Conversation Service for the auto-created conversation. If not set, the conversation is created in the default service.
	AutoCreationConversationServiceSid *string `json:"AutoCreation.ConversationServiceSid,omitempty"`
	// Enable/Disable auto-creating conversations for messages to this address
	AutoCreationEnabled *bool `json:"AutoCreation.Enabled,omitempty"`
	// For type `studio`, the studio flow SID where the webhook should be sent to.
	AutoCreationStudioFlowSid *string `json:"AutoCreation.StudioFlowSid,omitempty"`
	// For type `studio`, number of times to retry the webhook request
	AutoCreationStudioRetryCount *int `json:"AutoCreation.StudioRetryCount,omitempty"`
	// Type of Auto Creation. Value can be one of `webhook`, `studio` or `default`.
	AutoCreationType *string `json:"AutoCreation.Type,omitempty"`
	// The list of events, firing webhook event for this Conversation. Values can be any of the following: `onMessageAdded`, `onMessageUpdated`, `onMessageRemoved`, `onConversationUpdated`, `onConversationStateUpdated`, `onConversationRemoved`, `onParticipantAdded`, `onParticipantUpdated`, `onParticipantRemoved`, `onDeliveryUpdated`
	AutoCreationWebhookFilters *[]string `json:"AutoCreation.WebhookFilters,omitempty"`
	// For type `webhook`, the HTTP method to be used when sending a webhook request.
	AutoCreationWebhookMethod *string `json:"AutoCreation.WebhookMethod,omitempty"`
	// For type `webhook`, the url for the webhook request.
	AutoCreationWebhookUrl *string `json:"AutoCreation.WebhookUrl,omitempty"`
	// The human-readable name of this configuration, limited to 256 characters. Optional.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateConfigurationAddressParams) SetAutoCreationConversationServiceSid(AutoCreationConversationServiceSid string) *UpdateConfigurationAddressParams {
	params.AutoCreationConversationServiceSid = &AutoCreationConversationServiceSid
	return params
}
func (params *UpdateConfigurationAddressParams) SetAutoCreationEnabled(AutoCreationEnabled bool) *UpdateConfigurationAddressParams {
	params.AutoCreationEnabled = &AutoCreationEnabled
	return params
}
func (params *UpdateConfigurationAddressParams) SetAutoCreationStudioFlowSid(AutoCreationStudioFlowSid string) *UpdateConfigurationAddressParams {
	params.AutoCreationStudioFlowSid = &AutoCreationStudioFlowSid
	return params
}
func (params *UpdateConfigurationAddressParams) SetAutoCreationStudioRetryCount(AutoCreationStudioRetryCount int) *UpdateConfigurationAddressParams {
	params.AutoCreationStudioRetryCount = &AutoCreationStudioRetryCount
	return params
}
func (params *UpdateConfigurationAddressParams) SetAutoCreationType(AutoCreationType string) *UpdateConfigurationAddressParams {
	params.AutoCreationType = &AutoCreationType
	return params
}
func (params *UpdateConfigurationAddressParams) SetAutoCreationWebhookFilters(AutoCreationWebhookFilters []string) *UpdateConfigurationAddressParams {
	params.AutoCreationWebhookFilters = &AutoCreationWebhookFilters
	return params
}
func (params *UpdateConfigurationAddressParams) SetAutoCreationWebhookMethod(AutoCreationWebhookMethod string) *UpdateConfigurationAddressParams {
	params.AutoCreationWebhookMethod = &AutoCreationWebhookMethod
	return params
}
func (params *UpdateConfigurationAddressParams) SetAutoCreationWebhookUrl(AutoCreationWebhookUrl string) *UpdateConfigurationAddressParams {
	params.AutoCreationWebhookUrl = &AutoCreationWebhookUrl
	return params
}
func (params *UpdateConfigurationAddressParams) SetFriendlyName(FriendlyName string) *UpdateConfigurationAddressParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Update an existing address configuration
func (c *ApiService) UpdateConfigurationAddress(Sid string, params *UpdateConfigurationAddressParams) (*ConversationsV1ConfigurationAddress, error) {
	path := "/v1/Configuration/Addresses/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AutoCreationConversationServiceSid != nil {
		data.Set("AutoCreation.ConversationServiceSid", *params.AutoCreationConversationServiceSid)
	}
	if params != nil && params.AutoCreationEnabled != nil {
		data.Set("AutoCreation.Enabled", fmt.Sprint(*params.AutoCreationEnabled))
	}
	if params != nil && params.AutoCreationStudioFlowSid != nil {
		data.Set("AutoCreation.StudioFlowSid", *params.AutoCreationStudioFlowSid)
	}
	if params != nil && params.AutoCreationStudioRetryCount != nil {
		data.Set("AutoCreation.StudioRetryCount", fmt.Sprint(*params.AutoCreationStudioRetryCount))
	}
	if params != nil && params.AutoCreationType != nil {
		data.Set("AutoCreation.Type", *params.AutoCreationType)
	}
	if params != nil && params.AutoCreationWebhookFilters != nil {
		for _, item := range *params.AutoCreationWebhookFilters {
			data.Add("AutoCreation.WebhookFilters", item)
		}
	}
	if params != nil && params.AutoCreationWebhookMethod != nil {
		data.Set("AutoCreation.WebhookMethod", *params.AutoCreationWebhookMethod)
	}
	if params != nil && params.AutoCreationWebhookUrl != nil {
		data.Set("AutoCreation.WebhookUrl", *params.AutoCreationWebhookUrl)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConfigurationAddress{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
