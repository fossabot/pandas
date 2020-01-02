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
package rabbitmq

type rabbitmqDataSource struct{}

func (g *rabbitmqDataSource) Name() string                               { return "rabbitmq" }
func (g *rabbitmqDataSource) Start(configs map[string]interface{}) error { return nil }
func (g *rabbitmqDataSource) GracefulStop() error                        { return nil }
