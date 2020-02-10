export default {
  NodeParameter :{
    OriginatorTelemetry: {
      template: `<div :currentField="currentField">
        <el-form-item label="Latest timeseries" prop="timeseries">
          <el-input v-model="currentField.latesttime" placeholder="Latest timeseries"></el-input>
        </el-form-item>
        <el-form-item label="Fetch mode" prop="fetchmode">
          <el-select v-model="currentField.fetchmode" placeholder="请选择" style="width:100%;">
              <el-option
              v-for="item in currentField.allmode"
              :key="item.value"
              :label="item.label"
              :value="item.value">
              </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="" prop="">
          <el-checkbox style="zoom:120%;" label=" Use metadata interval patterns" name="type" v-model="currentField.delivery"></el-checkbox><br>
          <span style="font-size:12px;">If selected, rule node use start and end interval patterns from message metadata assuming that intervals are in the milliseconds</span>
        </el-form-item>
        <el-form-item label="" prop="" v-if="currentField.delivery === false">
          <el-row>
            <el-col :span="12">
              <span>Start Interval *</span>
            </el-col>
            <el-col :span="12">
              <span>Start Interval Time Unit *</span>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-input-number v-model="currentField.startlevel" controls-position="right" :min="1" style="width:90%"></el-input-number>
            </el-col>
            <el-col :span="12">
              <el-select v-model="currentField.startUnit" placeholder="请选择" style="width:90%">
                  <el-option
                  v-for="item in currentField.allUnit"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                  </el-option>
              </el-select>
            </el-col>
          </el-row>
        </el-form-item>
        <el-form-item label="" prop="" v-if="currentField.delivery === false">
          <el-row>
            <el-col :span="12">
              <span>End Interval *</span>
            </el-col>
            <el-col :span="12">
              <span>End Interval Time Unit *</span>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-input-number v-model="currentField.endlevel" controls-position="right" :min="1" style="width:90%"></el-input-number>
            </el-col>
            <el-col :span="12">
              <el-select v-model="currentField.endUnit" placeholder="请选择" style="width:90%">
                  <el-option
                  v-for="item in currentField.allUnit"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                  </el-option>
              </el-select>
            </el-col>
          </el-row>
        </el-form-item>
        <el-form-item label="Start interval pattern" prop="startpattern" v-if="currentField.delivery === true">
          <el-input v-model="currentField.startpattern"></el-input><br>
          <span style="font-size:12px;">Start interval pattern, use \${metaKeyName} to substitute variables from metadata</span>
        </el-form-item>
        <el-form-item label="End interval pattern" prop="endpattern" v-if="currentField.delivery === true">
          <el-input v-model="currentField.endpattern"></el-input><br>
          <span style="font-size:12px;">End interval pattern, use \${metaKeyName} to substitute variables from metadata</span>
        </el-form-item>
      </div> `,
      currentField: {
        latesttime: '',
        fetchmode: 'FIRST',
        delivery: false,
        startlevel: 1,
        startUnit: 'Milliseconds',
        endlevel: 1,
        endUnit: 'Milliseconds',
        startpattern: '',
        endpattern: '',
        allmode: [
          {
            label: 'FIRST',
            value: 'FIRST'
          }, {
            label: 'LAST',
            value: 'LAST'
          }, {
            label: 'ALL',
            value: 'ALL'
          }
        ],
        allUnit: [
          {
            label: 'Milliseconds',
            value: 'Milliseconds'
          }, {
            label: 'Seconds',
            value: 'Seconds'
          }, {
            label: 'Minutes',
            value: 'Minutes'
          }, {
            label: 'Hours',
            value: 'Hours'
          }, {
            label: 'Days',
            value: 'Days'
          }
        ]
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

