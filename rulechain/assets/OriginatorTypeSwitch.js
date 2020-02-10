export default {
  NodeParameter :{
    OriginatorTypeSwitch: {
      template: `<div :currentField="currentField"></div> `,
      currentField: {},
      linkType:'typeDevice'
    }
  },
  linkLabelOptions: {
    typeDevice: [{
      value: 'Device',
      label: 'Device'
    }, {
      value: 'Asset',
      label: 'Asset'
    }, {
      value: 'Entity view',
      label: 'Entity'
    }, {
      value: 'Tenant',
      label: 'Tenant'
    }, {
      value: 'Customer',
      label: 'Customer'
    }, {
      value: 'User',
      label: 'User'
    }, {
      value: 'Dashboard',
      label: 'Dashboard'
    }, {
      value: 'Rule chain',
      label: 'Rule chain'
    }, {
      value: 'Rule node',
      label: 'Rule node'
    }, {
      value: 'Failure',
      label: 'Failure'
    }]
  }
}

