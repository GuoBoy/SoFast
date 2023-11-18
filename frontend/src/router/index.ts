import { RouteRecordRaw, createRouter, createWebHashHistory } from 'vue-router'
import Index from '@/views/Chat.vue'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    children: [{ path: '/chat', component: Index }]
  }, {
    path: "/history",
    component: () => import("@/views/History.vue")
  }, {
    path: '/users',
    component: () => import('@/components/users/Index.vue')
  }, {
    path: '/user',
    component: () => import('@/components/cneter/Index.vue')
  }, {
    path: '/filemanage',
    component: () => import('@/components/filemanage/Index.vue')
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
