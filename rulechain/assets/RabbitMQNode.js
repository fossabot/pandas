export default {
  NodeParameter :{
    RabbitMQNode: {
      template: `<div :currentField="currentField">
        <el-form-item label="Exchange name pattern" prop="">
          <el-input v-model="currentField.exchangeName"></el-input>
        </el-form-item>
        <el-form-item label="Routing key pattern" prop="">
          <el-input v-model="currentField.routeKey"></el-input>
        </el-form-item>
        <el-form-item label="Message properties" prop="">
          <el-select v-model="currentField.messPropert" placeholder="请选择" style="width:100%;">
              <el-option
              v-for="item in currentField.allmesspropert"
              :key="item.value"
              :label="item.label"
              :value="item.value">
              </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="" prop="">
          <el-row>
            <el-col :span="11">
              <span>Host *</span>
            </el-col>
            <el-col :span="11">
              <span style="margin-left:10px;">Port *</span>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="11">
              <el-input v-model="currentField.host"></el-input>
            </el-col>
            <el-col :span="11">
              <el-input-number v-model="currentField.port" controls-position="right" :min="0" :max="65535" style="width:100%;margin-left:10px;"></el-input-number>
            </el-col>
          </el-row>
        </el-form-item>
        <el-form-item label="Virtual host" prop="">
          <el-input v-model="currentField.virtualHost"></el-input>
        </el-form-item>
        <el-form-item label="Username" prop="">
          <el-input name="username" v-model="currentField.username"></el-input>
        </el-form-item>
        <el-form-item label="Password" prop="">
          <el-input name="password" v-model="currentField.password"></el-input>
        </el-form-item>
        <el-form-item label="" prop="">
          <el-checkbox style="zoom:120%;" label=" Automatic recovery" name="type" v-model="currentField.autoRecovery"></el-checkbox>
        </el-form-item>
        <el-form-item label="Connection timeout (ms)" prop="">
          <el-input-number v-model="currentField.ConnectTimeout" controls-position="right" :min="0" style="width:90%;"></el-input-number>
        </el-form-item>
        <el-form-item label="Handshake timeout (ms)" prop="">
          <el-input-number v-model="currentField.HandshakeTimemout" controls-position="right" :min="0" style="width:90%;"></el-input-number>
        </el-form-item>
        <el-form-item label="Client properties" prop="">
        <el-row>
            <el-col :span="11">
              <span>Key</span>
            </el-col>
            <el-col :span="11">
              <span>Value</span>
            </el-col>
          </el-row>
          <div v-for="(param, index) in currentField.Parameters" :key="index" style="margin-bottom:10px;">
          <el-row>
            <el-col :span="10">
              <el-input v-model="param.parameter" placeholder="Key"></el-input>
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
        exchangeName: '',
        routeKey: '',
        messPropert: '',
        host: 'localhost',
        port: '5672',
        virtualHost: '/',
        username: '',
        password: '',
        autoRecovery: false,
        ConnectTimeout: 60000,
        HandshakeTimemout: 10000,
        allmesspropert: [
          {
            label: 'BASIC',
            value: 'BASIC'
          }, {
            label: 'TEXT_PLAIN',
            value: 'TEXT_PLAIN'
          }, {
            label: 'MINIMAL_BASIC',
            value: 'MINIMAL_BASIC'
          }, {
            label: 'MINIMAL_PERSISTENT_BASIC',
            value: 'MINIMAL_PERSISTENT_BASIC'
          }, {
            label: 'PERSISTENT_BASIC',
            value: 'PERSISTENT_BASIC'
          }, {
            label: 'PERSISTENT_TEXT_PLAIN',
            value: 'PERSISTENT_TEXT_PLAIN'
          }
        ],
        Parameters: [
          { parameter: '', value: '' }
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

