import { createRouter, createWebHistory } from 'vue-router'
export const routes = [
  { path: '/login', component: () => import('../pages/Login.vue') }
]
const router = createRouter({ history: createWebHistory(), routes })
export default router
