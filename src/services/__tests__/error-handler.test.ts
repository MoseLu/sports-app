import { describe, it, expect, vi, beforeEach } from 'vitest';
import { handleError, showSuccess, showWarning } from '../error-handler';
import { Notify } from 'quasar';
import type { AxiosError } from 'axios';

// Mock Quasar Notify
vi.mock('quasar', () => ({
  Notify: {
    create: vi.fn(),
  },
}));

describe('Error Handler', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  describe('handleError', () => {
    it('should handle error with response data message', () => {
      const error = {
        response: {
          data: {
            message: '自定义错误消息',
          },
          status: 400,
        },
      } as AxiosError;

      handleError(error);

      expect(Notify.create).toHaveBeenCalledWith({
        type: 'negative',
        message: '自定义错误消息',
        position: 'top',
        timeout: 3000,
        classes: 'error-notify',
        actions: [
          {
            icon: 'close',
            color: 'white',
            round: true,
            dense: true,
          },
        ],
      });
    });

    it('should handle error with response error field', () => {
      const error = {
        response: {
          data: {
            error: '错误信息',
          },
          status: 400,
        },
      } as AxiosError;

      handleError(error);

      expect(Notify.create).toHaveBeenCalledWith({
        type: 'negative',
        message: '错误信息',
        position: 'top',
        timeout: 3000,
        classes: 'error-notify',
        actions: [
          {
            icon: 'close',
            color: 'white',
            round: true,
            dense: true,
          },
        ],
      });
    });

    it('should handle error with status code message', () => {
      const error = {
        response: {
          status: 401,
        },
      } as AxiosError;

      handleError(error);

      expect(Notify.create).toHaveBeenCalledWith({
        type: 'negative',
        message: '未授权，请重新登录',
        position: 'top',
        timeout: 3000,
        classes: 'error-notify',
        actions: [
          {
            icon: 'close',
            color: 'white',
            round: true,
            dense: true,
          },
        ],
      });
    });

    it('should handle error with error message', () => {
      const error = {
        message: '服务器内部错误',
      } as AxiosError;

      handleError(error);

      expect(Notify.create).toHaveBeenCalledWith({
        type: 'negative',
        message: '服务器内部错误',
        position: 'top',
        timeout: 3000,
        classes: 'error-notify',
        actions: [
          {
            icon: 'close',
            color: 'white',
            round: true,
            dense: true,
          },
        ],
      });
    });
  });

  describe('showSuccess', () => {
    it('should show success notification', () => {
      showSuccess('操作成功');

      expect(Notify.create).toHaveBeenCalledWith({
        type: 'positive',
        message: '操作成功',
        position: 'top',
        timeout: 2000,
        classes: 'success-notify',
        actions: [
          {
            icon: 'close',
            color: 'white',
            round: true,
            dense: true,
          },
        ],
      });
    });
  });

  describe('showWarning', () => {
    it('should show warning notification', () => {
      showWarning('警告信息');

      expect(Notify.create).toHaveBeenCalledWith({
        type: 'warning',
        message: '警告信息',
        position: 'top',
        timeout: 2000,
        classes: 'warning-notify',
        actions: [
          {
            icon: 'close',
            color: 'white',
            round: true,
            dense: true,
          },
        ],
      });
    });
  });
});
