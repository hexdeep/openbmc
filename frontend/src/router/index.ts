import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      component: () => import('@/views/Login.vue'),
    },
    {
      path: '/settings',
      component: () => import('@/views/Setting.vue'),
    },
    {
      path: '/',
      component: () => import('@/views/NavbarLayout.vue'),
      children: [
        {
          path: 'dashboard',
          component: () => import('@/views/dashboard/Layout.vue'),
        },
        {
          path: 'slots',
          component: () => import('@/views/Slot.vue'),
        },
        {
          path: 'switch',
          component: () => import('@/views/switch/Layout.vue'),
        },
        {
          path: 'images',
          component: () => import('@/views/ImageList.vue'),
        },
        {
          path: 'file',
          component: () => import('@/views/FileManage.vue'),
        },
        {
          path: 'logs',
          component: () => import('@/views/Log.vue'),
        },
      ],
    },
  ],
})

export default router
