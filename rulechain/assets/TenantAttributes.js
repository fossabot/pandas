export default {
  NodeParameter :{
    TenantAttributes: {
      template: `<div :currentField="currentField">
            <el-form-item label="Attributes mapping *" prop="">
              <el-checkbox style="zoom:120%;" label=" Latest telemetry" name="type" v-model="currentField.delivery"></el-checkbox><br>
            </el-form-item>
          </el-form>
          <el-form label-width="120px" style="height:120px;overflow-y:auto;" :inline="true">
            <div v-for="(param, index) in currentField.parameters" :key="index">
              <el-row>
                <el-col :span="11">
                  <el-form-item v-if="currentField.delivery === false" label="Source attribute" style="margin-bottom:10px;">
                    <el-input v-model="param.parameter" placeholder="Source attribute" style="width:130px"/>
                  </el-form-item>
                  <el-form-item v-else label="Source telemetry" style="margin-bottom:10px;">
                    <el-input v-model="param.parameter" placeholder="Source telemetry" style="width:130px"/>
                  </el-form-item>
                </el-col>
                <el-col :span="11">
                  <el-form-item label="Target attribute" style="margin-bottom:10px;">
                    <el-input v-model="param.value" placeholder="Target attribute" style="width:130px"/>
                  </el-form-item>
                </el-col>
                <el-button v-if="index === 0" icon="el-icon-plus" size="medium" style="padding: 10px 12px;margin-left:10px;background-color:#f5f7fa;"/>
                <el-button v-if="index !== 0" icon="el-icon-minus" size="medium" style="padding: 10px 12px;margin-left:10px;background-color:#f5f7fa"/>
              </el-row>
            </div> 
          </el-form>
      </div> `,
      currentField: {
        delivery: false,
        type: '',
        typecontain: '',
        associatypes: '',
        parameters: [
          { parameter: 'temperature', value: 'tempo' },
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

