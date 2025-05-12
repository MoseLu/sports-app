import { describe, it, expect, beforeEach, vi } from 'vitest';
import { setActivePinia, createPinia } from 'pinia';
import { useUserStore } from '../user-store';
import { api } from '../../boot/axios';

// Mock axios
vi.mock('../../boot/axios', () => ({
  api: {
    post: vi.fn(),
    get: vi.fn(),
  },
}));

// Mock localStorage
const localStorageMock = (() => {
  let store: Record<string, string> = {};
  return {
    getItem: vi.fn((key: string) => store[key] || null),
    setItem: vi.fn((key: string, value: string) => {
      store[key] = value;
    }),
    removeItem: vi.fn((key: string) => {
      delete store[key];
    }),
    clear: vi.fn(() => {
      store = {};
    }),
  };
})();

Object.defineProperty(window, 'localStorage', { value: localStorageMock });

describe('User Store', () => {
  beforeEach(() => {
    setActivePinia(createPinia());
    vi.clearAllMocks();
  });

  describe('State', () => {
    it('should initialize with default values', () => {
      const store = useUserStore();
      expect(store.id).toBe(0);
      expect(store.username).toBe('');
      expect(store.email).toBe('');
      expect(store.token).toBeNull();
      expect(store.isLoggedIn).toBe(false);
    });
  });

  describe('Actions', () => {
    it('should login successfully', async () => {
      const mockResponse = {
        data: {
          id: 1,
          username: 'testuser',
          email: 'test@example.com',
          token: 'test-token',
        },
      };
      vi.mocked(api.post).mockResolvedValueOnce(mockResponse);

      const store = useUserStore();
      const response = await store.login('test@example.com', 'password');

      expect(api.post).toHaveBeenCalledWith('/auth/login', {
        email: 'test@example.com',
        password: 'password',
      });
      expect(store.id).toBe(1);
      expect(store.username).toBe('testuser');
      expect(store.email).toBe('test@example.com');
      expect(store.token).toBe('test-token');
      expect(store.isLoggedIn).toBe(true);
      expect(localStorage.setItem).toHaveBeenCalledWith('token', 'test-token');
    });

    it('should register successfully', async () => {
      const mockResponse = { data: { message: 'Registration successful' } };
      vi.mocked(api.post).mockResolvedValueOnce(mockResponse);

      const store = useUserStore();
      const response = await store.register('testuser', 'test@example.com', 'password');

      expect(api.post).toHaveBeenCalledWith('/auth/register', {
        username: 'testuser',
        email: 'test@example.com',
        password: 'password',
      });
      expect(response.data.message).toBe('Registration successful');
    });

    it('should logout successfully', async () => {
      const store = useUserStore();
      store.setUserData({
        id: 1,
        username: 'testuser',
        email: 'test@example.com',
        token: 'test-token',
      });

      await store.logout();

      expect(api.post).toHaveBeenCalledWith('/auth/logout');
      expect(store.id).toBe(0);
      expect(store.username).toBe('');
      expect(store.email).toBe('');
      expect(store.token).toBeNull();
      expect(store.isLoggedIn).toBe(false);
      expect(localStorage.removeItem).toHaveBeenCalledWith('token');
    });

    it('should get profile successfully', async () => {
      const mockResponse = {
        data: {
          id: 1,
          username: 'testuser',
          email: 'test@example.com',
        },
      };
      vi.mocked(api.get).mockResolvedValueOnce(mockResponse);

      const store = useUserStore();
      const response = await store.getProfile();

      expect(api.get).toHaveBeenCalledWith('/users/profile');
      expect(store.id).toBe(1);
      expect(store.username).toBe('testuser');
      expect(store.email).toBe('test@example.com');
    });

    it('should handle login failure', async () => {
      const mockError = new Error('Login failed');
      vi.mocked(api.post).mockRejectedValueOnce(mockError);

      const store = useUserStore();
      await expect(store.login('test@example.com', 'password')).rejects.toThrow('Login failed');
      expect(store.isLoggedIn).toBe(false);
      expect(store.token).toBeNull();
    });

    it('should handle register failure', async () => {
      const mockError = new Error('Registration failed');
      vi.mocked(api.post).mockRejectedValueOnce(mockError);

      const store = useUserStore();
      await expect(store.register('testuser', 'test@example.com', 'password')).rejects.toThrow(
        'Registration failed',
      );
    });

    it('should handle getProfile failure', async () => {
      const mockError = new Error('Profile fetch failed');
      vi.mocked(api.get).mockRejectedValueOnce(mockError);

      const store = useUserStore();
      await expect(store.getProfile()).rejects.toThrow('Profile fetch failed');
      expect(store.id).toBe(0);
      expect(store.username).toBe('');
      expect(store.email).toBe('');
    });

    it('should handle invalid token in localStorage', () => {
      localStorageMock.getItem.mockReturnValueOnce('invalid-token');
      const store = useUserStore();
      expect(store.token).toBe('invalid-token');
      expect(store.isLoggedIn).toBe(true);
      expect(store.id).toBe(0);
      expect(store.username).toBe('');
      expect(store.email).toBe('');
    });

    it('should handle missing token in localStorage', () => {
      localStorageMock.getItem.mockReturnValueOnce(null);
      const store = useUserStore();
      expect(store.token).toBeNull();
      expect(store.isLoggedIn).toBe(false);
    });
  });

  describe('Getters', () => {
    it('should return user info', () => {
      const store = useUserStore();
      store.setUserData({
        id: 1,
        username: 'testuser',
        email: 'test@example.com',
      });

      const userInfo = store.getUserInfo;
      expect(userInfo).toEqual({
        id: 1,
        username: 'testuser',
        email: 'test@example.com',
      });
    });
  });
});
