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
package runtime

type RelationFilter struct {
	Type        string `json:"type" yaml:"type"`
	EntityTypes string `json:"entity_types" yaml:"entity_types"`
}

type RelationQuery interface {
	QueryEntities(direction string, maxRelationLevel int, filters []RelationFilter) []string
}

func NewRelationQuery() RelationQuery {
	return nil
}
