// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/656080dee792f93a849cd7fbf566f528f858a5ea

package grantapikey

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/apikeygranttype"
)

// Request holds the request body struct for the package grantapikey
//
// https://github.com/elastic/elasticsearch-specification/blob/656080dee792f93a849cd7fbf566f528f858a5ea/specification/security/grant_api_key/SecurityGrantApiKeyRequest.ts#L24-L38
type Request struct {
	AccessToken *string                         `json:"access_token,omitempty"`
	ApiKey      types.GrantApiKey               `json:"api_key"`
	GrantType   apikeygranttype.ApiKeyGrantType `json:"grant_type"`
	Password    *string                         `json:"password,omitempty"`
	RunAs       *string                         `json:"run_as,omitempty"`
	Username    *string                         `json:"username,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Grantapikey request: %w", err)
	}

	return &req, nil
}
