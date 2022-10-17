// 这里是vuex 的配置入口
// 1.第一步引入Vue
// 2.vuex 全局注入 vue
//3.创建 vuex 实例
import Vue from 'vue'
import Vuex from 'vuex'

import tab from './tab'

Vue.use(Vuex)


//export default 对外导出实例  然后再挂载到 main.js 的Vue
export default new Vuex.Store({
    modules: {
      tab,
    },
  })
  

