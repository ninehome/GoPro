import mock from 'mockjs'

//拦截接口 模拟数据

mock.mock('/api/home/getData','get',function(){
           //拦截 后处理逻辑

           console.log('已经拦截到了')


})