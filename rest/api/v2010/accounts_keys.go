/*
 * Twilio - Api
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

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateNewKey'
type CreateNewKeyParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Key resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateNewKeyParams) SetPathAccountSid(PathAccountSid string) *CreateNewKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateNewKeyParams) SetFriendlyName(FriendlyName string) *CreateNewKeyParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) CreateNewKey(params *CreateNewKeyParams) (*ApiV2010NewKey, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Keys.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010NewKey{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteKey'
type DeleteKeyParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteKeyParams) SetPathAccountSid(PathAccountSid string) *DeleteKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) DeleteKey(Sid string, params *DeleteKeyParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Keys/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
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

// Optional parameters for the method 'FetchKey'
type FetchKeyParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchKeyParams) SetPathAccountSid(PathAccountSid string) *FetchKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) FetchKey(Sid string, params *FetchKeyParams) (*ApiV2010Key, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Keys/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Key{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListKey'
type ListKeyParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListKeyParams) SetPathAccountSid(PathAccountSid string) *ListKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListKeyParams) SetPageSize(PageSize int) *ListKeyParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListKeyParams) SetLimit(Limit int) *ListKeyParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Key records from the API. Request is executed immediately.
func (c *ApiService) PageKey(params *ListKeyParams, pageToken, pageNumber string) (*ListKeyResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Keys.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

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

	ps := &ListKeyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Key records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListKey(params *ListKeyParams) ([]ApiV2010Key, error) {
	response, err := c.StreamKey(params)
	if err != nil {
		return nil, err
	}

	records := make([]ApiV2010Key, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams Key records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamKey(params *ListKeyParams) (chan ApiV2010Key, error) {
	if params == nil {
		params = &ListKeyParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageKey(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010Key, 1)

	go func() {
		for response != nil {
			responseRecords := response.Keys
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListKeyResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListKeyResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListKeyResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListKeyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateKey'
type UpdateKeyParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateKeyParams) SetPathAccountSid(PathAccountSid string) *UpdateKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateKeyParams) SetFriendlyName(FriendlyName string) *UpdateKeyParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) UpdateKey(Sid string, params *UpdateKeyParams) (*ApiV2010Key, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Keys/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Key{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
