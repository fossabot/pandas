export default {
  NodeParameter :{
    RPCCallReplyNode: {
      template: `<div :currentField="currentField">
        <el-form-item label="Request Id Metadata attribute name" prop="">
          <el-input v-model="currentField.rimaName"></el-input>
        </el-form-item>
      </div> `, 
      currentField: {
        rimaName: 'requestId'
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

