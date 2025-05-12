import { Capacitor } from '@capacitor/core';

interface DownloadProgress {
  receivedBytes: number;
  totalBytes: number;
}

interface CodePushPackage {
  appVersion: string;
  description: string;
  label: string;
  packageSize: number;
  failedInstall: boolean;
  installTime: string;
  isFirstRun: boolean;
  isMandatory: boolean;
  isPending: boolean;
}

interface CodePushUpdate {
  description: string;
  label: string;
  packageSize: number;
  isMandatory: boolean;
}

interface SyncOptions {
  installMode: number;
  mandatoryInstallMode: number;
  minimumBackgroundDuration: number;
  updateDialog: {
    updateTitle: string;
    mandatoryUpdateMessage: string;
    mandatoryContinueButtonLabel: string;
    optionalUpdateMessage: string;
    optionalIgnoreButtonLabel: string;
    optionalInstallButtonLabel: string;
  };
  syncStatus?: (status: number) => void;
  syncProgress?: (progress: DownloadProgress) => void;
}

declare global {
  interface Window {
    codePush?: {
      checkForUpdate(): Promise<CodePushUpdate | null>;
      sync(options: SyncOptions): Promise<number>;
      getCurrentPackage(): Promise<CodePushPackage | null>;
    };
  }
}

export class CodePushService {
  private static instance: CodePushService;
  private isChecking = false;

  private constructor() {}

  static getInstance(): CodePushService {
    if (!CodePushService.instance) {
      CodePushService.instance = new CodePushService();
    }
    return CodePushService.instance;
  }

  async checkForUpdate(): Promise<boolean> {
    if (!this.isNativePlatform() || this.isChecking) {
      return false;
    }

    this.isChecking = true;

    try {
      const update = await window.codePush?.checkForUpdate();
      this.isChecking = false;

      if (update) {
        console.log('发现新版本:', update);
        return true;
      } else {
        console.log('当前已是最新版本');
        return false;
      }
    } catch (error) {
      console.error('检查更新失败:', error);
      this.isChecking = false;
      return false;
    }
  }

  async syncUpdate(
    onStatusChange?: (status: number) => void,
    onProgress?: (progress: DownloadProgress) => void,
  ): Promise<boolean> {
    if (!this.isNativePlatform()) {
      return false;
    }

    try {
      const syncOptions: SyncOptions = {
        installMode: 1, // IMMEDIATE
        mandatoryInstallMode: 1,
        minimumBackgroundDuration: 0,
        updateDialog: {
          updateTitle: '发现新版本',
          mandatoryUpdateMessage: '必须更新才能继续使用',
          mandatoryContinueButtonLabel: '立即更新',
          optionalUpdateMessage: '有新版本可用',
          optionalIgnoreButtonLabel: '稍后再说',
          optionalInstallButtonLabel: '立即更新',
        },
        syncStatus: (status: number) => {
          console.log('更新状态:', status);
          onStatusChange?.(status);
        },
        syncProgress: (progress: DownloadProgress) => {
          console.log('下载进度:', progress);
          onProgress?.(progress);
        },
      };

      const result = await window.codePush?.sync(syncOptions);
      return result === 0; // 0 = UP_TO_DATE
    } catch (error) {
      console.error('同步更新失败:', error);
      return false;
    }
  }

  async getCurrentPackage(): Promise<CodePushPackage | null> {
    if (!this.isNativePlatform()) {
      return null;
    }

    try {
      const pkg = await window.codePush?.getCurrentPackage();
      console.log('当前包信息:', pkg);
      return pkg ?? null;
    } catch (error) {
      console.error('获取当前包信息失败:', error);
      return null;
    }
  }

  private isNativePlatform(): boolean {
    return Capacitor.isNativePlatform();
  }
}
