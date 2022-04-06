/*
 * Twilio - Events
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

// Fetch a specific schema and version.
func (c *ApiService) FetchSchemaVersion(Id string, SchemaVersion int) (*EventsV1SchemaVersion, error) {
	path := "/v1/Schemas/{Id}/Versions/{SchemaVersion}"
	path = strings.Replace(path, "{"+"Id"+"}", Id, -1)
	path = strings.Replace(path, "{"+"SchemaVersion"+"}", fmt.Sprint(SchemaVersion), -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1SchemaVersion{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSchemaVersion'
type ListSchemaVersionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSchemaVersionParams) SetPageSize(PageSize int) *ListSchemaVersionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSchemaVersionParams) SetLimit(Limit int) *ListSchemaVersionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SchemaVersion records from the API. Request is executed immediately.
func (c *ApiService) PageSchemaVersion(Id string, params *ListSchemaVersionParams, pageToken, pageNumber string) (*ListSchemaVersionResponse, error) {
	path := "/v1/Schemas/{Id}/Versions"

	path = strings.Replace(path, "{"+"Id"+"}", Id, -1)

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

	ps := &ListSchemaVersionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SchemaVersion records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSchemaVersion(Id string, params *ListSchemaVersionParams) ([]EventsV1SchemaVersion, error) {
	response, err := c.StreamSchemaVersion(Id, params)
	if err != nil {
		return nil, err
	}

	records := make([]EventsV1SchemaVersion, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams SchemaVersion records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSchemaVersion(Id string, params *ListSchemaVersionParams) (chan EventsV1SchemaVersion, error) {
	if params == nil {
		params = &ListSchemaVersionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSchemaVersion(Id, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan EventsV1SchemaVersion, 1)

	go func() {
		for response != nil {
			responseRecords := response.SchemaVersions
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListSchemaVersionResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSchemaVersionResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSchemaVersionResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSchemaVersionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
