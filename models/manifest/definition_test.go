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
package manifest

import "testing"

var definitionTestData1 = `
name: electrical-machinery
description: sample
kind: deviceModel
domain: electrical
version: 0.1
network: "mqtt"
logical: false
compound: false
endpoints:
  - path: status
    dataModel: statusNotify
    permission: readonly
  - path: control
    dataModel: control
    permission: write
dataModels:
  - name: statusNotify
    fields:
      - key: direction
        type: int 
        defaultValue: 0
      - key: speed
        type: int
        defaultValue: 0
  - name: control
    fields:
      - key: speed
        type: int
        defaultValue: 0
      - key: direction
        type: int
        defaultValue: 0
childDeviceModels:
    - name: "xxyyxxzz"
    - name: "zzzzzzzz"
`

var definitionTestData2 = `
name: electrical-machinery
description: sample
kind: deviceModel
domain: electrical
version: 0.1
network: "mqtt"
logical: false
compound: false
endpoints:
  - path: status
    dataModel: statusNotifyXXX
    permission: readonly
  - path: control
    dataModel: control
    permission: write
dataModels:
  - name: statusNotify
    fields:
      - key: direction
        type: int 
        defaultValue: 0
      - key: speed
        type: int
        defaultValue: 0
  - name: control
    fields:
      - key: speed
        type: int
        defaultValue: 0
      - key: direction
        type: int
        defaultValue: 0
childDeviceModels:
    - name: "xxyyxxzz"
    - name: "zzzzzzzz"
`

func TestParseModelDefinition(t *testing.T) {
}
