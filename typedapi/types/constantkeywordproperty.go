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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
)

// ConstantKeywordProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ca0cc05d3ae3fa06c2cd7be91905b656a474334/specification/_types/mapping/specialized.ts#L44-L47
type ConstantKeywordProperty struct {
	Dynamic     *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields      map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove *int                           `json:"ignore_above,omitempty"`
	// Meta Metadata about the field.
	Meta       map[string]string   `json:"meta,omitempty"`
	Properties map[string]Property `json:"properties,omitempty"`
	Type       string              `json:"type,omitempty"`
	Value      interface{}         `json:"value,omitempty"`
}

// NewConstantKeywordProperty returns a ConstantKeywordProperty.
func NewConstantKeywordProperty() *ConstantKeywordProperty {
	r := &ConstantKeywordProperty{
		Fields:     make(map[string]Property, 0),
		Meta:       make(map[string]string, 0),
		Properties: make(map[string]Property, 0),
	}

	r.Type = "constant_keyword"

	return r
}