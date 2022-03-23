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
	"net/url"
	"strings"
)

// Optional parameters for the method 'FetchBalance'
type FetchBalanceParams struct {
	// The unique SID identifier of the Account.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchBalanceParams) SetPathAccountSid(PathAccountSid string) *FetchBalanceParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch the balance for an Account based on Account Sid. Balance changes may not be reflected immediately. Child accounts do not contain balance information
func (c *ApiService) FetchBalance(params *FetchBalanceParams) (*ApiV2010Balance, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Balance.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Balance{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
