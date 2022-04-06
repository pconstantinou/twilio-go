/*
 * Twilio - Monitor
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
	"time"

	"github.com/twilio/twilio-go/client"
)

func (c *ApiService) FetchAlert(Sid string) (*MonitorV1AlertInstance, error) {
	path := "/v1/Alerts/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MonitorV1AlertInstance{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListAlert'
type ListAlertParams struct {
	// Only show alerts for this log-level.  Can be: `error`, `warning`, `notice`, or `debug`.
	LogLevel *string `json:"LogLevel,omitempty"`
	// Only include alerts that occurred on or after this date and time. Specify the date and time in GMT and format as `YYYY-MM-DD` or `YYYY-MM-DDThh:mm:ssZ`. Queries for alerts older than 30 days are not supported.
	StartDate *time.Time `json:"StartDate,omitempty"`
	// Only include alerts that occurred on or before this date and time. Specify the date and time in GMT and format as `YYYY-MM-DD` or `YYYY-MM-DDThh:mm:ssZ`. Queries for alerts older than 30 days are not supported.
	EndDate *time.Time `json:"EndDate,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListAlertParams) SetLogLevel(LogLevel string) *ListAlertParams {
	params.LogLevel = &LogLevel
	return params
}
func (params *ListAlertParams) SetStartDate(StartDate time.Time) *ListAlertParams {
	params.StartDate = &StartDate
	return params
}
func (params *ListAlertParams) SetEndDate(EndDate time.Time) *ListAlertParams {
	params.EndDate = &EndDate
	return params
}
func (params *ListAlertParams) SetPageSize(PageSize int) *ListAlertParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListAlertParams) SetLimit(Limit int) *ListAlertParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Alert records from the API. Request is executed immediately.
func (c *ApiService) PageAlert(params *ListAlertParams, pageToken, pageNumber string) (*ListAlertResponse, error) {
	path := "/v1/Alerts"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.LogLevel != nil {
		data.Set("LogLevel", *params.LogLevel)
	}
	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint((*params.StartDate).Format(time.RFC3339)))
	}
	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint((*params.EndDate).Format(time.RFC3339)))
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

	ps := &ListAlertResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Alert records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListAlert(params *ListAlertParams) ([]MonitorV1Alert, error) {
	response, err := c.StreamAlert(params)
	if err != nil {
		return nil, err
	}

	records := make([]MonitorV1Alert, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams Alert records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamAlert(params *ListAlertParams) (chan MonitorV1Alert, error) {
	if params == nil {
		params = &ListAlertParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageAlert(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan MonitorV1Alert, 1)

	go func() {
		for response != nil {
			responseRecords := response.Alerts
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListAlertResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListAlertResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListAlertResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAlertResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
