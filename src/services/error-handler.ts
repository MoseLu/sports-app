import { Notify } from 'quasar';
import type { AxiosError } from 'axios';

interface ErrorResponse {
  message?: string;
  error?: string;
}

const errorMessages: Record<number, string> = {
  400: '请求参数错误',
  401: '未授权，请重新登录',
  403: '没有权限访问',
  404: '请求的资源不存在',
  500: '服务器内部错误',
};

export const handleError = (error: AxiosError) => {
  console.error('API Error:', error);

  // 显示错误消息
  const errorResponse = error.response?.data as ErrorResponse;
  const statusCode = error.response?.status || 500;
  const message =
    errorResponse?.message ||
    errorResponse?.error ||
    errorMessages[statusCode] ||
    error.message ||
    '发生错误';

  Notify.create({
    type: 'negative',
    message,
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
};

export const showSuccess = (message: string) => {
  Notify.create({
    type: 'positive',
    message,
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
};

export const showWarning = (message: string) => {
  Notify.create({
    type: 'warning',
    message,
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
};
