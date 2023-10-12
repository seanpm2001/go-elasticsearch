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
// https://github.com/elastic/elasticsearch-specification/tree/5260ec5b7c899ab1a7939f752218cae07ef07dd7

package usage

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package usage
//
// https://github.com/elastic/elasticsearch-specification/blob/5260ec5b7c899ab1a7939f752218cae07ef07dd7/specification/nodes/usage/NodesUsageResponse.ts#L30-L32

type Response struct {
	ClusterName string `json:"cluster_name"`
	// NodeStats Contains statistics about the number of nodes selected by the request’s node
	// filters.
	NodeStats *types.NodeStatistics      `json:"_nodes,omitempty"`
	Nodes     map[string]types.NodeUsage `json:"nodes"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Nodes: make(map[string]types.NodeUsage, 0),
	}
	return r
}