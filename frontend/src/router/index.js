import { createRouter, createWebHistory } from 'vue-router'
import FirstView from '../views/FirstView.vue'
import SecondView from '../views/SecondView.vue'
const routes = [
  { path: '/', component: FirstView }, // 第一个页面
  { path: '/second', component: SecondView }, // 第二个页面
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router