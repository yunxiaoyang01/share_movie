import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserStore = defineStore('user', () => {
  const token = ref(localStorage.getItem('token') || '')
  const username = ref('')
  const nickname = ref('')
  const role = ref(1)

  function setToken(newToken) {
    token.value = newToken
    localStorage.setItem('token', newToken)
  }

  function setUserInfo(userInfo) {
    username.value = userInfo.username
    nickname.value = userInfo.nickname
    role.value = userInfo.role
  }

  function clearUserInfo() {
    token.value = ''
    username.value = ''
    nickname.value = ''
    role.value = 1
    localStorage.removeItem('token')
  }

  return {
    token,
    username,
    nickname,
    role,
    setToken,
    setUserInfo,
    clearUserInfo
  }
})