export default{
// es6语法 对象导出
  state:{
    isCollapse:false  //控制 菜单 展开 和收起

  },
  mutations:{   //更改 Vuex 的 store 中的状态的唯一方法是提交 mutation
     collapsMenu(state){
        state.isCollapse = !state.isCollapse
     }
  }
  
}