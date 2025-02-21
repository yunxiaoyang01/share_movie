import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Layout',
    component: () => import('../layout/index.vue'),
    children: [
      {
        path: '/movie',
        name: 'MovieList',
        component: () => import('../views/movie/list.vue'),
        meta: { title: '影视资源管理' }
      },
      {
        path: '/movie/create',
        name: 'MovieCreate',
        component: () => import('../views/movie/form.vue'),
        meta: { title: '新增影视资源' }
      },
      {
        path: '/movie/edit/:id',
        name: 'MovieEdit',
        component: () => import('../views/movie/form.vue'),
        meta: { title: '编辑影视资源' }
      },
      {
        path: '/tag',
        name: 'TagList',
        component: () => import('../views/tag/list.vue'),
        meta: { title: '标签管理' }
      }
    ]
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/login/index.vue'),
    meta: { title: '登录' }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
  // 设置页面标题
  document.title = to.meta.title || '影视资源管理系统'

  // 检查是否需要登录
  const token = localStorage.getItem('token')
  if (to.path !== '/login' && !token) {
    next('/login')
  } else {
    next()
  }
})

export default router