<template>
  <div class="movie-form-container">
    <div class="header">
      <h2>{{ isEdit ? '编辑' : '新增' }}影视资源</h2>
    </div>

    <el-form
      ref="formRef"
      :model="formData"
      :rules="rules"
      label-width="100px"
      class="movie-form"
    >
      <el-form-item label="标题" prop="title">
        <el-input v-model="formData.title" placeholder="请输入影视标题" />
      </el-form-item>

      <el-form-item label="描述" prop="description">
        <el-input
          v-model="formData.description"
          type="textarea"
          rows="4"
          placeholder="请输入影视描述"
        />
      </el-form-item>

      <el-form-item label="封面" prop="cover">
        <el-input v-model="formData.cover" placeholder="请输入封面图片URL" />
      </el-form-item>

      <el-row :gutter="20">
        <el-col :span="8">
          <el-form-item label="年份" prop="year">
            <el-input-number
              v-model="formData.year"
              :min="1900"
              :max="2100"
              placeholder="请选择年份"
            />
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="类型" prop="type">
            <el-select v-model="formData.type" placeholder="请选择类型">
              <el-option label="电影" value="电影" />
              <el-option label="电视剧" value="电视剧" />
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="地区" prop="region">
            <el-input v-model="formData.region" placeholder="请输入地区" />
          </el-form-item>
        </el-col>
      </el-row>

      <el-form-item label="导演" prop="director">
        <el-input v-model="formData.director" placeholder="请输入导演" />
      </el-form-item>

      <el-form-item label="主演" prop="actors">
        <el-input v-model="formData.actors" placeholder="请输入主演" />
      </el-form-item>

      <el-form-item label="标签" prop="tags">
        <el-select
          v-model="formData.tags"
          multiple
          filterable
          placeholder="请选择标签"
        >
          <el-option
            v-for="tag in tags"
            :key="tag.id"
            :label="tag.name"
            :value="tag.id"
          />
        </el-select>
      </el-form-item>

      <el-form-item label="网盘链接">
        <div v-for="(link, index) in formData.links" :key="index" class="link-item">
          <el-row :gutter="20">
            <el-col :span="6">
              <el-select v-model="link.platform" placeholder="选择平台">
                <el-option label="百度网盘" value="百度网盘" />
                <el-option label="阿里云盘" value="阿里云盘" />
              </el-select>
            </el-col>
            <el-col :span="8">
              <el-input v-model="link.url" placeholder="分享链接" />
            </el-col>
            <el-col :span="6">
              <el-input v-model="link.password" placeholder="提取密码" />
            </el-col>
            <el-col :span="4">
              <el-button type="danger" @click="removeLink(index)">删除</el-button>
            </el-col>
          </el-row>
        </div>
        <el-button type="primary" @click="addLink">添加链接</el-button>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="handleSubmit" :loading="loading">
          保存
        </el-button>
        <el-button @click="$router.back()">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import axios from 'axios'

const route = useRoute()
const router = useRouter()
const formRef = ref(null)
const loading = ref(false)
const tags = ref([])
const isEdit = ref(false)

const formData = reactive({
  title: '',
  description: '',
  cover: '',
  year: new Date().getFullYear(),
  director: '',
  actors: '',
  type: '',
  region: '',
  tags: [],
  links: []
})

const rules = {
  title: [{ required: true, message: '请输入影视标题', trigger: 'blur' }],
  type: [{ required: true, message: '请选择类型', trigger: 'change' }],
  year: [{ required: true, message: '请选择年份', trigger: 'change' }]
}

// 获取所有标签
const fetchTags = async () => {
  try {
    const response = await axios.get('/api/tags')
    tags.value = response.data
  } catch (error) {
    ElMessage.error('获取标签列表失败')
  }
}

// 获取影视资源详情
const fetchMovie = async (id) => {
  try {
    const response = await axios.get(`/api/movies/${id}`)
    Object.assign(formData, response.data)
    formData.tags = response.data.tags.map(tag => tag.id)
  } catch (error) {
    ElMessage.error('获取影视资源详情失败')
  }
}

// 添加网盘链接
const addLink = () => {
  formData.links.push({
    platform: '',
    url: '',
    password: ''
  })
}

// 删除网盘链接
const removeLink = (index) => {
  formData.links.splice(index, 1)
}

// 提交表单
const handleSubmit = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        const method = isEdit.value ? 'put' : 'post'
        const url = isEdit.value ? `/api/movies/${route.params.id}` : '/api/movies'
        await axios[method](url, formData)
        ElMessage.success(`${isEdit.value ? '更新' : '创建'}成功`)
        router.push('/movie')
      } catch (error) {
        ElMessage.error(error.response?.data?.error || `${isEdit.value ? '更新' : '创建'}失败`)
      } finally {
        loading.value = false
      }
    }
  })
}

onMounted(async () => {
  await fetchTags()
  if (route.params.id) {
    isEdit.value = true
    await fetchMovie(route.params.id)
  }
})
</script>

<style scoped>
.movie-form-container {
  padding: 20px;
}

.header {
  margin-bottom: 20px;
}

.movie-form {
  max-width: 800px;
}

.link-item {
  margin-bottom: 10px;
}
</style>