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

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/valuetype"
)

// WeightedAverageAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/00fd9ffbc085e011cce9deb05bab4feaaa6b4115/specification/_types/aggregations/metric.ts#L432-L446
type WeightedAverageAggregation struct {
	// Format A numeric response formatter.
	Format *string  `json:"format,omitempty"`
	Meta   Metadata `json:"meta,omitempty"`
	Name   *string  `json:"name,omitempty"`
	// Value Configuration for the field that provides the values.
	Value     *WeightedAverageValue `json:"value,omitempty"`
	ValueType *valuetype.ValueType  `json:"value_type,omitempty"`
	// Weight Configuration for the field or script that provides the weights.
	Weight *WeightedAverageValue `json:"weight,omitempty"`
}

func (s *WeightedAverageAggregation) UnmarshalJSON(data []byte) error {

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

		case "format":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Format", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Format = &o

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return fmt.Errorf("%s | %w", "Meta", err)
			}

		case "name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name = &o

		case "value":
			if err := dec.Decode(&s.Value); err != nil {
				return fmt.Errorf("%s | %w", "Value", err)
			}

		case "value_type":
			if err := dec.Decode(&s.ValueType); err != nil {
				return fmt.Errorf("%s | %w", "ValueType", err)
			}

		case "weight":
			if err := dec.Decode(&s.Weight); err != nil {
				return fmt.Errorf("%s | %w", "Weight", err)
			}

		}
	}
	return nil
}

// NewWeightedAverageAggregation returns a WeightedAverageAggregation.
func NewWeightedAverageAggregation() *WeightedAverageAggregation {
	r := &WeightedAverageAggregation{}

	return r
}
