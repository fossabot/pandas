export default {
  NodeParameter :{
    AWSSQSNode: {
      template: `<div :currentField="currentField">
        <el-form-item label="Queue type" prop="">
          <el-select v-model="currentField.queueType" placeholder="请选择" style="width:100%;">
            <el-option label="Standard" value="Standard"></el-option>
            <el-option label="FIFO" value="FIFO"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="Queue URL pattern" prop="queueUrl">
          <el-input v-model="currentField.queueUrl"></el-input>
          <span style="font-size:12px;">Queue URL pattern, use \${metaKeyName} to substitute variables from metadata</span>
        </el-form-item>
        <el-form-item label="Delay (seconds)" prop="" v-if="currentField.queueType === 'Standard'">
          <el-input-number v-model="currentField.delayTime" controls-position="right" :min="0"></el-input-number>
        </el-form-item>
        <el-form-item label="Message attributes" prop="">
          <span style="font-size:12px;">Use \${metaKeyName} in name/value fields to substitute variables from metadata</span>
          <el-row>
            <el-col :span="11">
              <span style="color:#606266;font-weight:600;">Name</span>
            </el-col>
            <el-col :span="11">
              <span style="color:#606266;font-weight:600;">Value</span>
            </el-col>
          </el-row>
          <div v-for="(param, index) in currentField.Parameters" :key="index" style="margin-bottom:10px;">
          <el-row>
            <el-col :span="10">
              <el-input v-model="param.parameter" placeholder="Name"></el-input>
            </el-col>
            <el-col :span="10" style="margin-left:30px;">
              <el-input v-model="param.value" placeholder="Value"></el-input>
            </el-col>
            <el-button v-if="index === 0" icon="el-icon-plus" size="medium" style="padding: 10px 12px;margin-left:10px;background-color:#f5f7fa;"/>
            <el-button v-if="index !== 0" icon="el-icon-minus" size="medium" style="padding: 10px 12px;margin-left:10px;background-color:#f5f7fa"/>
          </el-row>
          </div>
        </el-form-item>
        <el-form-item label="AWS Access Key ID" prop="aaKeID">
          <el-input v-model="currentField.aaKeID"></el-input>
        </el-form-item>
        <el-form-item label="AWS Secret Access Key" prop="asecretAK">
          <el-input v-model="currentField.asecretAK"></el-input>
        </el-form-item>
        <el-form-item label="AWS Region" prop="awsRegion">
          <el-input v-model="currentField.awsRegion"></el-input>
        </el-form-item>
      </div> `,
      currentField: {
        queueType: 'Standard',
        queueUrl: 'https://sqs.us-east-1.amazonaws.com/123456789012/my-queue-name',
        delayTime: 0,
        aaKeID: '',
        asecretAK: '',
        awsRegion: 'us-east-1',
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

