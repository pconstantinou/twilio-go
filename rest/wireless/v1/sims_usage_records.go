/*
 * Twilio - Wireless
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

// Optional parameters for the method 'ListUsageRecord'
type ListUsageRecordParams struct {
	// Only include usage that occurred on or before this date, specified in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html). The default is the current time.
	End *time.Time `json:"End,omitempty"`
	// Only include usage that has occurred on or after this date, specified in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html). The default is one month before the `end` parameter value.
	Start *time.Time `json:"Start,omitempty"`
	// How to summarize the usage by time. Can be: `daily`, `hourly`, or `all`. The default is `all`. A value of `all` returns one Usage Record that describes the usage for the entire period.
	Granularity *string `json:"Granularity,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListUsageRecordParams) SetEnd(End time.Time) *ListUsageRecordParams {
	params.End = &End
	return params
}
func (params *ListUsageRecordParams) SetStart(Start time.Time) *ListUsageRecordParams {
	params.Start = &Start
	return params
}
func (params *ListUsageRecordParams) SetGranularity(Granularity string) *ListUsageRecordParams {
	params.Granularity = &Granularity
	return params
}
func (params *ListUsageRecordParams) SetPageSize(PageSize int) *ListUsageRecordParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListUsageRecordParams) SetLimit(Limit int) *ListUsageRecordParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of UsageRecord records from the API. Request is executed immediately.
func (c *ApiService) PageUsageRecord(SimSid string, params *ListUsageRecordParams, pageToken, pageNumber string) (*ListUsageRecordResponse, error) {
	path := "/v1/Sims/{SimSid}/UsageRecords"

	path = strings.Replace(path, "{"+"SimSid"+"}", SimSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.End != nil {
		data.Set("End", fmt.Sprint((*params.End).Format(time.RFC3339)))
	}
	if params != nil && params.Start != nil {
		data.Set("Start", fmt.Sprint((*params.Start).Format(time.RFC3339)))
	}
	if params != nil && params.Granularity != nil {
		data.Set("Granularity", *params.Granularity)
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

	ps := &ListUsageRecordResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists UsageRecord records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListUsageRecord(SimSid string, params *ListUsageRecordParams) ([]WirelessV1UsageRecord, error) {
	response, err := c.StreamUsageRecord(SimSid, params)
	if err != nil {
		return nil, err
	}

	records := make([]WirelessV1UsageRecord, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams UsageRecord records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamUsageRecord(SimSid string, params *ListUsageRecordParams) (chan WirelessV1UsageRecord, error) {
	if params == nil {
		params = &ListUsageRecordParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageUsageRecord(SimSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan WirelessV1UsageRecord, 1)

	go func() {
		for response != nil {
			responseRecords := response.UsageRecords
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListUsageRecordResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListUsageRecordResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListUsageRecordResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUsageRecordResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
