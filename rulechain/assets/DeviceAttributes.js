export default {
  NodeParameter :{
    DeviceAttributes: {
      template: `<div :currentField="currentField">	  
        <el-form-item label="Device relations query" prop="query">
          <div style="display:flex">
              <div>
              <span>{{$t('nodeAssociat.direction')}} *&nbsp;</span><br>
              <el-select v-model="currentField.direction" placeholder="" style="width:105px;">
                  <el-option :label="$t('nodeAssociat.start')" value="start"></el-option>
                  <el-option :label="$t('nodeAssociat.end')" value="end"></el-option>
              </el-select>
              </div>
              <div style="margin-left:15px;">
              <span>Max relation level</span><br>
                <el-input-number v-model="currentField.level" controls-position="right" :min="1"></el-input-number>
              </div>
          </div>
        </el-form-item>
        <el-form-item :label="$t('nodeAssociat.type')" prop="associatype">
          <el-select
          v-model="currentField.associatype"
          filterable
          allow-create
          default-first-option
          placeholder=""
          style="width:100%">
              <el-option
              v-for="item in currentField.restaurants"
              :key="item.value"
              :label="item.label"
              :value="item.value">
              </el-option>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('nodeAssociat.associatype')" prop="devicetype">
          <el-select
              v-model="currentField.devicetype"
              multiple
              filterable
              allow-create
              default-first-option
              placeholder=""
              style="width:100%">
                  <el-option
                  v-for="item in currentField.deviceOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                  </el-option>
          </el-select>
        </el-form-item>
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
        direction: '从',
        level: '1',
        associatype: '',
        devicetype: [],
        Clientattributes: [],
        Sharedattributes: [],
        Serverattributes: [],
        Latesttimeseries: [],
        restaurants: [
          {
            value: 'Contains',
            label: 'Contains'
          }, {
            value: 'Manages',
            label: 'Manages'
          }
        ],
        deviceOptions: [
          {
            value: 'default',
            label: 'default'
          }
        ],
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

