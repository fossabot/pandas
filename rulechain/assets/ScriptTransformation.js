export default {
  NodeParameter :{
    ScriptTransformation: {
      template: `<div :currentField="currentField">
        <el-form-item label="Transform" prop="">
          <span>function Transform(msg, metadata, msgType) { </span><br>
            <div class="codeditor" style="border: 1px solid #8888">
              <codemirror ref="myCm" v-model="currentField.item.contents" :options="currentField.cmOptions" class="code"></codemirror>
            </div>
          <span>}</span><br>
          <el-button type="primary">TEST TRANSFORMER FUNCTION</el-button>
        </el-form-item>
      </div> `,
      currentField: {
        item: {
          contents: 'return {msg: msg, metadata: metadata, msgType: msgType};'
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

