export default {
  NodeParameter :{
    OriginatorTypeFilter:{
      template: `<div :currentField="currentField"><el-form-item label="Originator types filter" prop="Originatortype"><el-select v-model="currentField.Originatortype" multiple placeholder="enter entity type" style="width:100%" ref="typeFocusStatus"><el-option v-for="item in currentField.options" :key="item.value" :label="item.label" :value="item.value"></el-option></el-select></el-form-item></div> `,
      currentField : {
        Originatortype:'',
        options: [{
            value: 'Device',
            label: 'Device'
        }, {
            value: 'Asset',
            label: 'Asset'
        }, {
            value: 'Entity View',
            label: 'Entity View'
        }, {
            value: 'Tenant',
            label: 'Tenant'
        }, {
            value: 'Customer',
            label: 'Customer'
        }, {
            value: 'User',
            label: 'User'
        }, {
            value: 'Dashboard',
            label: 'Dashboard'
        }, {
            value: 'Rule Chain',
            label: 'Rule Chain'
        }, {
            value: 'Rule node',
            label: 'Rule node'
        }]
      },
      linkType:'typeBoolean'
    }
},
linkLabelOptions: {
    typeBoolean: [{
        value: 'True',
        label: 'True'
      }, {
        value: 'False',
        label: 'False'
      }, {
        value: 'Failure',
        label: 'Failure'
    }]
  }
}

