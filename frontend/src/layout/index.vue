<template>
  <el-container class="layout-container">
    <!-- 侧边栏 -->
    <el-aside width="200px">
      <el-menu
        :default-active="route.path"
        class="el-menu-vertical"
        router
      >
        <el-menu-item index="/movie">
          <el-icon><VideoCamera /></el-icon>
          <span>影视资源管理</span>
        </el-menu-item>
        <el-menu-item index="/tag">
          <el-icon><Collection /></el-icon>
          <span>标签管理</span>
        </el-menu-item>
      </el-menu>
    </el-aside>

    <!-- 主要内容区域 -->
    <el-container>
      <el-header height="60px">
        <div class="header-left">
          <h2>影视资源管理系统</h2>
        </div>
        <div class="header-right">
          <el-dropdown @command="handleCommand">
            <span class="el-dropdown-link">
              {{ userStore.nickname || userStore.username }}
              <el-icon class="el-icon--right"><arrow-down /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="logout">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>

      <el-main>
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import { VideoCamera, Collection } from '@element-plus/icons-vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '../stores/user'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

const handleCommand = (command) => {
  if (command === 'logout') {
    localStorage.removeItem('token')
    router.push('/login')
  }
}
</script>

<style scoped>
.layout-container {
  height: 100vh;
}

.el-menu-vertical {
  height: 100%;
  border-right: none;
}

.el-header {
  background-color: #fff;
  border-bottom: 1px solid #dcdfe6;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-right {
  margin-right: 20px;
}

.el-dropdown-link {
  cursor: pointer;
  display: flex;
  align-items: center;
}
</style>