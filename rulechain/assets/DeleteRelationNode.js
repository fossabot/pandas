export default {
  NodeParameter :{
    DeleteRelationNode: {
      template: `<div :currentField="currentField">
        <el-form-item label="" prop="">
          <el-checkbox style="zoom:120%;" label=" Delete relation to specific entity" name="type" v-model="currentField.delRelation"></el-checkbox><br>
          <span style="font-size:12px;">Deletes relation from the originator of the incoming message to the specified entity or list of entities based on direction and type.</span>
        </el-form-item>
        <el-form-item :label="$t('nodeAssociat.direction')" prop="direction">
          <el-select v-model="currentField.direction" placeholder="" style="width:100%;">
            <el-option :label="$t('nodeAssociat.start')" value="start"></el-option>
            <el-option :label="$t('nodeAssociat.end')" value="end"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('nodeAssociat.type')" prop="type" v-if="currentField.delRelation === true">
          <el-select v-model="currentField.type" placeholder="">
            <el-option v-for="item in currentField.alltypes" :key="item.value" :value="item.value" :label="item.label"/>
          </el-select>
        </el-form-item>
        <el-form-item label="" prop="" v-if="currentField.delRelation === true">
          <el-row>
            <el-col :span="11">
              <span>Name pattern *</span>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="11">
              <el-input v-model="currentField.patternName"></el-input>
              <span style="font-size:12px;">Name pattern, use \${metaKeyName} to substitute variables from metadata</span>
            </el-col>
          </el-row>
        </el-form-item>
        <el-form-item label="Relation type pattern" prop="relationType">
          <el-input v-model="currentField.relationType"></el-input>
          <span style="font-size:12px;">Relation type pattern, use \${metaKeyName} to substitute variables from metadata</span>
        </el-form-item>
        <el-form-item label="Entities cache expiration time (sec)" prop="entitycetime">
          <el-input-number v-model="currentField.entitycetime" controls-position="right" :min="0"></el-input-number><br>
          <span style="font-size:12px;">Specifies maximum time interval allowed to store found entity records. 0 value means that records will never expire.</span>
        </el-form-item>
      </div> `,
      currentField: {
        delRelation: false,
        direction: '',
        type: '',
        patternName: '',
        relationType: 'Contains',
        entitycetime: 300,
        alltypes: [
          {
            label: '设备',
            value: '设备'
          }, {
            label: '资产',
            value: '资产'
          }, {
            label: 'Entity View',
            value: 'Entity View'
          }, {
            label: '租户',
            value: '租户'
          }, {
            label: '客户',
            value: '客户'
          }, {
            label: '仪表板',
            value: '仪表板'
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

