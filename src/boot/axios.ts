// src/boot/axios.ts
import { defineBoot } from '#q-app/wrappers';
import axios, { type AxiosInstance, type AxiosError, type AxiosRequestConfig } from 'axios';
import { Notify } from 'quasar';
import { useUserStore } from '../stores/user';

interface ErrorResponse {
  code?: string;
  error?: string;
  message?: string;
}

interface RetryConfig extends AxiosRequestConfig {
  retry?: number;
  retryDelay?: number;
  __retryCount?: number;
}

// 创建 axios 实例
export const api: AxiosInstance = axios.create({
  baseURL: 'https://redamancy.com.cn/api',
  timeout: 30000, // 增加超时时间到 30 秒
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json',
  },
});

// 添加重试拦截器
api.interceptors.response.use(undefined, async (err) => {
  const config = err.config as RetryConfig;
  if (!config || !config.retry) {
    return Promise.reject(new Error(err.message || '请求失败'));
  }

  config.__retryCount = config.__retryCount || 0;

  if (config.__retryCount >= config.retry) {
    return Promise.reject(new Error(err.message || '请求失败'));
  }

  config.__retryCount += 1;
  const backoff = new Promise((resolve) => {
    setTimeout(() => {
      resolve(null);
    }, config.retryDelay || 1000);
  });

  await backoff;
  return api(config);
});

export default defineBoot(({ app, router }) => {
  // 2️⃣ 请求拦截器：添加 Token
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

  // 3️⃣ 响应拦截器：用到 useUserStore 和 router
  api.interceptors.response.use(
    (res) => res,
    async (error: AxiosError) => {
      const status = error.response?.status;
      const userStore = useUserStore();
      const path = error.config?.url?.replace(api.defaults.baseURL || '', '') || '';
      const method = (error.config?.method || '').toUpperCase();

      // —— ① 认证失败 (401) ——
      if (status === 401) {
        // 清空 Pinia 状态 + 本地存储
        userStore.token = null;
        userStore.user = null;
        localStorage.removeItem('token');

        // 非登录页时跳转
        const current = router.currentRoute.value.fullPath;
        if (!['/login', '/auth/login'].includes(current)) {
          await router.push('/login');
        }

        // "吞掉"错误：业务层不会 catch，也不会在控制台报红
        return Promise.resolve(undefined);
      }

      // —— ② 资源不存在 (404) ——
      if (status === 404) {
        // 显示错误通知
        Notify.create({
          type: 'negative',
          message: '请求的资源不存在',
          position: 'top',
          timeout: 3000,
        });
        return Promise.reject(new Error('请求的资源不存在'));
      }

      // —— ③ 超时错误 ——
      if (error.code === 'ECONNABORTED' || error.message?.includes('timeout')) {
        Notify.create({
          type: 'warning',
          message: '网络连接超时，正在重试...',
          position: 'top',
          timeout: 3000,
        });
        return Promise.reject(new Error('网络连接超时'));
      }

      // —— ④ 其它 HTTP 错误 ——
      // 准备请求/响应体文本（日志上报用）
      let requestBody = '无请求数据';
      if (error.config?.data) {
        try {
          requestBody = JSON.stringify(error.config.data);
        } catch (e) {
          console.error('解析请求数据失败:', e);
        }
      } else if (error.config?.params) {
        try {
          requestBody = JSON.stringify(error.config.params);
        } catch (e) {
          console.error('解析请求参数失败:', e);
        }
      }
      let responseBody = '无响应数据';
      if (error.response?.data) {
        try {
          responseBody = JSON.stringify(error.response.data);
        } catch (e) {
          console.error('解析响应数据失败:', e);
        }
      }

      // 上报错误（跳过 /api/errors 自身的请求）
      if (!error.config?.url?.includes('/api/errors')) {
        try {
          await api.post('/errors', {
            error_type: 'http_error',
            message: error.message || '请求失败',
            status_code: status || 500,
            path,
            method,
            request_body: requestBody,
            response_body: responseBody,
          });
        } catch (e) {
          console.error('日志上报失败：', e);
        }
      }

      // 全局通知
      const errData = error.response?.data as ErrorResponse;
      Notify.create({
        type: 'negative',
        message: errData.message || errData.error || error.message || '请求失败',
        position: 'top',
        timeout: 3000,
      });

      // 继续抛给业务，让其 decide 是否要 catch
      return Promise.reject(
        new Error(errData.message || errData.error || error.message || '请求失败'),
      );
    },
  );

  // 4️⃣ 挂载到 Vue 实例
  app.config.globalProperties.$axios = axios;
  app.config.globalProperties.$api = api;
});
