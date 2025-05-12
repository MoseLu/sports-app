import { boot } from 'quasar/wrappers';
import { Capacitor } from '@capacitor/core';

// 获取当前版本号
function getCurrentVersion(): string {
  // 从 package.json 中获取版本号
  return '1.0.4'; // 这里应该从实际的配置中获取
}

// 检查更新的函数
async function checkForUpdates() {
  // 只在原生平台执行
  if (Capacitor.getPlatform() === 'web') {
    return;
  }

  try {
    // 动态导入 LiveUpdate
    const { LiveUpdate } = await import('@capawesome/capacitor-live-update');

    // 同步更新
    await LiveUpdate.sync({
      channel: 'production',
    });

    // 准备更新
    await LiveUpdate.ready();

    // 手动检查版本更新
    const response = await fetch(`/api/manifest?version=${getCurrentVersion()}`);
    if (!response.ok) {
      throw new Error('检查更新失败');
    }
  } catch (e) {
    console.error('[LiveUpdate] failed:', e);
  }
}

export default boot(async () => {
  // 只在原生平台执行
  if (Capacitor.getPlatform() === 'web') {
    return;
  }

  try {
    // 首次检查更新
    await checkForUpdates();

    // 设置定时器，每3分钟检查一次更新
    setInterval(() => {
      void checkForUpdates(); // 使用 void 操作符忽略 Promise 返回值
    }, 180000); // 180000ms = 3分钟
  } catch (e) {
    console.error('[LiveUpdate] initialization failed:', e);
  }
});
