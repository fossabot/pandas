export default {
  NodeParameter :{
    SendEmailNode: {
      template: `<div :currentField="currentField">
        <el-form-item label="" prop="">
          <el-checkbox style="zoom:120%;" label=" Use system SMTP settings" name="type" v-model="currentField.useSysSMTP"></el-checkbox>
        </el-form-item>
        <el-form-item label="Protocol" prop="" v-if="currentField.useSysSMTP === false">
          <el-select v-model="currentField.protocol" placeholder="请选择" style="width:100%;">
              <el-option
              v-for="item in currentField.allprotocol"
              :key="item.value"
              :label="item.label"
              :value="item.value">
              </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="" prop="hostport" v-if="currentField.useSysSMTP === false">
          <el-row>
            <el-col :span="11">
              <span style="color:#606266;font-weight: bold;">SMTP host *</span>
            </el-col>
            <el-col :span="11">
              <span style="color:#606266;font-weight: bold;margin-left:10px;">SMTP port *</span>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="11">
              <el-input v-model="currentField.smtpHost"></el-input>
            </el-col>
            <el-col :span="11">
              <el-input-number v-model="currentField.smtpPort" controls-position="right" :min="1" :max="65535" style="width:100%;margin-left:10px;"></el-input-number>
            </el-col>
          </el-row>
        </el-form-item>
        <el-form-item label="Timeout (ms)" prop="timeout" v-if="currentField.useSysSMTP === false">
          <el-input-number v-model="currentField.timeout" controls-position="right" :min="0"></el-input-number>
        </el-form-item>
        <el-form-item label="" prop="" v-if="currentField.useSysSMTP === false">
          <el-checkbox label=" Enable TLS" name="type" v-model="currentField.enableTLS"></el-checkbox>
        </el-form-item>
        <el-form-item label="Username" prop="" v-if="currentField.useSysSMTP === false">
          <el-input v-model="currentField.username" placeholder="Enter Username"></el-input>
        </el-form-item>
        <el-form-item label="Password" prop="" v-if="currentField.useSysSMTP === false">
          <el-input v-model="currentField.password" placeholder="Enter Password"></el-input>
        </el-form-item>
      </div> `, 
      currentField: {
        useSysSMTP: true,
        protocol: '',
        smtpHost: 'localhost',
        smtpPort: 25,
        timeout: 10000,
        enableTLS: false,
        username: '',
        password: '',
        allprotocol: [
          {
            label: 'SMTP',
            value: 'SMTP'
          }, {
            label: 'SMTPS',
            value: 'SMTPS'
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

