/*
 * Twilio - Media
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
)

// Optional parameters for the method 'CreatePlayerStreamerPlaybackGrant'
type CreatePlayerStreamerPlaybackGrantParams struct {
	// The full origin URL where the livestream can be streamed. If this is not provided, it can be streamed from any domain.
	AccessControlAllowOrigin *string `json:"AccessControlAllowOrigin,omitempty"`
	// The time to live of the PlaybackGrant. Default value is 15 seconds. Maximum value is 60 seconds.
	Ttl *int `json:"Ttl,omitempty"`
}

func (params *CreatePlayerStreamerPlaybackGrantParams) SetAccessControlAllowOrigin(AccessControlAllowOrigin string) *CreatePlayerStreamerPlaybackGrantParams {
	params.AccessControlAllowOrigin = &AccessControlAllowOrigin
	return params
}
func (params *CreatePlayerStreamerPlaybackGrantParams) SetTtl(Ttl int) *CreatePlayerStreamerPlaybackGrantParams {
	params.Ttl = &Ttl
	return params
}

func (c *ApiService) CreatePlayerStreamerPlaybackGrant(Sid string, params *CreatePlayerStreamerPlaybackGrantParams) (*MediaV1PlayerStreamerPlaybackGrant, error) {
	path := "/v1/PlayerStreamers/{Sid}/PlaybackGrant"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AccessControlAllowOrigin != nil {
		data.Set("AccessControlAllowOrigin", *params.AccessControlAllowOrigin)
	}
	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MediaV1PlayerStreamerPlaybackGrant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// **This method is not enabled.** Returns a single PlaybackGrant resource identified by a SID.
func (c *ApiService) FetchPlayerStreamerPlaybackGrant(Sid string) (*MediaV1PlayerStreamerPlaybackGrant, error) {
	path := "/v1/PlayerStreamers/{Sid}/PlaybackGrant"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MediaV1PlayerStreamerPlaybackGrant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
