import axios from 'axios';

// 创建一个简单的 axios 实例用于测试
export const api = axios.create({
  baseURL: 'https://test-api.example.com',
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json',
  },
});

// 添加请求拦截器
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (err) => Promise.reject(new Error(err.message || '请求配置错误')),
);

// 添加响应拦截器
api.interceptors.response.use(
  (res) => res,
  (error) => {
    const status = error.response?.status;

    // 401 未授权
    if (status === 401) {
      localStorage.removeItem('token');
      return Promise.resolve(undefined);
    }

    // 404 资源不存在
    if (status === 404) {
      return Promise.reject(new Error('请求的资源不存在'));
    }

    // 超时错误
    if (error.code === 'ECONNABORTED' || error.message?.includes('timeout')) {
      return Promise.reject(new Error('网络连接超时'));
    }

    // 其他错误
    const errData = error.response?.data;
    return Promise.reject(
      new Error(errData?.message || errData?.error || error.message || '请求失败'),
    );
  },
);

// 导出默认函数
export default function setupAxios({ app }: { app: any }) {
  app.config.globalProperties.$axios = axios;
  app.config.globalProperties.$api = api;
}
