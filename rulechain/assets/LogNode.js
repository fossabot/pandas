export default {
  NodeParameter :{
    LogNode: {
      template: `<div :currentField="currentField">
        <el-form-item label="To string" prop="">
          　<span>function Tostring(msg, metadata, msgType) { </span><br>
            <div class="codeditor" style="border: 1px solid #8888">
              <codemirror ref="myCm" v-model="currentField.item.contents" :options="currentField.cmOptions" class="code"></codemirror>
            </div>
            <span>}</span><br>
            <el-button type="primary">TEST To STRING FUNCTION</el-button>
        </el-form-item>
      </div> `,
      currentField: {
        item: {
          contents: 'return "Incoming message:\n" + JSON.stringify(msg) + "\nIncoming metadata:\n" + JSON.stringify(metadata);'
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

