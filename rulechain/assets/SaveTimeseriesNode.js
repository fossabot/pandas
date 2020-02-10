export default {
  NodeParameter :{
    SaveTimeseriesNode: {
      template: `<div :currentField="currentField">
        <el-form-item label="Default TTL in seconds" prop="defauTTl">
          <el-input-number v-model="currentField.defauTTl" controls-position="right" :min="0"></el-input-number>
        </el-form-item>
      </div> `,
      currentField: {
        defauTTl: 0
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

