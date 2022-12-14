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
// https://github.com/elastic/elasticsearch-specification/tree/d63a0e35ee85d84c83d0612ff3c0641a7a1e7e2e


package types

// PercentilesAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/d63a0e35ee85d84c83d0612ff3c0641a7a1e7e2e/specification/_types/aggregations/metric.ts#L112-L117
type PercentilesAggregation struct {
	Field    *string    `json:"field,omitempty"`
	Format   *string    `json:"format,omitempty"`
	Hdr      *HdrMethod `json:"hdr,omitempty"`
	Keyed    *bool      `json:"keyed,omitempty"`
	Missing  *Missing   `json:"missing,omitempty"`
	Percents []float64  `json:"percents,omitempty"`
	Script   *Script    `json:"script,omitempty"`
	Tdigest  *TDigest   `json:"tdigest,omitempty"`
}

// NewPercentilesAggregation returns a PercentilesAggregation.
func NewPercentilesAggregation() *PercentilesAggregation {
	r := &PercentilesAggregation{}

	return r
}
