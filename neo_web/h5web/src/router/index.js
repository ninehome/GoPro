import Vue from 'vue'
import VueRouter from 'vue-router'

import Login from '../views/HomeView.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/home',

    name: 'home',

    component: Login,
  },

]

const router = new VueRouter({
  routes,
})

export default router
