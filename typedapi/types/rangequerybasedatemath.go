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
// https://github.com/elastic/elasticsearch-specification/tree/cdb84fa39f1401846dab6e1c76781fb3090527ed

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/rangerelation"
)

// RangeQueryBaseDateMath type.
//
// https://github.com/elastic/elasticsearch-specification/blob/cdb84fa39f1401846dab6e1c76781fb3090527ed/specification/_types/query_dsl/term.ts#L109-L133
type RangeQueryBaseDateMath struct {
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost *float32 `json:"boost,omitempty"`
	From  *string  `json:"from,omitempty"`
	// Gt Greater than.
	Gt *string `json:"gt,omitempty"`
	// Gte Greater than or equal to.
	Gte *string `json:"gte,omitempty"`
	// Lt Less than.
	Lt *string `json:"lt,omitempty"`
	// Lte Less than or equal to.
	Lte        *string `json:"lte,omitempty"`
	QueryName_ *string `json:"_name,omitempty"`
	// Relation Indicates how the range query matches values for `range` fields.
	Relation *rangerelation.RangeRelation `json:"relation,omitempty"`
	To       *string                      `json:"to,omitempty"`
}

func (s *RangeQueryBaseDateMath) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "boost":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "Boost", err)
				}
				f := float32(value)
				s.Boost = &f
			case float64:
				f := float32(v)
				s.Boost = &f
			}

		case "from":
			if err := dec.Decode(&s.From); err != nil {
				return fmt.Errorf("%s | %w", "From", err)
			}

		case "gt":
			if err := dec.Decode(&s.Gt); err != nil {
				return fmt.Errorf("%s | %w", "Gt", err)
			}

		case "gte":
			if err := dec.Decode(&s.Gte); err != nil {
				return fmt.Errorf("%s | %w", "Gte", err)
			}

		case "lt":
			if err := dec.Decode(&s.Lt); err != nil {
				return fmt.Errorf("%s | %w", "Lt", err)
			}

		case "lte":
			if err := dec.Decode(&s.Lte); err != nil {
				return fmt.Errorf("%s | %w", "Lte", err)
			}

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "QueryName_", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueryName_ = &o

		case "relation":
			if err := dec.Decode(&s.Relation); err != nil {
				return fmt.Errorf("%s | %w", "Relation", err)
			}

		case "to":
			if err := dec.Decode(&s.To); err != nil {
				return fmt.Errorf("%s | %w", "To", err)
			}

		}
	}
	return nil
}

// NewRangeQueryBaseDateMath returns a RangeQueryBaseDateMath.
func NewRangeQueryBaseDateMath() *RangeQueryBaseDateMath {
	r := &RangeQueryBaseDateMath{}

	return r
}
