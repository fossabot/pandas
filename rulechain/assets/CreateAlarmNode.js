<template>
  <el-dialog :visible="dialogVisible" :before-close="initDialog" title="添加规则节点" top="10vh">
    <el-form :model="ruleForm" :rules="rules" label-position="top" ref="ruleForm" label-width="120px" style="padding:30px;" class="AssignToCustomerNode">
      <el-form-item label="名称" prop="name">
        <el-input v-model="ruleForm.name" ref="nameFocusStatus"></el-input>
      </el-form-item>
      <el-form-item label="描述" prop="desc">
        <el-input type="textarea" v-model="ruleForm.desc"></el-input>
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button @click="initDialog">取消</el-button>
      <el-button type="primary" @click="surechange()">确定</el-button>
    </div>
  </el-dialog>
</template>

<script>
export default {
  name: 'AssignToCustomerNode',
  data() {
    return {
      ruleForm: {
        name: '',
        desc: ''
      },
      ruleFormCopy: {
        name: '',
        desc: ''
      },
      rules: {
        name: [
          { required: true, message: '名称必填', trigger: 'blur' }
        ]
      },
      deleteNodeFlag: true,
      dialogVisible:false,
      firstTime:true
    }
  },
  mounted() {
  },
  methods: {
    surechange() {
      if (!this.ruleForm.name) {
        alert('Please input name')
        this.$refs.nameFocusStatus.focus()
      } else {
        this.ruleFormCopy = JSON.parse(JSON.stringify(this.ruleForm)) 
        this.$emit('onUpdateAssignToCustomData', this.ruleForm)
        this.dialogVisible = false
        this.firstTime = false
      }
    },
    initDialog() {
      if (this.deleteNodeFlag) {
        this.$emit('onDialogATCvisibleChange', this.deleteNodeFlag) // 关闭对话框
      } else {
        this.$emit('onUpdateAssignToCustomData', this.ruleFormCopy)
      }
      this.dialogVisible = false
    }
  }
}
</script>
<style>
</style>
