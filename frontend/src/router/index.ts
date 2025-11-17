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
          path: 'server',
          component: () => import('@/views/ServerView.vue'),
        },
        {
          path: 'devices',
          component: () => import('@/views/DeviceList.vue'),
        },
        {
          path: 'images',
          component: () => import('@/views/ImageList.vue'),
        },
        {
          path: 'file',
          component: () => import('@/views/FileManage.vue'),
        },
      ],
    },
  ],
})

export default router
