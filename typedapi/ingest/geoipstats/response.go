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

package geoipstats

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package geoipstats
//
// https://github.com/elastic/elasticsearch-specification/blob/656080dee792f93a849cd7fbf566f528f858a5ea/specification/ingest/geo_ip_stats/IngestGeoIpStatsResponse.ts#L24-L31

type Response struct {

	// Nodes Downloaded GeoIP2 databases for each node.
	Nodes map[string]types.GeoIpNodeDatabases `json:"nodes"`
	// Stats Download statistics for all GeoIP2 databases.
	Stats types.GeoIpDownloadStatistics `json:"stats"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Nodes: make(map[string]types.GeoIpNodeDatabases, 0),
	}
	return r
}
