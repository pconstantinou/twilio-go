/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Autopilot
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateSample'
type CreateSampleParams struct {
	// The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the new sample. For example: `en-US`.
	Language *string `json:"Language,omitempty"`
	// The communication channel from which the new sample was captured. Can be: `voice`, `sms`, `chat`, `alexa`, `google-assistant`, `slack`, or null if not included.
	SourceChannel *string `json:"SourceChannel,omitempty"`
	// The text example of how end users might express the task. The sample can contain [Field tag blocks](https://www.twilio.com/docs/autopilot/api/task-sample#field-tagging).
	TaggedText *string `json:"TaggedText,omitempty"`
}

func (params *CreateSampleParams) SetLanguage(Language string) *CreateSampleParams {
	params.Language = &Language
	return params
}
func (params *CreateSampleParams) SetSourceChannel(SourceChannel string) *CreateSampleParams {
	params.SourceChannel = &SourceChannel
	return params
}
func (params *CreateSampleParams) SetTaggedText(TaggedText string) *CreateSampleParams {
	params.TaggedText = &TaggedText
	return params
}

//
func (c *ApiService) CreateSample(AssistantSid string, TaskSid string, params *CreateSampleParams) (*AutopilotV1Sample, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Language != nil {
		data.Set("Language", *params.Language)
	}
	if params != nil && params.SourceChannel != nil {
		data.Set("SourceChannel", *params.SourceChannel)
	}
	if params != nil && params.TaggedText != nil {
		data.Set("TaggedText", *params.TaggedText)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1Sample{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//
func (c *ApiService) DeleteSample(AssistantSid string, TaskSid string, Sid string) error {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)
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

//
func (c *ApiService) FetchSample(AssistantSid string, TaskSid string, Sid string) (*AutopilotV1Sample, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1Sample{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSample'
type ListSampleParams struct {
	// The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the sample. For example: `en-US`.
	Language *string `json:"Language,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSampleParams) SetLanguage(Language string) *ListSampleParams {
	params.Language = &Language
	return params
}
func (params *ListSampleParams) SetPageSize(PageSize int) *ListSampleParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSampleParams) SetLimit(Limit int) *ListSampleParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Sample records from the API. Request is executed immediately.
func (c *ApiService) PageSample(AssistantSid string, TaskSid string, params *ListSampleParams, pageToken, pageNumber string) (*ListSampleResponse, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples"

	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Language != nil {
		data.Set("Language", *params.Language)
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

	ps := &ListSampleResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Sample records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSample(AssistantSid string, TaskSid string, params *ListSampleParams) ([]AutopilotV1Sample, error) {
	response, errors := c.StreamSample(AssistantSid, TaskSid, params)

	records := make([]AutopilotV1Sample, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Sample records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSample(AssistantSid string, TaskSid string, params *ListSampleParams) (chan AutopilotV1Sample, chan error) {
	if params == nil {
		params = &ListSampleParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan AutopilotV1Sample, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageSample(AssistantSid, TaskSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamSample(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamSample(response *ListSampleResponse, params *ListSampleParams, recordChannel chan AutopilotV1Sample, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Samples
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListSampleResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListSampleResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListSampleResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSampleResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSample'
type UpdateSampleParams struct {
	// The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) string that specifies the language used for the sample. For example: `en-US`.
	Language *string `json:"Language,omitempty"`
	// The communication channel from which the sample was captured. Can be: `voice`, `sms`, `chat`, `alexa`, `google-assistant`, `slack`, or null if not included.
	SourceChannel *string `json:"SourceChannel,omitempty"`
	// The text example of how end users might express the task. The sample can contain [Field tag blocks](https://www.twilio.com/docs/autopilot/api/task-sample#field-tagging).
	TaggedText *string `json:"TaggedText,omitempty"`
}

func (params *UpdateSampleParams) SetLanguage(Language string) *UpdateSampleParams {
	params.Language = &Language
	return params
}
func (params *UpdateSampleParams) SetSourceChannel(SourceChannel string) *UpdateSampleParams {
	params.SourceChannel = &SourceChannel
	return params
}
func (params *UpdateSampleParams) SetTaggedText(TaggedText string) *UpdateSampleParams {
	params.TaggedText = &TaggedText
	return params
}

//
func (c *ApiService) UpdateSample(AssistantSid string, TaskSid string, Sid string, params *UpdateSampleParams) (*AutopilotV1Sample, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Language != nil {
		data.Set("Language", *params.Language)
	}
	if params != nil && params.SourceChannel != nil {
		data.Set("SourceChannel", *params.SourceChannel)
	}
	if params != nil && params.TaggedText != nil {
		data.Set("TaggedText", *params.TaggedText)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1Sample{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
