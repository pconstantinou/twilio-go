/*
 * Twilio - Serverless
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

// Retrieve a specific Function Version resource.
func (c *ApiService) FetchFunctionVersion(ServiceSid string, FunctionSid string, Sid string) (*ServerlessV1FunctionVersion, error) {
	path := "/v1/Services/{ServiceSid}/Functions/{FunctionSid}/Versions/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"FunctionSid"+"}", FunctionSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1FunctionVersion{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFunctionVersion'
type ListFunctionVersionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListFunctionVersionParams) SetPageSize(PageSize int) *ListFunctionVersionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListFunctionVersionParams) SetLimit(Limit int) *ListFunctionVersionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of FunctionVersion records from the API. Request is executed immediately.
func (c *ApiService) PageFunctionVersion(ServiceSid string, FunctionSid string, params *ListFunctionVersionParams, pageToken, pageNumber string) (*ListFunctionVersionResponse, error) {
	path := "/v1/Services/{ServiceSid}/Functions/{FunctionSid}/Versions"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"FunctionSid"+"}", FunctionSid, -1)

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

	ps := &ListFunctionVersionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists FunctionVersion records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListFunctionVersion(ServiceSid string, FunctionSid string, params *ListFunctionVersionParams) ([]ServerlessV1FunctionVersion, error) {
	response, err := c.StreamFunctionVersion(ServiceSid, FunctionSid, params)
	if err != nil {
		return nil, err
	}

	records := make([]ServerlessV1FunctionVersion, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams FunctionVersion records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamFunctionVersion(ServiceSid string, FunctionSid string, params *ListFunctionVersionParams) (chan ServerlessV1FunctionVersion, error) {
	if params == nil {
		params = &ListFunctionVersionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageFunctionVersion(ServiceSid, FunctionSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan ServerlessV1FunctionVersion, 1)

	go func() {
		for response != nil {
			responseRecords := response.FunctionVersions
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListFunctionVersionResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListFunctionVersionResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListFunctionVersionResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFunctionVersionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
