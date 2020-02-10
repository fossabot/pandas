export default {
  NodeParameter :{
    KafkaNode: {
      template: `<div :currentField="currentField">
        <el-form-item label="Topic pattern" prop="topicPat">
          <el-input v-model="currentField.topicPat"></el-input>
        </el-form-item>
        <el-form-item label="Bootstrap servers" prop="serverIp">
          <el-input v-model="currentField.serverIp"></el-input>
        </el-form-item>
        <el-form-item label="Automatically retry times if fails" prop="">
          <el-input-number v-model="currentField.retryTime" controls-position="right" :min="0" style="width:100%"></el-input-number>
        </el-form-item>
        <el-form-item label="Produces batch size in bytes" prop="">
          <el-input-number v-model="currentField.batchSize" controls-position="right" :min="0" style="width:100%"></el-input-number>
        </el-form-item>
        <el-form-item label="Time to buffer locally (ms)" prop="">
          <el-input-number v-model="currentField.bufferTime" controls-position="right" :min="0" style="width:100%"></el-input-number>
        </el-form-item>
        <el-form-item label="Client buffer max size in bytes" prop="">
          <el-input-number v-model="currentField.bufferMaxsize" controls-position="right" :min="0" style="width:100%"></el-input-number>
        </el-form-item>
        <el-form-item label="Number of acknowledgments" prop="">
          <el-select v-model="currentField.acknowNum" placeholder="请选择">
              <el-option
              v-for="item in currentField.allnum"
              :key="item.value"
              :label="item.label"
              :value="item.value">
              </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="Key serializer" prop="keySeria">
          <el-input v-model="currentField.keySeria"></el-input>
        </el-form-item>
        <el-form-item label="Value serializer" prop="valueSeria">
          <el-input v-model="currentField.valueSeria"></el-input>
        </el-form-item>
        <el-form-item label="Other properties" prop="">
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
        topicPat: 'my-topic',
        serverIp: 'localhost:9092',
        retryTime: 0,
        batchSize: 16384,
        bufferTime: 0,
        bufferMaxsize: 33554432,
        acknowNum: '',
        keySeria: 'org.apache.kafka.common.serialization.StringSerializer',
        valueSeria: 'org.apache.kafka.common.serialization.StringSerializer',
        allnum: [
          {
            label: 'all',
            value: 'all'
          }, {
            label: '-1',
            value: '-1'
          }, {
            label: '0',
            value: '0'
          }, {
            label: '1',
            value: '1'
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

