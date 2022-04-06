/*
 * Twilio - Ip_messaging
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

func (c *ApiService) DeleteBinding(ServiceSid string, Sid string) error {
	path := "/v2/Services/{ServiceSid}/Bindings/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
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

func (c *ApiService) FetchBinding(ServiceSid string, Sid string) (*IpMessagingV2Binding, error) {
	path := "/v2/Services/{ServiceSid}/Bindings/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV2Binding{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListBinding'
type ListBindingParams struct {
	//
	BindingType *[]string `json:"BindingType,omitempty"`
	//
	Identity *[]string `json:"Identity,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListBindingParams) SetBindingType(BindingType []string) *ListBindingParams {
	params.BindingType = &BindingType
	return params
}
func (params *ListBindingParams) SetIdentity(Identity []string) *ListBindingParams {
	params.Identity = &Identity
	return params
}
func (params *ListBindingParams) SetPageSize(PageSize int) *ListBindingParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListBindingParams) SetLimit(Limit int) *ListBindingParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Binding records from the API. Request is executed immediately.
func (c *ApiService) PageBinding(ServiceSid string, params *ListBindingParams, pageToken, pageNumber string) (*ListBindingResponse, error) {
	path := "/v2/Services/{ServiceSid}/Bindings"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.BindingType != nil {
		for _, item := range *params.BindingType {
			data.Add("BindingType", item)
		}
	}
	if params != nil && params.Identity != nil {
		for _, item := range *params.Identity {
			data.Add("Identity", item)
		}
	}
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

	ps := &ListBindingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Binding records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListBinding(ServiceSid string, params *ListBindingParams) ([]IpMessagingV2Binding, error) {
	response, err := c.StreamBinding(ServiceSid, params)
	if err != nil {
		return nil, err
	}

	records := make([]IpMessagingV2Binding, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams Binding records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamBinding(ServiceSid string, params *ListBindingParams) (chan IpMessagingV2Binding, error) {
	if params == nil {
		params = &ListBindingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageBinding(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan IpMessagingV2Binding, 1)

	go func() {
		for response != nil {
			responseRecords := response.Bindings
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListBindingResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListBindingResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListBindingResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListBindingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
