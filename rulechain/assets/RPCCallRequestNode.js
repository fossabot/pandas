export default {
  NodeParameter :{
    RPCCallRequestNode: {
      template: `<div :currentField="currentField">
        <el-form-item label="Timeout in seconds" prop="timeout">
          <el-input-number v-model="currentField.timeout" controls-position="right" :min="0"></el-input-number>
        </el-form-item>
      </div> `, 
      currentField: {
        timeout: 60
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

