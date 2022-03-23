/*
 * Twilio - Video
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

// Optional parameters for the method 'CreateCompositionHook'
type CreateCompositionHookParams struct {
	// An array of track names from the same group room to merge into the compositions created by the composition hook. Can include zero or more track names. A composition triggered by the composition hook includes all audio sources specified in `audio_sources` except those specified in `audio_sources_excluded`. The track names in this parameter can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, `student*` includes tracks named `student` as well as `studentTeam`.
	AudioSources *[]string `json:"AudioSources,omitempty"`
	// An array of track names to exclude. A composition triggered by the composition hook includes all audio sources specified in `audio_sources` except for those specified in `audio_sources_excluded`. The track names in this parameter can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, `student*` excludes `student` as well as `studentTeam`. This parameter can also be empty.
	AudioSourcesExcluded *[]string `json:"AudioSourcesExcluded,omitempty"`
	// Whether the composition hook is active. When `true`, the composition hook will be triggered for every completed Group Room in the account. When `false`, the composition hook will never be triggered.
	Enabled *bool `json:"Enabled,omitempty"`
	// The container format of the media files used by the compositions created by the composition hook. Can be: `mp4` or `webm` and the default is `webm`. If `mp4` or `webm`, `audio_sources` must have one or more tracks and/or a `video_layout` element must contain a valid `video_sources` list, otherwise an error occurs.
	Format *string `json:"Format,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to  100 characters long and it must be unique within the account.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// A string that describes the columns (width) and rows (height) of the generated composed video in pixels. Defaults to `640x480`.  The string's format is `{width}x{height}` where:   * 16 <= `{width}` <= 1280 * 16 <= `{height}` <= 1280 * `{width}` * `{height}` <= 921,600  Typical values are:   * HD = `1280x720` * PAL = `1024x576` * VGA = `640x480` * CIF = `320x240`  Note that the `resolution` imposes an aspect ratio to the resulting composition. When the original video tracks are constrained by the aspect ratio, they are scaled to fit. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info.
	Resolution *string `json:"Resolution,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application on every composition event. If not provided, status callback events will not be dispatched.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The HTTP method we should use to call `status_callback`. Can be: `POST` or `GET` and the default is `POST`.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// Whether to clip the intervals where there is no active media in the Compositions triggered by the composition hook. The default is `true`. Compositions with `trim` enabled are shorter when the Room is created and no Participant joins for a while as well as if all the Participants leave the room and join later, because those gaps will be removed. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info.
	Trim *bool `json:"Trim,omitempty"`
	// An object that describes the video layout of the composition hook in terms of regions. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info.
	VideoLayout *map[string]interface{} `json:"VideoLayout,omitempty"`
}

func (params *CreateCompositionHookParams) SetAudioSources(AudioSources []string) *CreateCompositionHookParams {
	params.AudioSources = &AudioSources
	return params
}
func (params *CreateCompositionHookParams) SetAudioSourcesExcluded(AudioSourcesExcluded []string) *CreateCompositionHookParams {
	params.AudioSourcesExcluded = &AudioSourcesExcluded
	return params
}
func (params *CreateCompositionHookParams) SetEnabled(Enabled bool) *CreateCompositionHookParams {
	params.Enabled = &Enabled
	return params
}
func (params *CreateCompositionHookParams) SetFormat(Format string) *CreateCompositionHookParams {
	params.Format = &Format
	return params
}
func (params *CreateCompositionHookParams) SetFriendlyName(FriendlyName string) *CreateCompositionHookParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateCompositionHookParams) SetResolution(Resolution string) *CreateCompositionHookParams {
	params.Resolution = &Resolution
	return params
}
func (params *CreateCompositionHookParams) SetStatusCallback(StatusCallback string) *CreateCompositionHookParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *CreateCompositionHookParams) SetStatusCallbackMethod(StatusCallbackMethod string) *CreateCompositionHookParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}
func (params *CreateCompositionHookParams) SetTrim(Trim bool) *CreateCompositionHookParams {
	params.Trim = &Trim
	return params
}
func (params *CreateCompositionHookParams) SetVideoLayout(VideoLayout map[string]interface{}) *CreateCompositionHookParams {
	params.VideoLayout = &VideoLayout
	return params
}

func (c *ApiService) CreateCompositionHook(params *CreateCompositionHookParams) (*VideoV1CompositionHook, error) {
	path := "/v1/CompositionHooks"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AudioSources != nil {
		for _, item := range *params.AudioSources {
			data.Add("AudioSources", item)
		}
	}
	if params != nil && params.AudioSourcesExcluded != nil {
		for _, item := range *params.AudioSourcesExcluded {
			data.Add("AudioSourcesExcluded", item)
		}
	}
	if params != nil && params.Enabled != nil {
		data.Set("Enabled", fmt.Sprint(*params.Enabled))
	}
	if params != nil && params.Format != nil {
		data.Set("Format", *params.Format)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Resolution != nil {
		data.Set("Resolution", *params.Resolution)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}
	if params != nil && params.Trim != nil {
		data.Set("Trim", fmt.Sprint(*params.Trim))
	}
	if params != nil && params.VideoLayout != nil {
		v, err := json.Marshal(params.VideoLayout)

		if err != nil {
			return nil, err
		}

		data.Set("VideoLayout", string(v))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1CompositionHook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a Recording CompositionHook resource identified by a &#x60;CompositionHook SID&#x60;.
func (c *ApiService) DeleteCompositionHook(Sid string) error {
	path := "/v1/CompositionHooks/{Sid}"
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

// Returns a single CompositionHook resource identified by a CompositionHook SID.
func (c *ApiService) FetchCompositionHook(Sid string) (*VideoV1CompositionHook, error) {
	path := "/v1/CompositionHooks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1CompositionHook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCompositionHook'
type ListCompositionHookParams struct {
	// Read only CompositionHook resources with an `enabled` value that matches this parameter.
	Enabled *bool `json:"Enabled,omitempty"`
	// Read only CompositionHook resources created on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone.
	DateCreatedAfter *time.Time `json:"DateCreatedAfter,omitempty"`
	// Read only CompositionHook resources created before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone.
	DateCreatedBefore *time.Time `json:"DateCreatedBefore,omitempty"`
	// Read only CompositionHook resources with friendly names that match this string. The match is not case sensitive and can include asterisk `*` characters as wildcard match.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListCompositionHookParams) SetEnabled(Enabled bool) *ListCompositionHookParams {
	params.Enabled = &Enabled
	return params
}
func (params *ListCompositionHookParams) SetDateCreatedAfter(DateCreatedAfter time.Time) *ListCompositionHookParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListCompositionHookParams) SetDateCreatedBefore(DateCreatedBefore time.Time) *ListCompositionHookParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *ListCompositionHookParams) SetFriendlyName(FriendlyName string) *ListCompositionHookParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListCompositionHookParams) SetPageSize(PageSize int) *ListCompositionHookParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListCompositionHookParams) SetLimit(Limit int) *ListCompositionHookParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of CompositionHook records from the API. Request is executed immediately.
func (c *ApiService) PageCompositionHook(params *ListCompositionHookParams, pageToken, pageNumber string) (*ListCompositionHookResponse, error) {
	path := "/v1/CompositionHooks"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Enabled != nil {
		data.Set("Enabled", fmt.Sprint(*params.Enabled))
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreatedAfter", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreatedBefore", fmt.Sprint((*params.DateCreatedBefore).Format(time.RFC3339)))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
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

	ps := &ListCompositionHookResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists CompositionHook records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCompositionHook(params *ListCompositionHookParams) ([]VideoV1CompositionHook, error) {
	response, err := c.StreamCompositionHook(params)
	if err != nil {
		return nil, err
	}

	records := make([]VideoV1CompositionHook, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams CompositionHook records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCompositionHook(params *ListCompositionHookParams) (chan VideoV1CompositionHook, error) {
	if params == nil {
		params = &ListCompositionHookParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageCompositionHook(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan VideoV1CompositionHook, 1)

	go func() {
		for response != nil {
			responseRecords := response.CompositionHooks
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListCompositionHookResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListCompositionHookResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListCompositionHookResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCompositionHookResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateCompositionHook'
type UpdateCompositionHookParams struct {
	// An array of track names from the same group room to merge into the compositions created by the composition hook. Can include zero or more track names. A composition triggered by the composition hook includes all audio sources specified in `audio_sources` except those specified in `audio_sources_excluded`. The track names in this parameter can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, `student*` includes tracks named `student` as well as `studentTeam`.
	AudioSources *[]string `json:"AudioSources,omitempty"`
	// An array of track names to exclude. A composition triggered by the composition hook includes all audio sources specified in `audio_sources` except for those specified in `audio_sources_excluded`. The track names in this parameter can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, `student*` excludes `student` as well as `studentTeam`. This parameter can also be empty.
	AudioSourcesExcluded *[]string `json:"AudioSourcesExcluded,omitempty"`
	// Whether the composition hook is active. When `true`, the composition hook will be triggered for every completed Group Room in the account. When `false`, the composition hook never triggers.
	Enabled *bool `json:"Enabled,omitempty"`
	// The container format of the media files used by the compositions created by the composition hook. Can be: `mp4` or `webm` and the default is `webm`. If `mp4` or `webm`, `audio_sources` must have one or more tracks and/or a `video_layout` element must contain a valid `video_sources` list, otherwise an error occurs.
	Format *string `json:"Format,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to  100 characters long and it must be unique within the account.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// A string that describes the columns (width) and rows (height) of the generated composed video in pixels. Defaults to `640x480`.  The string's format is `{width}x{height}` where:   * 16 <= `{width}` <= 1280 * 16 <= `{height}` <= 1280 * `{width}` * `{height}` <= 921,600  Typical values are:   * HD = `1280x720` * PAL = `1024x576` * VGA = `640x480` * CIF = `320x240`  Note that the `resolution` imposes an aspect ratio to the resulting composition. When the original video tracks are constrained by the aspect ratio, they are scaled to fit. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info.
	Resolution *string `json:"Resolution,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application on every composition event. If not provided, status callback events will not be dispatched.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The HTTP method we should use to call `status_callback`. Can be: `POST` or `GET` and the default is `POST`.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// Whether to clip the intervals where there is no active media in the compositions triggered by the composition hook. The default is `true`. Compositions with `trim` enabled are shorter when the Room is created and no Participant joins for a while as well as if all the Participants leave the room and join later, because those gaps will be removed. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info.
	Trim *bool `json:"Trim,omitempty"`
	// A JSON object that describes the video layout of the composition hook in terms of regions. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info.
	VideoLayout *map[string]interface{} `json:"VideoLayout,omitempty"`
}

func (params *UpdateCompositionHookParams) SetAudioSources(AudioSources []string) *UpdateCompositionHookParams {
	params.AudioSources = &AudioSources
	return params
}
func (params *UpdateCompositionHookParams) SetAudioSourcesExcluded(AudioSourcesExcluded []string) *UpdateCompositionHookParams {
	params.AudioSourcesExcluded = &AudioSourcesExcluded
	return params
}
func (params *UpdateCompositionHookParams) SetEnabled(Enabled bool) *UpdateCompositionHookParams {
	params.Enabled = &Enabled
	return params
}
func (params *UpdateCompositionHookParams) SetFormat(Format string) *UpdateCompositionHookParams {
	params.Format = &Format
	return params
}
func (params *UpdateCompositionHookParams) SetFriendlyName(FriendlyName string) *UpdateCompositionHookParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateCompositionHookParams) SetResolution(Resolution string) *UpdateCompositionHookParams {
	params.Resolution = &Resolution
	return params
}
func (params *UpdateCompositionHookParams) SetStatusCallback(StatusCallback string) *UpdateCompositionHookParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *UpdateCompositionHookParams) SetStatusCallbackMethod(StatusCallbackMethod string) *UpdateCompositionHookParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}
func (params *UpdateCompositionHookParams) SetTrim(Trim bool) *UpdateCompositionHookParams {
	params.Trim = &Trim
	return params
}
func (params *UpdateCompositionHookParams) SetVideoLayout(VideoLayout map[string]interface{}) *UpdateCompositionHookParams {
	params.VideoLayout = &VideoLayout
	return params
}

func (c *ApiService) UpdateCompositionHook(Sid string, params *UpdateCompositionHookParams) (*VideoV1CompositionHook, error) {
	path := "/v1/CompositionHooks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AudioSources != nil {
		for _, item := range *params.AudioSources {
			data.Add("AudioSources", item)
		}
	}
	if params != nil && params.AudioSourcesExcluded != nil {
		for _, item := range *params.AudioSourcesExcluded {
			data.Add("AudioSourcesExcluded", item)
		}
	}
	if params != nil && params.Enabled != nil {
		data.Set("Enabled", fmt.Sprint(*params.Enabled))
	}
	if params != nil && params.Format != nil {
		data.Set("Format", *params.Format)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Resolution != nil {
		data.Set("Resolution", *params.Resolution)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}
	if params != nil && params.Trim != nil {
		data.Set("Trim", fmt.Sprint(*params.Trim))
	}
	if params != nil && params.VideoLayout != nil {
		v, err := json.Marshal(params.VideoLayout)

		if err != nil {
			return nil, err
		}

		data.Set("VideoLayout", string(v))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1CompositionHook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
