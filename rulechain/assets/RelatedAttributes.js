export default {
  NodeParameter :{
    RelatedAttributes: {
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
       <el-form-item :label="$t('nodeAssociat.associatypes')" prop="associatypes">
         <el-row>
           <el-col :span="11">
             <span>{{$t('nodeAssociat.type')}}</span>
           </el-col>
           <el-col :span="11">
             <span>{{$t('nodeAssociat.associatype')}}</span>
           </el-col>
         </el-row>
         <div v-for="(param, index) in currentField.filterParameters" :key="index" style="margin-bottom:10px;">
         <el-row>
           <el-col :span="10">			
       <el-select v-model="param.parameter" placeholder="" style="width:100%">
                 <el-option
                 v-for="item in currentField.restaurants"
                 :key="item.value"
                 :label="item.label"
                 :value="item.value">
                 </el-option>
             </el-select>			
           </el-col>
           <el-col :span="10" style="margin-left:30px;">
             <el-select v-model="param.value" multiple placeholder="" style="width:100%">
                 <el-option
                 v-for="item in currentField.options"
                 :key="item.value"
                 :label="item.label"
                 :value="item.value">
                 </el-option>
             </el-select>
           </el-col>
           <el-button v-if="index === 0" icon="el-icon-plus" size="medium" style="padding: 10px 12px;margin-left:10px;background-color:#f5f7fa;"/>
           <el-button v-if="index !== 0" icon="el-icon-minus" size="medium" style="padding: 10px 12px;margin-left:10px;background-color:#f5f7fa"/>
         </el-row>
         </div>
       </el-form-item>
       <el-form-item label="Attributes mapping *" prop="">
         <el-checkbox style="zoom:120%;" label=" Latest telemetry" name="type" v-model="currentField.delivery"></el-checkbox>
         <el-row>
           <el-col :span="11">
             <span v-if="currentField.delivery === false">Source attribute</span>
             <span v-else>Source telemetry</span>
           </el-col>
           <el-col :span="11">
             <span>Target attribute</span>
           </el-col>
         </el-row>
         <div v-for="(param, index) in currentField.mapParameters" :key="index" style="margin-bottom:10px;">
         <el-row>
           <el-col :span="10">
             <el-input v-model="param.parameter" placeholder="Source"/>
           </el-col>
           <el-col :span="10" style="margin-left:30px;">
             <el-input v-model="param.value" placeholder="Target attribute"/>
           </el-col>
           <el-button v-if="index === 0" icon="el-icon-plus" size="medium" style="padding: 10px 12px;margin-left:10px;background-color:#f5f7fa;"/>
           <el-button v-if="index !== 0" icon="el-icon-minus" size="medium" style="padding: 10px 12px;margin-left:10px;background-color:#f5f7fa"/>
         </el-row>
         </div>
       </el-form-item>
      </div> `, 
      currentField: {
        direction: '从',
        level: '1',
        delivery: false,
        filterParameters: [
          { parameter: '', value: [] },
          { parameter: '', value: [] }
        ],
        mapParameters: [
          { parameter: 'temperature', value: 'tempo' },
          { parameter: '', value: '' }
        ],
        restaurants: [
          {
            value: 'Contains'
          }, {
            value: 'Manages'
          }
        ],
        options: [{
          value: '设备',
          label: '设备'
        }, {
          value: '资产',
          label: '资产'
        }, {
          value: 'Entity View',
          label: 'Entity View'
        }, {
          value: '租户',
          label: '租户'
        }, {
          value: '客户',
          label: '客户'
        }, {
          value: '仪表板',
          label: '仪表板'
        }]
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

