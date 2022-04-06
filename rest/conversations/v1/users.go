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

// Optional parameters for the method 'CreateUser'
type CreateUserParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// The JSON Object string that stores application-specific data. If attributes have not been set, `{}` is returned.
	Attributes *string `json:"Attributes,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The application-defined string that uniquely identifies the resource's User within the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource). This value is often a username or an email address, and is case-sensitive.
	Identity *string `json:"Identity,omitempty"`
	// The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.
	RoleSid *string `json:"RoleSid,omitempty"`
}

func (params *CreateUserParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *CreateUserParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *CreateUserParams) SetAttributes(Attributes string) *CreateUserParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateUserParams) SetFriendlyName(FriendlyName string) *CreateUserParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateUserParams) SetIdentity(Identity string) *CreateUserParams {
	params.Identity = &Identity
	return params
}
func (params *CreateUserParams) SetRoleSid(RoleSid string) *CreateUserParams {
	params.RoleSid = &RoleSid
	return params
}

// Add a new conversation user to your account&#39;s default service
func (c *ApiService) CreateUser(params *CreateUserParams) (*ConversationsV1User, error) {
	path := "/v1/Users"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
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

	ps := &ConversationsV1User{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteUser'
type DeleteUserParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
}

func (params *DeleteUserParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *DeleteUserParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}

// Remove a conversation user from your account&#39;s default service
func (c *ApiService) DeleteUser(Sid string, params *DeleteUserParams) error {
	path := "/v1/Users/{Sid}"
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

// Fetch a conversation user from your account&#39;s default service
func (c *ApiService) FetchUser(Sid string) (*ConversationsV1User, error) {
	path := "/v1/Users/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1User{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListUser'
type ListUserParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListUserParams) SetPageSize(PageSize int) *ListUserParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListUserParams) SetLimit(Limit int) *ListUserParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of User records from the API. Request is executed immediately.
func (c *ApiService) PageUser(params *ListUserParams, pageToken, pageNumber string) (*ListUserResponse, error) {
	path := "/v1/Users"

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

	ps := &ListUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists User records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListUser(params *ListUserParams) ([]ConversationsV1User, error) {
	response, err := c.StreamUser(params)
	if err != nil {
		return nil, err
	}

	records := make([]ConversationsV1User, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams User records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamUser(params *ListUserParams) (chan ConversationsV1User, error) {
	if params == nil {
		params = &ListUserParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageUser(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan ConversationsV1User, 1)

	go func() {
		for response != nil {
			responseRecords := response.Users
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListUserResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListUserResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListUserResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateUser'
type UpdateUserParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// The JSON Object string that stores application-specific data. If attributes have not been set, `{}` is returned.
	Attributes *string `json:"Attributes,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.
	RoleSid *string `json:"RoleSid,omitempty"`
}

func (params *UpdateUserParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *UpdateUserParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *UpdateUserParams) SetAttributes(Attributes string) *UpdateUserParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateUserParams) SetFriendlyName(FriendlyName string) *UpdateUserParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateUserParams) SetRoleSid(RoleSid string) *UpdateUserParams {
	params.RoleSid = &RoleSid
	return params
}

// Update an existing conversation user in your account&#39;s default service
func (c *ApiService) UpdateUser(Sid string, params *UpdateUserParams) (*ConversationsV1User, error) {
	path := "/v1/Users/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
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

	ps := &ConversationsV1User{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
