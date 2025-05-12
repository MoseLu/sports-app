import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest';
import { Capacitor } from '@capacitor/core';

// Mock Capacitor
vi.mock('@capacitor/core', () => ({
  Capacitor: {
    getPlatform: vi.fn().mockReturnValue('web'),
  },
}));

// Mock LiveUpdate 模块
vi.mock('@capawesome/capacitor-live-update', async () => {
  const original = await vi.importActual('@capawesome/capacitor-live-update');
  return {
    LiveUpdate: {
      sync: vi.fn().mockResolvedValue(undefined),
      ready: vi.fn().mockResolvedValue(undefined),
    },
  };
});

// Mock fetch
global.fetch = vi.fn();

// 保存原始的 console.error 和 setInterval
const originalConsoleError = console.error;
const originalSetInterval = global.setInterval;

describe('Live Update Boot', () => {
  beforeEach(() => {
    vi.clearAllMocks();
    // 模拟 console.error
    console.error = vi.fn();
    // 模拟 setInterval
    global.setInterval = vi.fn();
    // 模拟 fetch 成功返回
    global.fetch = vi.fn().mockResolvedValue({
      ok: true,
      json: vi.fn().mockResolvedValue({ version: '1.0.5' }),
    });
  });

  afterEach(() => {
    // 恢复原始的 console.error 和 setInterval
    console.error = originalConsoleError;
    global.setInterval = originalSetInterval;
  });

  it('should not perform update check on web platform', async () => {
    // 确保 Capacitor.getPlatform() 返回 'web'
    vi.mocked(Capacitor.getPlatform).mockReturnValue('web');

    // 导入 boot 函数
    const { default: liveupdateBoot } = await import('../live-update');

    // 执行 boot 函数
    await liveupdateBoot({});

    // 验证未调用 fetch（web 平台不执行更新检查）
    expect(global.fetch).not.toHaveBeenCalled();
    // 验证未设置定时器
    expect(global.setInterval).not.toHaveBeenCalled();
  });

  it('should perform update check on native platform', async () => {
    // 设置 Capacitor.getPlatform() 返回 'android'
    vi.mocked(Capacitor.getPlatform).mockReturnValue('android');

    // 导入 boot 函数
    const { default: liveupdateBoot } = await import('../live-update');

    // 执行 boot 函数
    await liveupdateBoot({});

    // 验证调用了 fetch
    expect(global.fetch).toHaveBeenCalled();
    expect(global.fetch).toHaveBeenCalledWith(expect.stringContaining('/api/manifest?version='));

    // 验证设置了定时器（检查周期性更新）
    expect(global.setInterval).toHaveBeenCalled();
    expect(global.setInterval).toHaveBeenCalledWith(expect.any(Function), 180000);
  });

  it('should handle errors during update check', async () => {
    // 设置 Capacitor.getPlatform() 返回 'ios'
    vi.mocked(Capacitor.getPlatform).mockReturnValue('ios');

    // 模拟 fetch 失败
    global.fetch = vi.fn().mockRejectedValue(new Error('Network error'));

    // 导入 boot 函数
    const { default: liveupdateBoot } = await import('../live-update');

    // 执行 boot 函数
    await liveupdateBoot({});

    // 验证错误被正确处理
    expect(console.error).toHaveBeenCalled();
    expect(console.error).toHaveBeenCalledWith('[LiveUpdate] failed:', expect.any(Error));
  });

  it('should handle API errors during update check', async () => {
    // 设置 Capacitor.getPlatform() 返回 'android'
    vi.mocked(Capacitor.getPlatform).mockReturnValue('android');

    // 模拟 API 返回错误
    global.fetch = vi.fn().mockResolvedValue({
      ok: false,
      status: 404,
      statusText: 'Not Found',
    });

    // 导入 boot 函数
    const { default: liveupdateBoot } = await import('../live-update');

    // 执行 boot 函数
    await liveupdateBoot({});

    // 验证错误被正确处理
    expect(console.error).toHaveBeenCalled();
  });
});
