export default {
  NodeParameter :{
    OriginatorAttributes: {
      template: `<div :currentField="currentField">
      <el-form-item label="Client attributes" prop="Clientattributes">
        <el-select
            v-model="currentField.Clientattributes"
            multiple
            filterable
            allow-create
            default-first-option
            placeholder="Client attributes"
            style="width:100%">
                <el-option
                v-for="item in currentField.ClientOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value">
                </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="Shared attributes" prop="Sharedattributes">
        <el-select
            v-model="currentField.Sharedattributes"
            multiple
            filterable
            allow-create
            default-first-option
            placeholder="Shared attributes"
            style="width:100%">
                <el-option
                v-for="item in currentField.SharedOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value">
                </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="Server attributes" prop="Serverattributes">
        <el-select
            v-model="currentField.Serverattributes"
            multiple
            filterable
            allow-create
            default-first-option
            placeholder="Server attributes"
            style="width:100%">
                <el-option
                v-for="item in currentField.ServerdOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value">
                </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="Latest timeseries" prop="Latesttimeseries">
        <el-select
            v-model="currentField.Latesttimeseries"
            multiple
            filterable
            allow-create
            default-first-option
            placeholder="Latest timeseries"
            style="width:100%">
                <el-option
                v-for="item in currentField.LatestOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value">
                </el-option>
        </el-select>
      </el-form-item>
      </div> `, 
      currentField: {
        Clientattributes: [],
        Sharedattributes: [],
        Serverattributes: [],
        Latesttimeseries: [],
        ClientOptions: [],
        SharedOptions: [],
        ServerdOptions: [],
        LatestOptions: [],
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

