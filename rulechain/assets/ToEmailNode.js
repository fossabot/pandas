export default {
  NodeParameter :{
    ToEmailNode: {
      template: `<div :currentField="currentField">
        <el-form-item label="From Template" prop="fromTemp">
          <el-input type="textarea" :rows="2" placeholder="" v-model="currentField.fromTemp"></el-input><br>
          <span style="font-size:12px;">From address template, use \${metaKeyName} to substitute variables from metadata</span>
        </el-form-item>
        <el-form-item label="To Template" prop="toTemp">
          <el-input type="textarea" :rows="2" placeholder="" v-model="currentField.toTemp"></el-input><br>
          <span style="font-size:12px;">Comma separated address list, use \${metaKeyName} to substitute variables from metadata</span>
        </el-form-item>
        <el-form-item label="Cc Template" prop="ccTemp" >
          <el-input type="textarea" :rows="2" placeholder="" v-model="currentField.ccTemp"></el-input><br>
          <span style="font-size:12px;">Comma separated address list, use \${metaKeyName} to substitute variables from metadata</span>
        </el-form-item>
        <el-form-item label="Bcc Template" prop="bccTemp" >
          <el-input type="textarea" :rows="2" placeholder="" v-model="currentField.bccTemp"></el-input><br>
          <span style="font-size:12px;">Comma separated address list, use \${metaKeyName} to substitute variables from metadata</span>
        </el-form-item>
        <el-form-item label="Subject Template" prop="subjectTemp" >
          <el-input type="textarea" :rows="2" placeholder="" v-model="currentField.subjectTemp"></el-input><br>
          <span style="font-size:12px;">Mail subject template, use \${metaKeyName} to substitute variables from metadata</span>
        </el-form-item>
        <el-form-item label="Body Template" prop="bodyTemp" >
          <el-input type="textarea" :rows="5" placeholder="" v-model="currentField.bodyTemp"></el-input><br>
          <span style="font-size:12px;">Mail body template, use \${metaKeyName} to substitute variables from metadata</span>
        </el-form-item>
      </div> `, 
      currentField: {
        fromTemp: 'info@testmail.org',
        toTemp: '\${userEmail}',
        ccTemp: '',
        bccTemp: '',
        subjectTemp: 'Device \${deviceType} temperature high',
        bodyTemp: 'Device \${deviceName} has high temperature \${temp}'
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

