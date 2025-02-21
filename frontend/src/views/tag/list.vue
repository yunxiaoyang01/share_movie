<template>
  <div class="tag-list-container">
    <div class="header">
      <h2>标签管理</h2>
      <el-button type="primary" @click="handleAdd">
        新增标签
      </el-button>
    </div>

    <el-table :data="tags" v-loading="loading" style="width: 100%">
      <el-table-column prop="name" label="标签名称" />
      <el-table-column prop="created_at" label="创建时间" />
      <el-table-column label="操作" width="200" fixed="right">
        <template #default="{ row }">
          <el-button-group>
            <el-button type="primary" link @click="handleEdit(row)">
              编辑
            </el-button>
            <el-button type="danger" link @click="handleDelete(row)">
              删除
            </el-button>
          </el-button-group>
        </template>
      </el-table-column>
    </el-table>

    <!-- 新增/编辑标签对话框 -->
    <el-dialog
      :title="dialogTitle"
      v-model="dialogVisible"
      width="30%"
      @close="resetForm"
    >
      <el-form
        ref="formRef"
        :model="formData"
        :rules="rules"
        label-width="80px"
      >
        <el-form-item label="标签名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入标签名称" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit" :loading="submitting">
            确定
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import axios from 'axios'

const loading = ref(false)
const submitting = ref(false)
const dialogVisible = ref(false)
const dialogTitle = ref('新增标签')
const formRef = ref(null)
const tags = ref([])
const isEdit = ref(false)

const formData = reactive({
  name: ''
})

const rules = {
  name: [{ required: true, message: '请输入标签名称', trigger: 'blur' }]
}

// 获取标签列表
const fetchTags = async () => {
  loading.value = true
  try {
    const response = await axios.get('/api/tags')
    tags.value = response.data
  } catch (error) {
    ElMessage.error('获取标签列表失败')
  } finally {
    loading.value = false
  }
}

// 新增标签
const handleAdd = () => {
  isEdit.value = false
  dialogTitle.value = '新增标签'
  dialogVisible.value = true
}

// 编辑标签
const handleEdit = (row) => {
  isEdit.value = true
  dialogTitle.value = '编辑标签'
  formData.name = row.name
  formData.id = row.id
  dialogVisible.value = true
}

// 删除标签
const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该标签吗？', '提示', {
      type: 'warning'
    })
    await axios.delete(`/api/tags/${row.id}`)
    ElMessage.success('删除成功')
    fetchTags()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}

// 提交表单
const handleSubmit = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (valid) {
      submitting.value = true
      try {
        if (isEdit.value) {
          await axios.put(`/api/tags/${formData.id}`, formData)
          ElMessage.success('更新成功')
        } else {
          await axios.post('/api/tags', formData)
          ElMessage.success('创建成功')
        }
        dialogVisible.value = false
        fetchTags()
      } catch (error) {
        ElMessage.error(isEdit.value ? '更新失败' : '创建失败')
      } finally {
        submitting.value = false
      }
    }
  })
}

// 重置表单
const resetForm = () => {
  if (formRef.value) {
    formRef.value.resetFields()
  }
  formData.name = ''
  formData.id = undefined
}

onMounted(() => {
  fetchTags()
})
</script>

<style scoped>
.tag-list-container {
  padding: 20px;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}
</style>