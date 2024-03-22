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

// GarbageCollector type.
//
// https://github.com/elastic/elasticsearch-specification/blob/00fd9ffbc085e011cce9deb05bab4feaaa6b4115/specification/nodes/_types/Stats.ts#L923-L928
type GarbageCollector struct {
	// Collectors Contains statistics about JVM garbage collectors for the node.
	Collectors map[string]GarbageCollectorTotal `json:"collectors,omitempty"`
}

// NewGarbageCollector returns a GarbageCollector.
func NewGarbageCollector() *GarbageCollector {
	r := &GarbageCollector{
		Collectors: make(map[string]GarbageCollectorTotal, 0),
	}

	return r
}
