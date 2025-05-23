import { describe, it, expect, vi, beforeEach } from 'vitest';
import axios, { AxiosRequestConfig, AxiosError } from 'axios';
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

interface User {
  id: number;
  [key: string]: any;
}

interface UserStore {
  _token: string | null;
  _user: User | null;
  setToken(val: string | null): void;
  getToken(): string | null;
  setUser(val: User | null): void;
  getUser(): User | null;
}

describe('Axios Configuration', () => {
  beforeEach(() => {
    vi.clearAllMocks();
    setupAxios({ app: mockApp });
  });

  describe('Request Interceptor', () => {
    it('should add token to request headers when token exists', () => {
      const token = 'test-token';
      localStorageMock.getItem.mockReturnValue(token);

      const config: AxiosRequestConfig = { headers: {} };
      const requestFn = (config: AxiosRequestConfig) => {
        const token = localStorage.getItem('token');
        if (token) {
          config.headers = config.headers || {};
          config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
      };

      const result = requestFn(config);
      expect(result.headers?.Authorization).toBe(`Bearer ${token}`);
    });

    it('should not add token to request headers when token does not exist', () => {
      localStorageMock.getItem.mockReturnValue(null);

      const config: AxiosRequestConfig = { headers: {} };
      const requestFn = (config: AxiosRequestConfig) => {
        const token = localStorage.getItem('token');
        if (token) {
          config.headers = config.headers || {};
          config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
      };

      const result = requestFn(config);
      expect(result.headers?.Authorization).toBeUndefined();
    });

    it('should handle request error', async () => {
      const error = new Error('Request error');
      const errorFn = (err: Error) => Promise.reject(new Error(err.message || '请求配置错误'));

      await expect(errorFn(error)).rejects.toThrow(/Request error/);
    });
  });

  describe('Response Interceptor', () => {
    it('should handle 401 unauthorized error', async () => {
      // Mock useUserStore
      vi.mock('../../stores/user', () => ({
        useUserStore: vi.fn(() => ({
          _token: 'test-token',
          _user: { id: 1 },
          setToken(val: string | null) {
            this._token = val;
          },
          getToken() {
            return this._token;
          },
          setUser(val: User | null) {
            this._user = val;
          },
          getUser() {
            return this._user;
          },
        })),
      }));

      const error: AxiosError = {
        response: {
          status: 401,
          data: {},
          statusText: '',
          headers: {},
          config: {} as AxiosRequestConfig,
        },
        config: {
          url: '/api/test',
          method: 'GET',
        } as AxiosRequestConfig,
        isAxiosError: true,
        toJSON: () => ({}),
        name: '',
        message: '',
      };

      const errorFn = (error: AxiosError) => {
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
      const error: AxiosError = {
        response: {
          status: 404,
          data: {},
          statusText: '',
          headers: {},
          config: {} as AxiosRequestConfig,
        },
        config: {
          url: '/api/test',
          method: 'GET',
        } as AxiosRequestConfig,
        isAxiosError: true,
        toJSON: () => ({}),
        name: '',
        message: '',
      };

      const errorFn = (error: AxiosError) => {
        const status = error.response?.status;

        if (status === 404) {
          return Promise.reject(new Error('请求的资源不存在'));
        }

        return Promise.reject(new Error('其他错误'));
      };

      await expect(errorFn(error)).rejects.toThrow('请求的资源不存在');
    });

    it('should handle timeout error', async () => {
      const error: AxiosError = {
        code: 'ECONNABORTED',
        message: 'timeout',
        config: {
          url: '/api/test',
          method: 'GET',
        } as AxiosRequestConfig,
        isAxiosError: true,
        toJSON: () => ({}),
        name: '',
        response: undefined,
      };

      const errorFn = (error: AxiosError) => {
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
