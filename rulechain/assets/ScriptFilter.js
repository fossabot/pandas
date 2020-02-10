export default {
  template: `<div :currentField="currentField">
    <el-form-item label="Filter" prop="">
      <span>function Filter(msg, metadata, msgType) { </span><br>
      <div class="codeditor" style="border: 1px solid #8888">
        <codemirror ref="myCm" v-model="currentField.item.contents" :options="currentField.cmOptions" class="code"></codemirror>
      </div>
      <span>}</span><br>
        <el-button type="primary">TEST FILTER FUNCTION</el-button>
    </el-form-item>
  </div> `,
  currentField: {
    item: {
      contents: "return msg.temperature > 20;"
    },
    cmOptions: {
      value: "",
      mode: "text/html",
      them: "foo bar",
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
  linkType:[{
    value: 'True',
    label: 'True'
  }, {
    value: 'False',
    label: 'False'
  }, {
    value: 'Failure',
    label: 'Failure'
  }]
}

