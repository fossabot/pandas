export default {
  NodeParameter :{
    GeneratorNode: {
      template: `<div :currentField="currentField">
        <el-form-item label="Message count (0 - unlimited)" prop="messCount">
          <el-input-number v-model="currentField.messCount" controls-position="right" :min="0"></el-input-number>
        </el-form-item>
        <el-form-item label="Period in seconds" prop="period">
          <el-input-number v-model="currentField.period" controls-position="right" :min="1"></el-input-number>
        </el-form-item>
        <el-form-item label="Originator" prop="">
          <el-select v-model="currentField.type" placeholder="类型">
            <el-option v-for="item in currentField.allOritypes" :key="item.value" :value="item.value"/>
          </el-select>
          <el-select v-model="currentField.typecontain" placeholder="请选择">
            <el-option v-for="con in currentField.allOritypes" :key="con.value" :value="con.value"/>
          </el-select>
        </el-form-item>
        <el-form-item label="Generate" prop="">
          　<span>function Generate(prevMsg, prevMetadata, prevMsgType) { </span><br>
            <div class="codeditor" style="border: 1px solid #8888">
              <codemirror ref="myCm" v-model="currentField.item.contents" :options="currentField.cmOptions" class="code"></codemirror>
            </div>
            <span>}</span><br>
            <el-button type="primary">TEST GENERATOR FUNCTION</el-button>
        </el-form-item>
      </div> `,
      currentField: {
        messCount: 0,
        period: 1,
        type: '',
        typecontain: '',
        allOritypes: [
          {
            value: '设备',
            contains: [
              {
                label: 'device1',
                value: 'device1'
              }, {
                label: 'device2',
                value: 'device2'
              }
            ]
          }, {
            value: '资产',
            contains: [
              {
                label: '1元',
                value: '1元'
              }, {
                label: '2元',
                value: '2元'
              }
            ]
          }, {
            value: '租户',
            contains: []
          }, {
            value: '客户',
            contains: []
          }, {
            value: 'Entity View',
            contains: []
          }, {
            value: '仪表盘',
            contains: []
          }
        ],
        item: {
          contents: 'var msg = { temp: 42, humidity: 77 };var metadata = { data: 40 };var msgType = "POST_TELEMETRY_REQUEST";return { msg: msg, metadata: metadata, msgType: msgType };'
        },
        cmOptions: {
          value: '',
          mode: 'text/html', // application/json
          them: 'foo bar',
          indentUnit: 2,
          smartIndent: true,
          tabSize: 4,
          readOnly: false,
          showCursorWhenSelecting: true,
          lineNumbers: true,
          firstLineNumber: 1,
          cursorHeight: 1,
          hintOptions: {
          }
        }
      },
      linkType:'typeSuccess'
    }
  },
  linkLabelOptions: {
    typeSuccess: [{
      value: 'Success',
      label: 'Success'
    }, {
      value: 'Failure',
      label: 'Failure'
    }]
  }
}

