import axios from 'axios';

// 创建 axios 实例
const instance = axios.create({
  baseURL: 'http://localhost:8080', // 默认的请求地址
});

// 请求拦截器
instance.interceptors.request.use(
  config => {
    // 可以在这里添加 token 或其他请求前逻辑
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

// 响应拦截器
instance.interceptors.response.use(
  response => {
    // 处理响应数据
    return response;
  },
  (error) => {
    if (error.response) {
      const { status } = error.response;
      if (status === 401) {
        // Token 无效或已过期
        alert("未授权或登录已过期，请重新登录");
        // 可选：重定向到登录页
        window.location.href = "/login";
      }
    }
    return Promise.reject(error);
  }
);

export default instance;
