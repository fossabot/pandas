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
package nodes

// init register all node's factory
func init() {
	// Input Node
	RegisterFactory(inputNodeFactory{})

	// Filter Nodes
	RegisterFactory(checkRelationFilterNodeFactory{})
	RegisterFactory(messageTypeFilterNodeFactory{})
	RegisterFactory(messageTypeSwitchNodeFactory{})
	RegisterFactory(originatorFilterNodeFactory{})
	RegisterFactory(originatorTypeSwitchNodeFactory{})
	RegisterFactory(scriptFilterNodeFactory{})
	RegisterFactory(switchFilterNodeFactory{})

	// Action Nodes
	RegisterFactory(assignToCustomerFactory{})
	RegisterFactory(createAlarmNodeFactory{})
	RegisterFactory(clearAlarmNodeFactory{})
	RegisterFactory(createRelationNodeFactory{})
	RegisterFactory(delayNodeFactory{})
	RegisterFactory(deleteRelationNodeFactory{})
	RegisterFactory(messageGeneratorNodeFactory{})
	RegisterFactory(logNodeFactory{})
	RegisterFactory(rpcCallRequestNodeFactory{})
	RegisterFactory(rpcCallReplyNodeFactory{})
	RegisterFactory(saveAttributesNodeFactory{})
	RegisterFactory(saveTimeSeriesNodeFactory{})
	RegisterFactory(unassignFromCustomerNodeFactory{})

	// Enrichment Nodes
	RegisterFactory(enrichmentCustomerAttrNodeFactory{})
	RegisterFactory(enrichmentDeviceAttrNodeFactory{})
	RegisterFactory(enrichmentOriginatorAttrNodeFactory{})
	RegisterFactory(enrichmentOriginatorFieldsNodeFactory{})
	RegisterFactory(enrichmentOriginatorTelemetryNodeFactory{})

	// Transform Nodes
	RegisterFactory(transformChangeOriginatorNodeFactory{})
	RegisterFactory(transformScriptNodeFactory{})
	RegisterFactory(transformToEmailNodeFactory{})

	// External Nodes
	RegisterFactory(externalMqttNodeFactory{})
	RegisterFactory(externalRestapiNodeFactory{})
}
