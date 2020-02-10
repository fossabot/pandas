export default {
  NodeParameter :{
    SwitchNode: {
      template: `<div :currentField="currentField">
        <el-form-item label="Switch" prop="">
          <span>function Switch(msg, metadata, msgType) { </span><br>
            <div class="codeditor" style="border: 1px solid #8888">
              <codemirror ref="myCm" v-model="currentField.item.contents" :options="currentField.cmOptions" class="code"></codemirror>
            </div>
          <span>}</span><br>
          <el-button type="primary">TEST SWITCH FUNCTION</el-button>
        </el-form-item>
      </div> `,
      currentField: {
        item: {
          contents: 'function nextRelation(metadata, msg) {return ["one","nine"];} if(msgType === "POST_TELEMETRY_REQUEST") {return ["two"];}return nextRelation(metadata, msg);'
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
            // lineWrapping: true,
          lineNumbers: true,
          firstLineNumber: 1,
          cursorHeight: 1,
          hintOptions: {
                  // 代码提示功能  根据Mode设置
          }
        }
      },
      linkType:'typeFailure'
    }
  },
  linkLabelOptions: {
    typeFailure: [{
      value: 'Failure',
      label: 'Failure'
    }]
  }
}

