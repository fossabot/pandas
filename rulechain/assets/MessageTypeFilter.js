export default {
  NodeParameter :{
    MessageTypeFilter :{
      template: `<div :currentField="currentField">
        <el-form-item label="Message types filter" prop="messagetype">
          <el-select v-model="currentField.messagetype" multiple placeholder="Message type" style="width:100%" ref="typeFocusStatus">
            <el-option v-for="item in currentField.options" 
            :key="item.value" 
            :label="item.label" 
            :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
      </div> `,
      currentField : {
        messagetype:'',
        options: [{
          value: 'Post attributes',
          label: 'Post attributes'
        }, {
          value: 'Post telemetry',
          label: 'Post telemetry'
        }, {
          value: 'RPC Request from Device',
          label: 'RPC Request from Device'
        }, {
          value: 'RPC Request to Device',
          label: 'RPC Request to Device'
        }, {
          value: 'Activity Event',
          label: 'Activity Event'
        }, {
          value: 'Inactivity Event',
          label: 'Inactivity Event'
        }, {
          value: 'Connect Event',
          label: 'Connect Event'
        }, {
          value: 'Disconnect Event',
          label: 'Disconnect Event'
        }, {
          value: 'Entity Created',
          label: 'Entity Created'
        }, {
          value: 'Entity Updated',
          label: 'Entity Updated'
        }, {
          value: 'Entity Deleted',
          label: 'Entity Deleted'
        }, {
          value: 'Entity Assigned',
          label: 'Entity Assigned'
        }, {
          value: 'Entity Unassigned',
          label: 'Entity Unassigned'
        }, {
          value: 'Attributes Updated',
          label: 'Attributes Updated'
        }, {
          value: 'Attributes Deleted',
          label: 'Attributes Deleted'
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

