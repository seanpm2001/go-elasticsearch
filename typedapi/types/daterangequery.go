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
// https://github.com/elastic/elasticsearch-specification/tree/4ca0cc05d3ae3fa06c2cd7be91905b656a474334


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/rangerelation"
)

// DateRangeQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ca0cc05d3ae3fa06c2cd7be91905b656a474334/specification/_types/query_dsl/term.ts#L72-L81
type DateRangeQuery struct {
	Boost      *float32                     `json:"boost,omitempty"`
	Format     *string                      `json:"format,omitempty"`
	From       string                       `json:"from,omitempty"`
	Gt         *string                      `json:"gt,omitempty"`
	Gte        *string                      `json:"gte,omitempty"`
	Lt         *string                      `json:"lt,omitempty"`
	Lte        *string                      `json:"lte,omitempty"`
	QueryName_ *string                      `json:"_name,omitempty"`
	Relation   *rangerelation.RangeRelation `json:"relation,omitempty"`
	TimeZone   *string                      `json:"time_zone,omitempty"`
	To         string                       `json:"to,omitempty"`
}

// NewDateRangeQuery returns a DateRangeQuery.
func NewDateRangeQuery() *DateRangeQuery {
	r := &DateRangeQuery{}

	return r
}