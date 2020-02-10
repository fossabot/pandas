export default {
  NodeParameter :{
    MQTTNode: {
      template: `<div :currentField="currentField">
        <el-form-item label="Topic pattern" prop="topicPatt">
          <el-input v-model="currentField.topicPatt"></el-input>
          <span style="font-size:12px;">MQTT topic pattern, use \${metaKeyName} to substitute variables from metadata</span>
        </el-form-item>
        <el-form-item label="" prop="hostport">
          <el-row>
            <el-col :span="7">
              <span style="color:#606266;font-weight: bold;">Host *</span>
            </el-col>
            <el-col :span="7">
              <span style="color:#606266;font-weight: bold;margin-left:10px;">Port *</span>
            </el-col>
            <el-col :span="7">
              <span style="color:#606266;font-weight: bold;">Connection timeout(sec) *</span>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="7">
              <el-input v-model="currentField.host" style=width:90%;></el-input>
            </el-col>
            <el-col :span="7">
              <el-input-number v-model="currentField.port" controls-position="right" :min="1" :max="65535" style="width:100%;margin-left:1px;"></el-input-number>
            </el-col>
            <el-col :span="7">
              <el-input-number v-model="currentField.timeout" controls-position="right" :min="1" :max="200" style="width:100%;margin-left:15px;"></el-input-number>
            </el-col>
          </el-row>
        </el-form-item>
        <el-form-item label="Client ID" prop="">
          <el-input v-model="currentField.clientId"></el-input>
        </el-form-item>
        <el-form-item label="" prop="">
          <el-checkbox style="zoom:120%;" label=" Clean session" name="type" v-model="currentField.cleanSession"></el-checkbox><br>
          <el-checkbox style="zoom:120%;" label=" Enable SSL" name="type" v-model="currentField.enableSSL"></el-checkbox>
        </el-form-item>
        <el-card>
          <div slot="header" class="clearfix">
            <span style="font-size:16px;">Credentials</span>
            <span style="margin-left: 100px;">{{$t('nodeOther.anonymous')}}</span>
            <span style="margin-left: 100px;">{{$t('nodeOther.anonymous')}}</span>
            <span style="margin-left: 100px;">{{$t('nodeOther.anonymous')}}</span>
          </div>
          <div>
          </div>
        </el-card>
      </div> `,
      currentField: {
        topicPatt: 'my-topic',
        host: 'localhost',
        port: 1883,
        timeout: 10,
        clientId: '',
        cleanSession: true,
        enableSSL: false,
        activeName: '1',
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

