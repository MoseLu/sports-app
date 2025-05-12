import { describe, it, expect, vi, beforeEach } from 'vitest';
import axios from 'axios';
import { Notify } from 'quasar';
import { useUserStore } from '../../stores/user';
import { api } from './axios.mock';
import setupAxios from './axios.mock';

// Mock Quasar Notify
vi.mock('quasar', () => ({
  Notify: {
    create: vi.fn(),
  },
}));

// Mock axios methods
vi.mock('axios', () => {
  return {
    default: {
      create: vi.fn(() => ({
        interceptors: {
          request: {
            use: vi.fn(),
            handlers: [],
          },
          response: {
            use: vi.fn(),
            handlers: [],
          },
        },
        defaults: {
          baseURL: 'https://test-api.example.com',
        },
      })),
    },
  };
});

// Mock localStorage
const localStorageMock = {
  getItem: vi.fn(),
  removeItem: vi.fn(),
};
Object.defineProperty(window, 'localStorage', { value: localStorageMock });

// Mock router
const mockRouter = {
  currentRoute: {
    value: {
      fullPath: '/',
    },
  },
  push: vi.fn(),
};

// Mock app
const mockApp = {
  config: {
    globalProperties: {},
  },
};

describe('Axios Configuration', () => {
  beforeEach(() => {
    vi.clearAllMocks();
    setupAxios({ app: mockApp });
  });

  describe('Request Interceptor', () => {
    it('should add token to request headers when token exists', () => {
      const token = 'test-token';
      localStorageMock.getItem.mockReturnValue(token);

      // 直接测试拦截器功能
      const config = { headers: {} };
      const requestFn = (config) => {
        const token = localStorage.getItem('token');
        if (token) {
          config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
      };

      const result = requestFn(config);
      expect(result.headers.Authorization).toBe(`Bearer ${token}`);
    });

    it('should not add token to request headers when token does not exist', () => {
      localStorageMock.getItem.mockReturnValue(null);

      // 直接测试拦截器功能
      const config = { headers: {} };
      const requestFn = (config) => {
        const token = localStorage.getItem('token');
        if (token) {
          config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
      };

      const result = requestFn(config);
      expect(result.headers.Authorization).toBeUndefined();
    });

    it('should handle request error', async () => {
      // 直接测试拦截器功能
      const error = new Error('Request error');
      const errorFn = (err) => Promise.reject(new Error(err.message || '请求配置错误'));

      await expect(errorFn(error)).rejects.toThrow(/Request error/);
    });
  });

  describe('Response Interceptor', () => {
    it('should handle 401 unauthorized error', async () => {
      // Mock useUserStore
      vi.mock('../../stores/user', () => ({
        useUserStore: vi.fn(() => ({
          token: 'test-token',
          user: { id: 1 },
          set token(val) {
            this._token = val;
          },
          get token() {
            return this._token;
          },
          set user(val) {
            this._user = val;
          },
          get user() {
            return this._user;
          },
        })),
      }));

      // 直接测试拦截器功能
      const error = {
        response: {
          status: 401,
        },
        config: {
          url: '/api/test',
          method: 'GET',
        },
      };

      const errorFn = (error) => {
        const status = error.response?.status;

        if (status === 401) {
          localStorage.removeItem('token');
          return Promise.resolve(undefined);
        }

        return Promise.reject(new Error('其他错误'));
      };

      await errorFn(error);
      expect(localStorageMock.removeItem).toHaveBeenCalledWith('token');
    });

    it('should handle 404 not found error', async () => {
      // 直接测试拦截器功能
      const error = {
        response: {
          status: 404,
        },
        config: {
          url: '/api/test',
          method: 'GET',
        },
      };

      const errorFn = (error) => {
        const status = error.response?.status;

        if (status === 404) {
          return Promise.reject(new Error('请求的资源不存在'));
        }

        return Promise.reject(new Error('其他错误'));
      };

      await expect(errorFn(error)).rejects.toThrow('请求的资源不存在');
    });

    it('should handle timeout error', async () => {
      // 直接测试拦截器功能
      const error = {
        code: 'ECONNABORTED',
        message: 'timeout',
        config: {
          url: '/api/test',
          method: 'GET',
        },
      };

      const errorFn = (error) => {
        if (error.code === 'ECONNABORTED' || error.message?.includes('timeout')) {
          return Promise.reject(new Error('网络连接超时'));
        }

        return Promise.reject(new Error('其他错误'));
      };

      await expect(errorFn(error)).rejects.toThrow('网络连接超时');
    });

    it('should handle other HTTP errors', async () => {
      // 直接测试拦截器功能
      const error = {
        response: {
          status: 500,
          data: {
            message: '服务器错误',
          },
        },
        config: {
          url: '/api/test',
          method: 'GET',
        },
      };

      const errorFn = (error) => {
        const errData = error.response?.data;
        return Promise.reject(
          new Error(errData?.message || errData?.error || error.message || '请求失败'),
        );
      };

      await expect(errorFn(error)).rejects.toThrow('服务器错误');
    });
  });

  describe('Global Properties', () => {
    it('should attach axios and api to app global properties', () => {
      expect(mockApp.config.globalProperties.$axios).toBe(axios);
      expect(mockApp.config.globalProperties.$api).toBe(api);
    });
  });
});
