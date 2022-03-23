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
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateServiceConversationParticipant'
type CreateServiceConversationParticipantParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
	Attributes *string `json:"Attributes,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversation SDK to communicate. Limited to 256 characters.
	Identity *string `json:"Identity,omitempty"`
	// The address of the participant's device, e.g. a phone or WhatsApp number. Together with the Proxy address, this determines a participant uniquely. This field (with proxy_address) is only null when the participant is interacting from an SDK endpoint (see the 'identity' field).
	MessagingBindingAddress *string `json:"MessagingBinding.Address,omitempty"`
	// The address of the Twilio phone number that is used in Group MMS. Communication mask for the Conversation participant with Identity.
	MessagingBindingProjectedAddress *string `json:"MessagingBinding.ProjectedAddress,omitempty"`
	// The address of the Twilio phone number (or WhatsApp number) that the participant is in contact with. This field, together with participant address, is only null when the participant is interacting from an SDK endpoint (see the 'identity' field).
	MessagingBindingProxyAddress *string `json:"MessagingBinding.ProxyAddress,omitempty"`
	// The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant.
	RoleSid *string `json:"RoleSid,omitempty"`
}

func (params *CreateServiceConversationParticipantParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *CreateServiceConversationParticipantParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *CreateServiceConversationParticipantParams) SetAttributes(Attributes string) *CreateServiceConversationParticipantParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateServiceConversationParticipantParams) SetDateCreated(DateCreated time.Time) *CreateServiceConversationParticipantParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *CreateServiceConversationParticipantParams) SetDateUpdated(DateUpdated time.Time) *CreateServiceConversationParticipantParams {
	params.DateUpdated = &DateUpdated
	return params
}
func (params *CreateServiceConversationParticipantParams) SetIdentity(Identity string) *CreateServiceConversationParticipantParams {
	params.Identity = &Identity
	return params
}
func (params *CreateServiceConversationParticipantParams) SetMessagingBindingAddress(MessagingBindingAddress string) *CreateServiceConversationParticipantParams {
	params.MessagingBindingAddress = &MessagingBindingAddress
	return params
}
func (params *CreateServiceConversationParticipantParams) SetMessagingBindingProjectedAddress(MessagingBindingProjectedAddress string) *CreateServiceConversationParticipantParams {
	params.MessagingBindingProjectedAddress = &MessagingBindingProjectedAddress
	return params
}
func (params *CreateServiceConversationParticipantParams) SetMessagingBindingProxyAddress(MessagingBindingProxyAddress string) *CreateServiceConversationParticipantParams {
	params.MessagingBindingProxyAddress = &MessagingBindingProxyAddress
	return params
}
func (params *CreateServiceConversationParticipantParams) SetRoleSid(RoleSid string) *CreateServiceConversationParticipantParams {
	params.RoleSid = &RoleSid
	return params
}

// Add a new participant to the conversation in a specific service
func (c *ApiService) CreateServiceConversationParticipant(ChatServiceSid string, ConversationSid string, params *CreateServiceConversationParticipantParams) (*ConversationsV1ServiceConversationParticipant, error) {
	path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateUpdated != nil {
		data.Set("DateUpdated", fmt.Sprint((*params.DateUpdated).Format(time.RFC3339)))
	}
	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.MessagingBindingAddress != nil {
		data.Set("MessagingBinding.Address", *params.MessagingBindingAddress)
	}
	if params != nil && params.MessagingBindingProjectedAddress != nil {
		data.Set("MessagingBinding.ProjectedAddress", *params.MessagingBindingProjectedAddress)
	}
	if params != nil && params.MessagingBindingProxyAddress != nil {
		data.Set("MessagingBinding.ProxyAddress", *params.MessagingBindingProxyAddress)
	}
	if params != nil && params.RoleSid != nil {
		data.Set("RoleSid", *params.RoleSid)
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceConversationParticipant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteServiceConversationParticipant'
type DeleteServiceConversationParticipantParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
}

func (params *DeleteServiceConversationParticipantParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *DeleteServiceConversationParticipantParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}

// Remove a participant from the conversation
func (c *ApiService) DeleteServiceConversationParticipant(ChatServiceSid string, ConversationSid string, Sid string, params *DeleteServiceConversationParticipantParams) error {
	path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a participant of the conversation
func (c *ApiService) FetchServiceConversationParticipant(ChatServiceSid string, ConversationSid string, Sid string) (*ConversationsV1ServiceConversationParticipant, error) {
	path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceConversationParticipant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListServiceConversationParticipant'
type ListServiceConversationParticipantParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListServiceConversationParticipantParams) SetPageSize(PageSize int) *ListServiceConversationParticipantParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListServiceConversationParticipantParams) SetLimit(Limit int) *ListServiceConversationParticipantParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ServiceConversationParticipant records from the API. Request is executed immediately.
func (c *ApiService) PageServiceConversationParticipant(ChatServiceSid string, ConversationSid string, params *ListServiceConversationParticipantParams, pageToken, pageNumber string) (*ListServiceConversationParticipantResponse, error) {
	path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants"

	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

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

	ps := &ListServiceConversationParticipantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ServiceConversationParticipant records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListServiceConversationParticipant(ChatServiceSid string, ConversationSid string, params *ListServiceConversationParticipantParams) ([]ConversationsV1ServiceConversationParticipant, error) {
	response, err := c.StreamServiceConversationParticipant(ChatServiceSid, ConversationSid, params)
	if err != nil {
		return nil, err
	}

	records := make([]ConversationsV1ServiceConversationParticipant, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams ServiceConversationParticipant records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamServiceConversationParticipant(ChatServiceSid string, ConversationSid string, params *ListServiceConversationParticipantParams) (chan ConversationsV1ServiceConversationParticipant, error) {
	if params == nil {
		params = &ListServiceConversationParticipantParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageServiceConversationParticipant(ChatServiceSid, ConversationSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan ConversationsV1ServiceConversationParticipant, 1)

	go func() {
		for response != nil {
			responseRecords := response.Participants
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListServiceConversationParticipantResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListServiceConversationParticipantResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListServiceConversationParticipantResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceConversationParticipantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateServiceConversationParticipant'
type UpdateServiceConversationParticipantParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
	Attributes *string `json:"Attributes,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversation SDK to communicate. Limited to 256 characters.
	Identity *string `json:"Identity,omitempty"`
	// Index of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant.
	LastReadMessageIndex *int `json:"LastReadMessageIndex,omitempty"`
	// Timestamp of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant.
	LastReadTimestamp *string `json:"LastReadTimestamp,omitempty"`
	// The address of the Twilio phone number that is used in Group MMS. 'null' value will remove it.
	MessagingBindingProjectedAddress *string `json:"MessagingBinding.ProjectedAddress,omitempty"`
	// The address of the Twilio phone number that the participant is in contact with. 'null' value will remove it.
	MessagingBindingProxyAddress *string `json:"MessagingBinding.ProxyAddress,omitempty"`
	// The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant.
	RoleSid *string `json:"RoleSid,omitempty"`
}

func (params *UpdateServiceConversationParticipantParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *UpdateServiceConversationParticipantParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *UpdateServiceConversationParticipantParams) SetAttributes(Attributes string) *UpdateServiceConversationParticipantParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateServiceConversationParticipantParams) SetDateCreated(DateCreated time.Time) *UpdateServiceConversationParticipantParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *UpdateServiceConversationParticipantParams) SetDateUpdated(DateUpdated time.Time) *UpdateServiceConversationParticipantParams {
	params.DateUpdated = &DateUpdated
	return params
}
func (params *UpdateServiceConversationParticipantParams) SetIdentity(Identity string) *UpdateServiceConversationParticipantParams {
	params.Identity = &Identity
	return params
}
func (params *UpdateServiceConversationParticipantParams) SetLastReadMessageIndex(LastReadMessageIndex int) *UpdateServiceConversationParticipantParams {
	params.LastReadMessageIndex = &LastReadMessageIndex
	return params
}
func (params *UpdateServiceConversationParticipantParams) SetLastReadTimestamp(LastReadTimestamp string) *UpdateServiceConversationParticipantParams {
	params.LastReadTimestamp = &LastReadTimestamp
	return params
}
func (params *UpdateServiceConversationParticipantParams) SetMessagingBindingProjectedAddress(MessagingBindingProjectedAddress string) *UpdateServiceConversationParticipantParams {
	params.MessagingBindingProjectedAddress = &MessagingBindingProjectedAddress
	return params
}
func (params *UpdateServiceConversationParticipantParams) SetMessagingBindingProxyAddress(MessagingBindingProxyAddress string) *UpdateServiceConversationParticipantParams {
	params.MessagingBindingProxyAddress = &MessagingBindingProxyAddress
	return params
}
func (params *UpdateServiceConversationParticipantParams) SetRoleSid(RoleSid string) *UpdateServiceConversationParticipantParams {
	params.RoleSid = &RoleSid
	return params
}

// Update an existing participant in the conversation
func (c *ApiService) UpdateServiceConversationParticipant(ChatServiceSid string, ConversationSid string, Sid string, params *UpdateServiceConversationParticipantParams) (*ConversationsV1ServiceConversationParticipant, error) {
	path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateUpdated != nil {
		data.Set("DateUpdated", fmt.Sprint((*params.DateUpdated).Format(time.RFC3339)))
	}
	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.LastReadMessageIndex != nil {
		data.Set("LastReadMessageIndex", fmt.Sprint(*params.LastReadMessageIndex))
	}
	if params != nil && params.LastReadTimestamp != nil {
		data.Set("LastReadTimestamp", *params.LastReadTimestamp)
	}
	if params != nil && params.MessagingBindingProjectedAddress != nil {
		data.Set("MessagingBinding.ProjectedAddress", *params.MessagingBindingProjectedAddress)
	}
	if params != nil && params.MessagingBindingProxyAddress != nil {
		data.Set("MessagingBinding.ProxyAddress", *params.MessagingBindingProxyAddress)
	}
	if params != nil && params.RoleSid != nil {
		data.Set("RoleSid", *params.RoleSid)
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceConversationParticipant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
