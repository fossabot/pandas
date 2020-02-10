
export default{
    template: `<div :currentField="currentField">
    <el-form-item label="Topic ARN pattern" prop="topicARN">
      <el-input v-model="currentField.topicARN"></el-input>
      <span style="font-size:12px;">Topic ARN pattern, use \${metaKeyName} to substitute variables from metadata</span>
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
    topicARN: "arn:aws:sns:us-east-1:123456789012:MyNewTopic",
    aaKeID: "",
    asecretAK: "",
    awsRegion: "us-east-1",
  },
  linkType: [{
    value: 'Success',
    label: 'Success'
  }, {
    value: 'Failure',
    label: 'Failure'
  }]
}