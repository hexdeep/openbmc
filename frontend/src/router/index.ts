import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
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
          path: '/images',
          component: () => import('@/views/ImageList.vue'),
        },
        {
          path: '/settings',
          component: () => import('@/views/SettingView.vue'),
        },
      ],
    },
  ],
})

export default router
