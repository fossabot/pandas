export default {
  NodeParameter :{
    CheckRelationFilter:{
      template: `<div :currentField="currentField">
      <el-form-item label="" prop="">
        <el-checkbox style="zoom:120%;" label=" Check relation to specific entity" name="type" v-model="currentField.delivery"></el-checkbox><br>
        <span stye="font-size:12px;">Checks existence of relation to specific entity or to any entity based on direction and relation type.</span>
      </el-form-item>
      <el-form-item :label="$t('nodeAssociat.direction')" prop="direction">
        <el-select v-model="currentField.direction" :placeholder="$t('nodeAssociat.select')"  ref="directionFocusStatus">
          <el-option :label="$t('nodeAssociat.start')" value="start"></el-option>
          <el-option :label="$t('nodeAssociat.end')" value="end"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item v-if="currentField.delivery === true" :label="$t('nodeAssociat.type')" prop="type">
        <el-select v-model="currentField.type" :placeholder="$t('nodeAssociat.select')" ref="typeFocusStatus" :id="currentField.selectEvent">
          <el-option v-for="item in currentField.alltypes" :key="item.value" :value="item.value"/>
        </el-select>
        <el-select v-model="currentField.typecontain" :placeholder="$t('nodeAssociat.select')" ref="typecontainFocusStatus" @change>
          <el-option v-for="con in currentField.tempContains" :key="con.value" :value="con.value"/>
        </el-select>
      </el-form-item>
      <el-form-item :label="$t('nodeAssociat.associatype')" prop="associatype">
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
    </div> `, 
    currentField : {
        direction: '',
        delivery: false,
        type: '',
        typecontain: '',
        associatype: '',
        querySearch:'',
        tempContains: [],
        selectEvent:'selectEvent',
        alltypes: [
          {
            value: '设备',
            contains: [
              {
                label: '设备1',
                value: '设备1'
              }, {
                label: '设备2',
                value: '设备2'
              }
            ]
          }, {
            value: '资产',
            contains: [
              {
                label: '1元',
                value: '1元'
              }, {
                label: '2元',
                value: '2元'
              }
            ]
          }, {
            value: '用户',
            contains: [
              {
                label: '用户1',
                value: '用户1'
              }, {
                label: '用户2',
                value: '用户2'
              }
            ]
          }, {
            value: 'Entity View',
            contains: [
              {
                label: 'Entity1',
                value: 'Entity1'
              }, {
                label: 'Entity2',
                value: 'Entity2'
              }
            ]
          }, {
            value: '仪表盘',
            contains: [
              {
                label: '仪表盘1',
                value: '仪表盘1'
              }, {
                label: '仪表盘2',
                value: '仪表盘2'
              }
            ]
          }
        ],
        restaurants: [
          {
            value: 'Contains'
          }, {
            value: 'Manages'
          }
        ],    
        alltypecontain(containsType) {
          console.log('containsType',containsType)  
          var tempAlltypes = this.alltypes      
          for (var i = 0; i < tempAlltypes.length; i++) {
            if (tempAlltypes[i].value === containsType) {
              this.tempContains = tempAlltypes[i].contains
            }
          }
        },
      },
      linkType:[{
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
}

