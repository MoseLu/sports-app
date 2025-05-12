import { describe, it, expect, vi, beforeEach } from 'vitest';
import { CodePushService } from '../code-push';
import { Capacitor } from '@capacitor/core';

// Mock Capacitor
vi.mock('@capacitor/core', () => ({
  Capacitor: {
    isNativePlatform: vi.fn(),
  },
}));

// Mock window.codePush
const mockCodePush = {
  checkForUpdate: vi.fn(),
  sync: vi.fn(),
  getCurrentPackage: vi.fn(),
};

describe('CodePush Service', () => {
  beforeEach(() => {
    vi.clearAllMocks();
    // Reset window.codePush before each test
    window.codePush = undefined;
  });

  describe('getInstance', () => {
    it('should return the same instance', () => {
      const instance1 = CodePushService.getInstance();
      const instance2 = CodePushService.getInstance();
      expect(instance1).toBe(instance2);
    });
  });

  describe('checkForUpdate', () => {
    it('should return false on non-native platform', async () => {
      vi.mocked(Capacitor.isNativePlatform).mockReturnValue(false);
      const service = CodePushService.getInstance();
      const result = await service.checkForUpdate();
      expect(result).toBe(false);
    });

    it('should return false when no update is available', async () => {
      vi.mocked(Capacitor.isNativePlatform).mockReturnValue(true);
      window.codePush = mockCodePush;
      mockCodePush.checkForUpdate.mockResolvedValueOnce(null);

      const service = CodePushService.getInstance();
      const result = await service.checkForUpdate();

      expect(mockCodePush.checkForUpdate).toHaveBeenCalled();
      expect(result).toBe(false);
    });

    it('should return true when update is available', async () => {
      vi.mocked(Capacitor.isNativePlatform).mockReturnValue(true);
      window.codePush = mockCodePush;
      mockCodePush.checkForUpdate.mockResolvedValueOnce({
        description: '新版本',
        label: 'v1.0.1',
        packageSize: 1000,
        isMandatory: false,
      });

      const service = CodePushService.getInstance();
      const result = await service.checkForUpdate();

      expect(mockCodePush.checkForUpdate).toHaveBeenCalled();
      expect(result).toBe(true);
    });

    it('should handle errors gracefully', async () => {
      vi.mocked(Capacitor.isNativePlatform).mockReturnValue(true);
      window.codePush = mockCodePush;
      mockCodePush.checkForUpdate.mockRejectedValueOnce(new Error('Network error'));

      const service = CodePushService.getInstance();
      const result = await service.checkForUpdate();

      expect(mockCodePush.checkForUpdate).toHaveBeenCalled();
      expect(result).toBe(false);
    });
  });

  describe('syncUpdate', () => {
    it('should return false on non-native platform', async () => {
      vi.mocked(Capacitor.isNativePlatform).mockReturnValue(false);
      const service = CodePushService.getInstance();
      const result = await service.syncUpdate();
      expect(result).toBe(false);
    });

    it('should sync update successfully', async () => {
      vi.mocked(Capacitor.isNativePlatform).mockReturnValue(true);
      window.codePush = mockCodePush;
      mockCodePush.sync.mockResolvedValueOnce(0);

      const service = CodePushService.getInstance();
      const result = await service.syncUpdate();

      expect(mockCodePush.sync).toHaveBeenCalledWith({
        installMode: 1,
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
        syncStatus: expect.any(Function),
        syncProgress: expect.any(Function),
      });
      expect(result).toBe(true);
    });

    it('should handle sync errors gracefully', async () => {
      vi.mocked(Capacitor.isNativePlatform).mockReturnValue(true);
      window.codePush = mockCodePush;
      mockCodePush.sync.mockRejectedValueOnce(new Error('Sync error'));

      const service = CodePushService.getInstance();
      const result = await service.syncUpdate();

      expect(mockCodePush.sync).toHaveBeenCalled();
      expect(result).toBe(false);
    });

    it('should call status and progress callbacks', async () => {
      vi.mocked(Capacitor.isNativePlatform).mockReturnValue(true);
      window.codePush = mockCodePush;
      mockCodePush.sync.mockImplementationOnce(async (options) => {
        options.syncStatus?.(1);
        options.syncProgress?.({ receivedBytes: 50, totalBytes: 100 });
        return 0;
      });

      const statusCallback = vi.fn();
      const progressCallback = vi.fn();

      const service = CodePushService.getInstance();
      await service.syncUpdate(statusCallback, progressCallback);

      expect(statusCallback).toHaveBeenCalledWith(1);
      expect(progressCallback).toHaveBeenCalledWith({
        receivedBytes: 50,
        totalBytes: 100,
      });
    });
  });

  describe('getCurrentPackage', () => {
    it('should return null on non-native platform', async () => {
      vi.mocked(Capacitor.isNativePlatform).mockReturnValue(false);
      const service = CodePushService.getInstance();
      const result = await service.getCurrentPackage();
      expect(result).toBeNull();
    });

    it('should return current package info', async () => {
      vi.mocked(Capacitor.isNativePlatform).mockReturnValue(true);
      window.codePush = mockCodePush;
      const mockPackage = {
        appVersion: '1.0.0',
        description: '当前版本',
        label: 'v1.0.0',
        packageSize: 1000,
        failedInstall: false,
        installTime: '2024-03-20',
        isFirstRun: false,
        isMandatory: false,
        isPending: false,
      };
      mockCodePush.getCurrentPackage.mockResolvedValueOnce(mockPackage);

      const service = CodePushService.getInstance();
      const result = await service.getCurrentPackage();

      expect(mockCodePush.getCurrentPackage).toHaveBeenCalled();
      expect(result).toEqual(mockPackage);
    });

    it('should handle errors gracefully', async () => {
      vi.mocked(Capacitor.isNativePlatform).mockReturnValue(true);
      window.codePush = mockCodePush;
      mockCodePush.getCurrentPackage.mockRejectedValueOnce(new Error('Network error'));

      const service = CodePushService.getInstance();
      const result = await service.getCurrentPackage();

      expect(mockCodePush.getCurrentPackage).toHaveBeenCalled();
      expect(result).toBeNull();
    });
  });
});
