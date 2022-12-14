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


// Package dfraftereffect
package dfraftereffect

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/d63a0e35ee85d84c83d0612ff3c0641a7a1e7e2e/specification/_types/Similarity.ts#L26-L30
type DFRAfterEffect struct {
	Name string
}

var (
	No = DFRAfterEffect{"no"}

	B = DFRAfterEffect{"b"}

	L = DFRAfterEffect{"l"}
)

func (d DFRAfterEffect) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *DFRAfterEffect) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "no":
		*d = No
	case "b":
		*d = B
	case "l":
		*d = L
	default:
		*d = DFRAfterEffect{string(text)}
	}

	return nil
}

func (d DFRAfterEffect) String() string {
	return d.Name
}
