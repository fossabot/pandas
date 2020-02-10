export default {
  NodeParameter :{
    AssignToCustomerNode: {
      template: `<div :currentField="currentField">
        <el-form-item label="Customer name pattern" prop="customerName">
          <el-input v-model="currentField.customerName"></el-input>
          <span style="font-size:12px;">Customer name pattern, use \${metaKeyName} to substitute variables from metadata</span>
        </el-form-item>
        <el-form-item label="" prop="">
          <el-checkbox style="zoom:120%;" label=" Create new customer if not exists" name="type" v-model="currentField.delivery"></el-checkbox><br>
        </el-form-item>
        <el-form-item label="Customers cache expiration time (sec)" prop="expirationTime">
          <el-input-number v-model="currentField.expirationTime" controls-position="right" :min="0"></el-input-number><br>
          <span style="font-size:12px;">Specifies maximum time interval allowed to store found customer records. 0 value means that records will never expire.</span>
        </el-form-item>
      </div> `, 
      currentField: {
        customerName: '',
        delivery: false,
        expirationTime: 300
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

