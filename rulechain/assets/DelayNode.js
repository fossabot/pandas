export default {
  NodeParameter :{
    DelayNode: {
      template: `<div :currentField="currentField">
        <el-form-item label="Period in seconds" prop="period">
          <el-input-number v-model="currentField.period" controls-position="right" :min="0"></el-input-number>
        </el-form-item>
        <el-form-item label="Maximum pending messages" prop="maxnummess">
          <el-input-number v-model="currentField.maxnummess" controls-position="right" :min="1" :max="100000"></el-input-number>
        </el-form-item>
      </div> `, 
      currentField: {
        period: 60,
        maxnummess: 1000
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

