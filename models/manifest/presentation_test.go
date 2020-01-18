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

func TestParseModelPresentation(t *testing.T) {
	_, err := ParseModelPresentation([]byte(presentationTestData), "json")
	if err != nil {
		t.Error(err)
	}
}

var presentationTestData = `{
	"nodes": [{
		"name": "LoadBalancer_1",
		"type": "LoadBalancer",
		"id": "1hokg0mgnmww",
		"x": 840,
		"y": 180,
		"icon": "hello",
		"width": 115,
		"height": 60,
		"initW": 115,
		"initH": 60,
		"classType": "T1",
		"isLeftConnectShow": true,
		"isRightConnectShow": false,
		"containNodes": [],
		"attrs": [],
		"isSelect": false
	}, {
		"name": "Router_1",
		"type": "Router",
		"id": "dbfp8krc8ow",
		"x": 680,
		"y": 220,
		"icon": "hello2",
		"width": 50,
		"height": 50,
		"initW": 50,
		"initH": 50,
		"classType": "T2",
		"isLeftConnectShow": true,
		"isRightConnectShow": true,
		"containNodes": [],
		"attrs": [],
		"isSelect": false
	}, {
		"name": "Service_1",
		"type": "Service",
		"id": "1orsbg8m30zk",
		"x": 60,
		"y": 40,
		"icon": "hello3", 
		"width": 552,
		"height": 160,
		"initW": 140,
		"initH": 80,
		"classType": "T1",
		"isLeftConnectShow": false,
		"isRightConnectShow": true,
		"containNodes": ["jdturwsp9xc", "6isr7j0hiec0"],
		"attrs": [],
		"isSelect": false
	}, {
		"name": "Pod_1",
		"type": "Pod",
		"id": "jdturwsp9xc",
		"x": 82,
		"y": 70,
		"icon": "hello4", 
		"width": 346,
		"height": 120,
		"initW": 140,
		"initH": 80,
		"classType": "T1",
		"isLeftConnectShow": false,
		"isRightConnectShow": false,
		"containNodes": ["66", "88"],
		"attrs": [],
		"isSelect": false
	}, {
		"x": 266,
		"y": 100,
		"width": 140,
		"height": 80,
		"id": "88",
		"isLeftConnectShow": false,
		"isRightConnectShow": false,
		"name": "Container_c",
		"isSelect": false,
		"initW": 140,
		"initH": 80,
		"icon": "hello5", 
		"classType": "T1",
		"containNodes": [],
		"attrs": [],
		"type": "Container"
	}, {
		"name": "Pod_2",
		"type": "Pod",
		"id": "6isr7j0hiec0",
		"x": 450,
		"y": 70,
		"icon": "hello6", 
		"width": 140,
		"height": 80,
		"initW": 140,
		"initH": 80,
		"classType": "T1",
		"isLeftConnectShow": false,
		"isRightConnectShow": false,
		"containNodes": [],
		"attrs": [],
		"isSelect": false
	}, {
		"x": 104,
		"y": 100,
		"width": 140,
		"height": 80,
		"id": "66",
		"isLeftConnectShow": false,
		"isRightConnectShow": false,
		"name": "Container_a",
		"isSelect": false,
		"initW": 140,
		"initH": 80,
		"icon": "hello7", 
		"classType": "T1",
		"containNodes": [],
		"type": "Container",
		"attrs": [{
			"type": "input",
			"name": "portId",
			"value": "2222141",
			"placeholder": "请输入portId",
			"rules": [{
				"required": true,
				"message": "请输入活动名称",
				"trigger": "blur"
			}],
			"disabled": true
		}, {
			"type": "select",
			"name": "server",
			"value": "",
			"placeholder": "请选择服务器",
			"options": [{
				"label": "上海服务器",
				"value": "shagnhai"
			}, {
				"label": "北京服务器",
				"value": "beijing"
			}],
			"disabled": false
		}, {
			"type": "checkbox",
			"name": "数据库类型",
			"value": [],
			"options": [{
				"label": "SQL server"
			}, {
				"label": "Access"
			}, {
				"label": "mySQL"
			}, {
				"label": "Oracle"
			}],
			"disabled": false
		}, {
			"type": "textarea",
			"name": "数据库",
			"value": "",
			"rules": [],
			"disabled": false
		}, {
			"type": "radio",
			"name": "数据类型",
			"value": "",
			"options": [{
				"label": "sql"
			}, {
				"label": "oracle"
			}],
			"disabled": true
		}, {
			"type": "keyVal",
			"name": "service",
			"value": "dbms"
		}, {
			"type": "keyVal",
			"name": "sql",
			"value": "sq"
		}]
	}],
	"connectors": [{
		"type": "Contain",
		"sourceNode": {
			"id": "66",
			"width": 140,
			"height": 80,
			"x": 104,
			"y": 100
		},
		"targetNode": {
			"id": "jdturwsp9xc",
			"width": 346,
			"height": 120,
			"x": 82,
			"y": 70
		},
		"isSelect": false
	}, {
		"type": "Contain",
		"sourceNode": {
			"id": "jdturwsp9xc",
			"width": 346,
			"height": 120,
			"x": 82,
			"y": 70
		},
		"targetNode": {
			"id": "1orsbg8m30zk",
			"width": 552,
			"height": 160,
			"x": 60,
			"y": 40
		},
		"isSelect": false
	}, {
		"type": "Line",
		"targetNode": {
			"x": 840,
			"y": 180,
			"id": "1hokg0mgnmww",
			"width": 115,
			"height": 60
		},
		"sourceNode": {
			"x": 680,
			"y": 220,
			"id": "dbfp8krc8ow",
			"width": 50,
			"height": 50
		},
		"isSelect": false
	}, {
		"type": "Line",
		"targetNode": {
			"x": 680,
			"y": 220,
			"id": "dbfp8krc8ow",
			"width": 50,
			"height": 50
		},
		"sourceNode": {
			"x": 60,
			"y": 40,
			"id": "1orsbg8m30zk",
			"width": 552,
			"height": 160
		},
		"isSelect": false
	}, {
		"type": "Contain",
		"sourceNode": {
			"id": "6isr7j0hiec0",
			"width": 140,
			"height": 80,
			"x": 450,
			"y": 70
		},
		"targetNode": {
			"id": "1orsbg8m30zk",
			"width": 552,
			"height": 160,
			"x": 60,
			"y": 40
		},
		"isSelect": false
	}, {
		"type": "Contain",
		"sourceNode": {
			"id": "88",
			"width": 140,
			"height": 80,
			"x": 266,
			"y": 100
		},
		"targetNode": {
			"id": "jdturwsp9xc",
			"width": 346,
			"height": 120,
			"x": 82,
			"y": 70
		},
		"isSelect": false
	}]
}`
