export default {
  NodeParameter :{
    ClearAlarmNode: {
      template: `<div :currentField="currentField">
        <el-form-item label="Alarm details builder" prop="">
        　<span>function Details(msg, metadata, msgType) { </span><br>
          <div class="codeditor" style="border: 1px solid #8888">
            <codemirror ref="myCm" v-model="currentField.item.contents" :options="currentField.cmOptions" class="code"></codemirror>
          </div>
          <span>}</span><br>
          <el-button type="primary">TEST DETAILS FUNCTION</el-button>
        </el-form-item>
        <el-form-item label="Alarm type" prop="alarmType">
          <el-input v-model="currentField.alarmType"></el-input>
          <span style="font-size:12px;">Type pattern, use \${metaKeyName} to substitute variables from metadata</span>
        </el-form-item>
      </div> `,
      currentField: {
        alarmType: 'General Alarm',
        item: {
          contents: 'var details = {};if (metadata.prevAlarmDetails) {details = JSON.parse(metadata.prevAlarmDetails);return details;}'
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
            // 代码提示功能  根据Mode设置
          }
        }
      },
      linkType:'typeCleard'  
    }
  },
  linkLabelOptions: {
    typeCleard: [{
      value: 'Cleard',
      label: 'Cleard'
    }, {
      value: 'False',
      label: 'False'
    }, {
      value: 'Failure',
      label: 'Failure'
    }]
  }
}

