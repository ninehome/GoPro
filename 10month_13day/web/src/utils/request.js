import axios from 'axios'

//创建 和配置请求实例
const http  = axios.create({

    baseURL:'/api',
    timeout:10000 , //超时10 S

})



// 添加请求拦截器
axios.interceptors.request.use(function (config) {
    // 在发送请求之前做些什么
    return config;
  }, function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
  });




// 添加响应拦截器
axios.interceptors.response.use(function (response) {
    // 对响应数据做点什么
    return response;
  }, function (error) {
    // 对响应错误做点什么
    return Promise.reject(error);
  });




//对外导出实例
export default http