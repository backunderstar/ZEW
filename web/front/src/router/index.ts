import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      children: [ // 添加children属性来定义嵌套路由
        {
          path: '', // 空路径表示这是默认显示的子路由
          name: 'first',
          component: () => import('../views/home/firstView.vue')
        },
        {
          path: 'about',
          name: 'about',
          component: () => import('../views/home/AboutView.vue')
        }
      ]
    }
  ]
})

// 路由守卫
// to : to将去的路由对象
// from : from将离开的路由对象
// next : next函数，控制导航行为，接受一个参数，指定目标路由
/* router.beforeEach((to, from, next) => {
  if (to.meta.noAuth || userInfoStore.authFromLocal()) {
    next()
  } else {
    router.push('/login')
  }
}) */

export default router
