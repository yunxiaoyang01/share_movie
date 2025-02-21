<template>
  <div class="movie-list-container">
    <div class="header">
      <h2>影视资源列表</h2>
      <el-button type="primary" @click="$router.push('/movie/create')">
        新增影视资源
      </el-button>
    </div>

    <el-table :data="movies" v-loading="loading" style="width: 100%">
      <el-table-column prop="title" label="标题" min-width="150" />
      <el-table-column prop="year" label="年份" width="100" />
      <el-table-column prop="director" label="导演" width="120" />
      <el-table-column prop="type" label="类型" width="100" />
      <el-table-column prop="region" label="地区" width="100" />
      <el-table-column label="标签" width="200">
        <template #default="{ row }">
          <el-tag v-for="tag in row.tags" :key="tag.id" size="small" style="margin-right: 5px">
            {{ tag.name }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="200" fixed="right">
        <template #default="{ row }">
          <el-button-group>
            <el-button type="primary" link @click="$router.push(`/movie/edit/${row.id}`)">
              编辑
            </el-button>
            <el-button type="danger" link @click="handleDelete(row)">
              删除
            </el-button>
          </el-button-group>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import axios from 'axios'

const movies = ref([])
const loading = ref(false)

// 获取影视资源列表
const fetchMovies = async () => {
  loading.value = true
  try {
    const response = await axios.get('/api/movies')
    movies.value = response.data
  } catch (error) {
    ElMessage.error('获取影视资源列表失败')
  } finally {
    loading.value = false
  }
}

// 删除影视资源
const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确认删除该影视资源吗？', '提示', {
      type: 'warning',
      confirmButtonText: '确定',
      cancelButtonText: '取消'
    })

    await axios.delete(`/api/movies/${row.id}`)
    ElMessage.success('删除成功')
    await fetchMovies()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}

onMounted(() => {
  fetchMovies()
})
</script>

<style scoped>
.movie-list-container {
  padding: 20px;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}
</style>