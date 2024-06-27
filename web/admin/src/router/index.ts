import { createRouter, createWebHistory } from 'vue-router'
import { useUserInfoStore } from '@/stores/userInfo.store'
import AdminView from '../views/home/AdminView.vue'

const userInfoStore = useUserInfoStore()
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/home',
      name: 'home',
      children: [
        {
          path: '',
          name: 'admin',
          redirect: '/home/dashboard',
          component: AdminView,
          children: [
            {
              path: 'dashboard',
              name: 'dashboard',
              component: () => import('../components/home/DashBoard.vue')
            },
            {
              path: 'user',
              name: 'user',
              component: () => import('../components/user/userManagement.vue')
            },
            {
              path: 'article',
              name: 'article',
              component: () =>
                import('../components/article/articleManagement.vue')
            },
            {
              path: 'category',
              name: 'category',
              component: () =>
                import('../components/article/categoryManagement.vue')
            },
            {
              path: 'label',
              name: 'label',
              component: () =>
                import('../components/article/labelManagement.vue')
            }
          ]
        },
        {
          path: 'login',
          name: 'login',
          meta: { noAuth: true },
          component: () => import('../views/home/LoginView.vue')
        }
      ]
    }
  ]
})

// 路由守卫
// to : to将去的路由对象
// from : from将离开的路由对象
// next : next函数，控制导航行为，接受一个参数，指定目标路由
router.beforeEach((to, from, next) => {
  if (to.meta.noAuth || userInfoStore.authFromLocal()) {
    next()
  } else {
    /* router.push('/login') */
    next()
  }
})

export default router
