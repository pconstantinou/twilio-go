/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Taskrouter
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

// Optional parameters for the method 'CreateWorker'
type CreateWorkerParams struct {
	// The SID of a valid Activity that will describe the new Worker's initial state. See [Activities](https://www.twilio.com/docs/taskrouter/api/activity) for more information. If not provided, the new Worker's initial state is the `default_activity_sid` configured on the Workspace.
	ActivitySid *string `json:"ActivitySid,omitempty"`
	// A valid JSON string that describes the new Worker. For example: `{ \\\"email\\\": \\\"Bob@example.com\\\", \\\"phone\\\": \\\"+5095551234\\\" }`. This data is passed to the `assignment_callback_url` when TaskRouter assigns a Task to the Worker. Defaults to {}.
	Attributes *string `json:"Attributes,omitempty"`
	// A descriptive string that you create to describe the new Worker. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateWorkerParams) SetActivitySid(ActivitySid string) *CreateWorkerParams {
	params.ActivitySid = &ActivitySid
	return params
}
func (params *CreateWorkerParams) SetAttributes(Attributes string) *CreateWorkerParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateWorkerParams) SetFriendlyName(FriendlyName string) *CreateWorkerParams {
	params.FriendlyName = &FriendlyName
	return params
}

//
func (c *ApiService) CreateWorker(WorkspaceSid string, params *CreateWorkerParams) (*TaskrouterV1Worker, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ActivitySid != nil {
		data.Set("ActivitySid", *params.ActivitySid)
	}
	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1Worker{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteWorker'
type DeleteWorkerParams struct {
	// The If-Match HTTP request header
	IfMatch *string `json:"If-Match,omitempty"`
}

func (params *DeleteWorkerParams) SetIfMatch(IfMatch string) *DeleteWorkerParams {
	params.IfMatch = &IfMatch
	return params
}

//
func (c *ApiService) DeleteWorker(WorkspaceSid string, Sid string, params *DeleteWorkerParams) error {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.IfMatch != nil {
		headers["If-Match"] = *params.IfMatch
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

//
func (c *ApiService) FetchWorker(WorkspaceSid string, Sid string) (*TaskrouterV1Worker, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1Worker{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListWorker'
type ListWorkerParams struct {
	// The `activity_name` of the Worker resources to read.
	ActivityName *string `json:"ActivityName,omitempty"`
	// The `activity_sid` of the Worker resources to read.
	ActivitySid *string `json:"ActivitySid,omitempty"`
	// Whether to return only Worker resources that are available or unavailable. Can be `true`, `1`, or `yes` to return Worker resources that are available, and `false`, or any value returns the Worker resources that are not available.
	Available *string `json:"Available,omitempty"`
	// The `friendly_name` of the Worker resources to read.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Filter by Workers that would match an expression on a TaskQueue. This is helpful for debugging which Workers would match a potential queue.
	TargetWorkersExpression *string `json:"TargetWorkersExpression,omitempty"`
	// The `friendly_name` of the TaskQueue that the Workers to read are eligible for.
	TaskQueueName *string `json:"TaskQueueName,omitempty"`
	// The SID of the TaskQueue that the Workers to read are eligible for.
	TaskQueueSid *string `json:"TaskQueueSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListWorkerParams) SetActivityName(ActivityName string) *ListWorkerParams {
	params.ActivityName = &ActivityName
	return params
}
func (params *ListWorkerParams) SetActivitySid(ActivitySid string) *ListWorkerParams {
	params.ActivitySid = &ActivitySid
	return params
}
func (params *ListWorkerParams) SetAvailable(Available string) *ListWorkerParams {
	params.Available = &Available
	return params
}
func (params *ListWorkerParams) SetFriendlyName(FriendlyName string) *ListWorkerParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListWorkerParams) SetTargetWorkersExpression(TargetWorkersExpression string) *ListWorkerParams {
	params.TargetWorkersExpression = &TargetWorkersExpression
	return params
}
func (params *ListWorkerParams) SetTaskQueueName(TaskQueueName string) *ListWorkerParams {
	params.TaskQueueName = &TaskQueueName
	return params
}
func (params *ListWorkerParams) SetTaskQueueSid(TaskQueueSid string) *ListWorkerParams {
	params.TaskQueueSid = &TaskQueueSid
	return params
}
func (params *ListWorkerParams) SetPageSize(PageSize int) *ListWorkerParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListWorkerParams) SetLimit(Limit int) *ListWorkerParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Worker records from the API. Request is executed immediately.
func (c *ApiService) PageWorker(WorkspaceSid string, params *ListWorkerParams, pageToken, pageNumber string) (*ListWorkerResponse, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers"

	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ActivityName != nil {
		data.Set("ActivityName", *params.ActivityName)
	}
	if params != nil && params.ActivitySid != nil {
		data.Set("ActivitySid", *params.ActivitySid)
	}
	if params != nil && params.Available != nil {
		data.Set("Available", *params.Available)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.TargetWorkersExpression != nil {
		data.Set("TargetWorkersExpression", *params.TargetWorkersExpression)
	}
	if params != nil && params.TaskQueueName != nil {
		data.Set("TaskQueueName", *params.TaskQueueName)
	}
	if params != nil && params.TaskQueueSid != nil {
		data.Set("TaskQueueSid", *params.TaskQueueSid)
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

	ps := &ListWorkerResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Worker records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListWorker(WorkspaceSid string, params *ListWorkerParams) ([]TaskrouterV1Worker, error) {
	response, errors := c.StreamWorker(WorkspaceSid, params)

	records := make([]TaskrouterV1Worker, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Worker records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamWorker(WorkspaceSid string, params *ListWorkerParams) (chan TaskrouterV1Worker, chan error) {
	if params == nil {
		params = &ListWorkerParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan TaskrouterV1Worker, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageWorker(WorkspaceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamWorker(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamWorker(response *ListWorkerResponse, params *ListWorkerParams, recordChannel chan TaskrouterV1Worker, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Workers
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListWorkerResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListWorkerResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListWorkerResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListWorkerResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateWorker'
type UpdateWorkerParams struct {
	// The If-Match HTTP request header
	IfMatch *string `json:"If-Match,omitempty"`
	// The SID of a valid Activity that will describe the Worker's initial state. See [Activities](https://www.twilio.com/docs/taskrouter/api/activity) for more information.
	ActivitySid *string `json:"ActivitySid,omitempty"`
	// The JSON string that describes the Worker. For example: `{ \\\"email\\\": \\\"Bob@example.com\\\", \\\"phone\\\": \\\"+5095551234\\\" }`. This data is passed to the `assignment_callback_url` when TaskRouter assigns a Task to the Worker. Defaults to {}.
	Attributes *string `json:"Attributes,omitempty"`
	// A descriptive string that you create to describe the Worker. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether to reject the Worker's pending reservations. This option is only valid if the Worker's new [Activity](https://www.twilio.com/docs/taskrouter/api/activity) resource has its `availability` property set to `False`.
	RejectPendingReservations *bool `json:"RejectPendingReservations,omitempty"`
}

func (params *UpdateWorkerParams) SetIfMatch(IfMatch string) *UpdateWorkerParams {
	params.IfMatch = &IfMatch
	return params
}
func (params *UpdateWorkerParams) SetActivitySid(ActivitySid string) *UpdateWorkerParams {
	params.ActivitySid = &ActivitySid
	return params
}
func (params *UpdateWorkerParams) SetAttributes(Attributes string) *UpdateWorkerParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateWorkerParams) SetFriendlyName(FriendlyName string) *UpdateWorkerParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateWorkerParams) SetRejectPendingReservations(RejectPendingReservations bool) *UpdateWorkerParams {
	params.RejectPendingReservations = &RejectPendingReservations
	return params
}

//
func (c *ApiService) UpdateWorker(WorkspaceSid string, Sid string, params *UpdateWorkerParams) (*TaskrouterV1Worker, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ActivitySid != nil {
		data.Set("ActivitySid", *params.ActivitySid)
	}
	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.RejectPendingReservations != nil {
		data.Set("RejectPendingReservations", fmt.Sprint(*params.RejectPendingReservations))
	}

	if params != nil && params.IfMatch != nil {
		headers["If-Match"] = *params.IfMatch
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1Worker{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
