import Vue from 'vue'
import VueRouter from 'vue-router'
import Mian from '../views/MainView.vue'
import User from '../views/User.vue'
import Home from '../views/HomeView.vue'
import Mall from '../views/MallView.vue'
import page1 from '../views/PageOne.vue'
import page2 from '../views/PageTwo.vue'
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    component: Mian,
    redirect:'/home', // 如果 匹配到是 / 就重定向到 /home
    children: [  //子路由
      {
        path: '/home',
        component: Home,
      },
      {
        path: '/user',
        component: User,
      },
      {
        path: '/mall',
        component: Mall,
      }
      ,
      {
        path: '/page1',
        component: page1,
      }
      ,
      {
        path: '/page2',
        component: page2,
      }
    ],
  },
]

const router = new VueRouter({
  routes,
})

export default router
