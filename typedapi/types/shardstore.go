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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/shardstoreallocation"
)

// ShardStore type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ca0cc05d3ae3fa06c2cd7be91905b656a474334/specification/indices/shard_stores/types.ts#L29-L38
type ShardStore struct {
	Allocation       shardstoreallocation.ShardStoreAllocation `json:"allocation"`
	AllocationId     string                                    `json:"allocation_id"`
	Attributes       map[string]interface{}                    `json:"attributes"`
	Id               string                                    `json:"id"`
	LegacyVersion    int64                                     `json:"legacy_version"`
	Name             string                                    `json:"name"`
	StoreException   ShardStoreException                       `json:"store_exception"`
	TransportAddress string                                    `json:"transport_address"`
}

// NewShardStore returns a ShardStore.
func NewShardStore() *ShardStore {
	r := &ShardStore{
		Attributes: make(map[string]interface{}, 0),
	}

	return r
}