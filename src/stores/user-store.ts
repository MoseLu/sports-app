import { defineStore } from 'pinia';
import { api } from '../boot/axios';

interface UserState {
  id: number;
  username: string;
  email: string;
  token: string | null;
  isLoggedIn: boolean;
}

interface UserData {
  id: number;
  username: string;
  email: string;
  token?: string;
}

export const useUserStore = defineStore('user', {
  state: (): UserState => {
    const token = localStorage.getItem('token');
    return {
      id: 0,
      username: '',
      email: '',
      token,
      isLoggedIn: !!token,
    };
  },

  actions: {
    async login(email: string, password: string) {
      const response = await api.post('/auth/login', { email, password });
      this.setUserData(response.data);
      return response;
    },

    async register(username: string, email: string, password: string) {
      const response = await api.post('/auth/register', {
        username,
        email,
        password,
      });
      return response;
    },

    async logout() {
      try {
        await api.post('/auth/logout');
      } catch (error) {
        console.error('Logout error:', error);
      } finally {
        this.clearUserData();
      }
    },

    async getProfile() {
      const response = await api.get('/users/profile');
      this.setUserData(response.data);
      return response;
    },

    setUserData(data: UserData) {
      this.id = data.id;
      this.username = data.username;
      this.email = data.email;
      this.token = data.token || null;
      this.isLoggedIn = true;
      if (data.token) {
        localStorage.setItem('token', data.token);
      }
    },

    clearUserData() {
      this.id = 0;
      this.username = '';
      this.email = '';
      this.token = null;
      this.isLoggedIn = false;
      localStorage.removeItem('token');
    },
  },

  getters: {
    getUserInfo: (state) => ({
      id: state.id,
      username: state.username,
      email: state.email,
    }),
  },
});
