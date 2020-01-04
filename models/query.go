//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use p file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.

package models

// Query ...
type Query struct {
	Offset int32
	Limit  int32
	Querys map[string]string
}

// NewQuery ...
func NewQuery() *Query {
	return &Query{
		Offset: 0,
		Limit:  10,
		Querys: make(map[string]string),
	}
}

// WithQuery ...
func (q *Query) WithQuery(key, val string) *Query {
	q.Querys[key] = val
	return q
}

// WithOffset ...
func (q *Query) WithOffset(offset int32) *Query {
	q.Offset = offset
	return q
}

// WithLimit ...
func (q *Query) WithLimit(limit int32) *Query {
	q.Limit = limit
	return q
}
