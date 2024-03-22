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
// https://github.com/elastic/elasticsearch-specification/tree/00fd9ffbc085e011cce9deb05bab4feaaa6b4115

package types

// TermVector type.
//
// https://github.com/elastic/elasticsearch-specification/blob/00fd9ffbc085e011cce9deb05bab4feaaa6b4115/specification/_global/termvectors/types.ts#L23-L26
type TermVector struct {
	FieldStatistics FieldStatistics `json:"field_statistics"`
	Terms           map[string]Term `json:"terms"`
}

// NewTermVector returns a TermVector.
func NewTermVector() *TermVector {
	r := &TermVector{
		Terms: make(map[string]Term, 0),
	}

	return r
}
