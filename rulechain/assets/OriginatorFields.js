export default {
  NodeParameter :{
    OriginatorFields: {
      template: `<div :currentField="currentField">
        <el-form-item label="" prop="">
          <span>Fields mapping *</span>   
        </el-form-item>
        <el-form label-width="110px" :inline="true" style="height:120px;overflow-y:auto">
          <div v-for="(param, index) in currentField.parameters" :key="index">
            <el-row>
              <el-col :span="11">
                <el-form-item label="Source field" style="margin-bottom:10px;">
                  <el-input placeholder="Source field" v-model="param.parameter" style="width:150px"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="11">
                <el-form-item label="Target attribute" style="margin-bottom:10px;">
                  <el-input placeholder="Target attribute" v-model="param.value" style="width:150px"></el-input>
                </el-form-item>
              </el-col>
              <el-button v-if="index === 0" icon="el-icon-plus" size="medium" style="padding: 10px 12px;margin-left:10px;background-color:#f5f7fa;"/>
              <el-button v-if="index !== 0" icon="el-icon-minus" size="medium" style="padding: 10px 12px;margin-left:10px;background-color:#f5f7fa"/>
            </el-row>
          </div> 
        </el-form>
      </div> `, 
      currentField: {
        parameters: [
          { parameter: '', value: '' },
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

