<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="租户名:" prop="username">
          <el-input v-model="formData.username" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="密码:" prop="password">
          <el-input v-model="formData.password" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="运营商昵称:" prop="nickname">
          <el-input v-model="formData.nickname" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="到期字段:" prop="endTime">
          <el-date-picker v-model="formData.endTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="分润方式:" prop="allocation">
          <el-select v-model="formData.allocation" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in allocationOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="分润比例:" prop="allocationProportion">
          <el-input-number v-model="formData.allocationProportion" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'CsOperator'
}
</script>

<script setup>
import {
  createCsOperator,
  updateCsOperator,
  findCsOperator
} from '@/api/csOperator'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const allocationOptions = ref([])
const formData = ref({
            username: '',
            password: '',
            nickname: '',
            endTime: new Date(),
            allocation: undefined,
            allocationProportion: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCsOperator({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recsOperator
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    allocationOptions.value = await getDictFunc('allocation')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createCsOperator(formData.value)
               break
             case 'update':
               res = await updateCsOperator(formData.value)
               break
             default:
               res = await createCsOperator(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
