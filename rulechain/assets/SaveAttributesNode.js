export default {
  NodeParameter :{
    SaveAttributesNode: {
      template: `<div :currentField="currentField">
        <el-form-item :label="$t('nodeOther.attrange')" prop="attrange">
          <el-select v-model="currentField.attrange" placeholder="" style="width:100%;">
              <el-option
              v-for="item in currentField.options"
              :key="item.value"
              :label="item.label"
              :value="item.value">
              </el-option>
          </el-select>
        </el-form-item>
      </div> `,
      currentField: {
        attrange: '',
        options: [
          {
            value: '客户端属性',
            label: '客户端属性'
          }, {
            value: '服务端属性',
            label: '服务端属性'
          }, {
            value: '共享属性',
            label: '共享属性'
          }
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

