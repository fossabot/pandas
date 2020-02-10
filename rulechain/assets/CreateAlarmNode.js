export default {
  NodeParameter :{
    CreateAlarmNode: {
      template: `<div :currentField="currentField">
        <el-form-item label="Alarm details builder" prop="">
          　<span>function Details(msg, metadata, msgType) { </span><br>
            <div class="codeditor" style="border: 1px solid #8888">
              <codemirror ref="myCm" v-model="currentField.item.contents" :options="currentField.cmOptions" class="code"></codemirror>
            </div>
            <span>}</span><br>
            <el-button type="primary">TEST DETAILS FUNCTION</el-button>
        </el-form-item>
        <el-form-item label="" prop="">
          <el-checkbox style="zoom:120%;" label=" Use message alarm data" name="type" v-model="currentField.delivery"></el-checkbox><br>
        </el-form-item>
        <el-form-item label="" prop="alarmType" v-if="currentField.delivery === false">
          <el-row>
            <el-col :span="12">
              <span>Alarm type *</span>
            </el-col>
            <el-col :span="12">
              <span>Alarm severity *</span>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-input v-model="currentField.alarmType" style="width:90%;"></el-input>
            </el-col>
            <el-col :span="12">
              <el-select v-model="currentField.alarmSeverity" placeholder="请选择" style="width:90%;">
                  <el-option
                  v-for="item in currentField.allSeverity"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                  </el-option>
              </el-select>
            </el-col>
          </el-row>
        </el-form-item>
        <el-form-item label="" prop="" v-if="currentField.delivery === false">
          <el-checkbox style="zoom:120%;" label=" Propagate" name="type" v-model="currentField.propagate"></el-checkbox><br>
        </el-form-item>
      </div> `,
      currentField: {
        delivery: false,
        alarmType: 'General Alarm',
        alarmSeverity: '',
        propagate: false,
        allSeverity: [
          {
            label: '危险',
            value: '危险'
          }, {
            label: '重要',
            value: '重要'
          }, {
            label: '次要',
            value: '次要'
          }, {
            label: '警告',
            value: '警告'
          }, {
            label: '不确定',
            value: '不确定'
          }
        ],
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

