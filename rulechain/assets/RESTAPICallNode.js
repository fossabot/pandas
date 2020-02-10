export default {
  NodeParameter :{
    RESTAPICallNode: {
      template: `<div :currentField="currentField">
        <el-form-item label="Endpoint URL pattern" prop="endpointUrl">
          <el-input v-model="currentField.endpointUrl"></el-input>
          <span style="font-size:12px;">HTTP URL address pattern, use \${metaKeyName} to substitute variables from metadata</span>
        </el-form-item>
        <el-form-item label="Request method" prop="">
          <el-select v-model="currentField.requestMethod" placeholder="请选择" style="width:100%;">
              <el-option
              v-for="item in currentField.allmethods"
              :key="item.value"
              :label="item.label"
              :value="item.value">
              </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="" prop="">
          <el-checkbox style="zoom:120%;" label=" Use simple client HTTP factory" name="type" v-model="currentField.uschf"></el-checkbox>
        </el-form-item>
        <el-form-item label="Headers" prop="">
          <span style="font-size:12px;">Use \${metaKeyName} in header/value fields to substitute variables from metadata</span>
          <el-row>
            <el-col :span="11">
              <span style="color:#606266;">Header</span>
            </el-col>
            <el-col :span="11">
              <span style="color:#606266;">Value</span>
            </el-col>
          </el-row>
          <div v-for="(param, index) in currentField.Parameters" :key="index" style="margin-bottom:10px;">
          <el-row>
            <el-col :span="10">
              <el-input v-model="param.parameter" placeholder="Header"></el-input>
            </el-col>
            <el-col :span="10" style="margin-left:30px;">
              <el-input v-model="param.value" placeholder="Value"></el-input>
            </el-col>
            <el-button v-if="index === 0" icon="el-icon-plus" size="medium" style="padding: 10px 12px;margin-left:10px;background-color:#f5f7fa;"/>
            <el-button v-if="index !== 0" icon="el-icon-minus" size="medium" style="padding: 10px 12px;margin-left:10px;background-color:#f5f7fa"/>
          </el-row>
          </div>
        </el-form-item>
      </div> `,
      currentField: {
        endpointUrl: 'http://localhost/api',
        requestMethod: 'POST',
        uschf: false,
        Parameters: [
          { parameter: '', value: '' }
        ],
        allmethods: [
          {
            label: 'GET',
            value: 'GET'
          }, {
            label: 'POST',
            value: 'POST'
          }, {
            label: 'PUT',
            value: 'PUT'
          }, {
            label: 'DELETE',
            value: 'DELETE'
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

