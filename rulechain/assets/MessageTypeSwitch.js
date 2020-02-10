export default {
  NodeParameter :{
    MessageTypeSwitch: {
      template: `<div :currentField="currentField"></div> `,
      currentField: {},
      linkType:'typeAttribute'
    }
  },
  linkLabelOptions: {
    typeAttribute: [{
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
    }, {
      value: 'Alarm Acknowledged',
      label: 'Alarm Acknowledged'
    }, {
      value: 'Alarm Cleared',
      label: 'Alarm Cleared'
    }, {
      value: 'Other',
      label: 'Other'
    }, {
      value: 'Failure',
      label: 'Failure'
    }]
  }
}

