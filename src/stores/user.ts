import { defineStore } from 'pinia';
import { api } from 'boot/axios';
import type { AxiosError } from 'axios';
import { handleError, showSuccess } from '../services/error-handler';

interface User {
  id: number;
  username: string;
  email: string;
  nickname: string;
  bio: string;
}

interface PageHistory {
  name: string;
  path: string;
  timestamp: string;
}

interface UserState {
  user: User | null;
  token: string | null;
  loading: boolean;
  error: string | null;
  history: PageHistory[];
}

export const useUserStore = defineStore('user', {
  state: (): UserState => ({
    user: null,
    token: null,
    loading: false,
    error: null,
    history: [],
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
    currentUser: (state) => state.user,
    isLoading: (state) => state.loading,
  },

  actions: {
    async init() {
      const token = localStorage.getItem('token');
      if (token) {
        this.token = token;
        try {
          await this.fetchUser();
        } catch (error) {
          console.error('Token validation failed:', error);
          this.token = null;
          localStorage.removeItem('token');
        }
      }
    },

    async login(data: { username: string; password: string }) {
      this.loading = true;
      this.error = null;
      try {
        const response = await api.post('/auth/login', data);
        const { token } = response.data;

        // 保存 token 到 localStorage
        if (token) {
          localStorage.setItem('token', token);
          this.token = token;
          // 设置 axios 默认 header
          api.defaults.headers.common['Authorization'] = `Bearer ${token}`;

          // 立即获取用户信息以验证 token
          try {
            await this.fetchUser();
          } catch (error) {
            console.error('Token validation failed:', error);
            this.token = null;
            localStorage.removeItem('token');
            throw new Error('登录验证失败，请重试');
          }
        } else {
          throw new Error('登录失败：未收到有效的token');
        }

        return response;
      } catch (error) {
        console.error('Login error:', error);
        handleError(error as AxiosError);
        throw error;
      } finally {
        this.loading = false;
      }
    },

    async register(data: { username: string; email: string; password: string }) {
      this.loading = true;
      this.error = null;
      try {
        const response = await api.post('/auth/register', data);
        this.token = response.data.token;
        this.user = response.data.user;
        if (this.token) {
          localStorage.setItem('token', this.token);
        }
      } catch (error) {
        handleError(error as AxiosError);
        throw error;
      } finally {
        this.loading = false;
      }
    },

    async logout() {
      this.loading = true;
      this.error = null;
      try {
        await api.post('/auth/logout');
        this.token = null;
        this.user = null;
        localStorage.removeItem('token');
        showSuccess('退出成功');
      } catch (error) {
        handleError(error as AxiosError);
        throw error;
      } finally {
        this.loading = false;
      }
    },

    async fetchUser() {
      this.error = null;
      try {
        const response = await api.get('/users/profile');
        if (response?.data) {
          this.user = response.data;
        } else {
          console.error('获取用户信息失败：响应数据为空');
          this.user = null;
        }
      } catch (error) {
        console.error('获取用户信息失败:', error);
        this.user = null;
        handleError(error as AxiosError);
      }
    },

    async updateProfile(userData: Partial<User>) {
      this.loading = true;
      this.error = null;
      try {
        const response = await api.put('/users/profile', userData);
        this.user = response.data;
      } catch (error) {
        handleError(error as AxiosError);
        throw error;
      } finally {
        this.loading = false;
      }
    },

    addToHistory(page: PageHistory) {
      this.history.push(page);
    },
  },
});
