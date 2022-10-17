import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './assets/css/style.css'
import 'ant-design-vue/dist/antd.css'
import axios from 'axios'
import './plugins/antui'
import store from './store'

//引入 element 需要加载 css
import 'element-ui/lib/theme-chalk/index.css'
import ElementUI from 'element-ui'

import './api/mock'


// 全部引入
Vue.use(ElementUI)

axios.defaults.baseURL = 'http:localhost:3000/api/v1'
Vue.prototype.$http = axios

Vue.config.productionTip = false

new Vue({
  store,  //挂载 store -- >vuex
  router,
  render: (h) => h(App),
}).$mount('#app')
