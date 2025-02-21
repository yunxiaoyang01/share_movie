import axios from 'axios'
import { useUserStore } from '../stores/user'

const instance = axios.create({
  baseURL: import.meta.env.VITE_API_URL || '',
  timeout: 5000
})

// 请求拦截器
instance.interceptors.request.use(
  config => {
    const userStore = useUserStore()
    const token = userStore.token
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

// 响应拦截器
instance.interceptors.response.use(
  response => {
    return response
  },
  error => {
    if (error.response) {
      switch (error.response.status) {
        case 401:
          // 未授权，清除token并跳转到登录页
          const userStore = useUserStore()
          userStore.clearToken()
          window.location.href = '/login'
          break
        case 403:
          console.error('没有权限访问该资源')
          break
        case 404:
          console.error('请求的资源不存在')
          break
        case 500:
          console.error('服务器内部错误')
          break
        default:
          console.error('发生未知错误')
      }
    }
    return Promise.reject(error)
  }
)

export default instance