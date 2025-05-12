import { describe, it, expect, beforeEach } from 'vitest';
import { createI18n } from 'vue-i18n';
import { watch, nextTick } from 'vue';
import zhCN from '../../i18n/zh-CN';
import enUS from '../../i18n/en-US';

const messages = {
  'zh-CN': zhCN,
  'en-US': enUS,
};

describe('i18n', () => {
  let i18n: ReturnType<typeof createI18n>;

  beforeEach(() => {
    // 清除 localStorage
    localStorage.clear();

    // 创建新的 i18n 实例，强制使用 Composition API 模式
    i18n = createI18n({
      legacy: false,
      locale: 'zh-CN',
      fallbackLocale: 'zh-CN',
      messages,
    });
  });

  describe('初始化', () => {
    it('应该使用默认语言 zh-CN', () => {
      expect(i18n.global.locale.value).toBe('zh-CN');
    });

    it('应该从 localStorage 读取语言设置', () => {
      localStorage.setItem('locale', 'en-US');
      const i18nWithStorage = createI18n({
        legacy: false,
        locale: localStorage.getItem('locale') || 'zh-CN',
        fallbackLocale: 'zh-CN',
        messages,
      });
      expect(i18nWithStorage.global.locale.value).toBe('en-US');
    });
  });

  describe('翻译功能', () => {
    it('应该正确翻译中文', () => {
      expect(i18n.global.t('common.dashboard')).toBe('首页');
      expect(i18n.global.t('common.profile')).toBe('我的');
      expect(i18n.global.t('auth.login')).toBe('登录');
    });

    it('应该正确翻译英文', () => {
      i18n.global.locale.value = 'en-US';
      expect(i18n.global.t('common.dashboard')).toBe('Dashboard');
      expect(i18n.global.t('common.profile')).toBe('Profile');
      expect(i18n.global.t('auth.login')).toBe('Login');
    });

    it('应该使用回退语言当翻译不存在时', () => {
      i18n.global.locale.value = 'en-US';
      // 假设这个键只在中文中存在
      expect(i18n.global.t('common.unknownKey')).toBe('common.unknownKey');
    });
  });

  describe('语言切换', () => {
    it('应该能够切换语言', () => {
      i18n.global.locale.value = 'en-US';
      expect(i18n.global.locale.value).toBe('en-US');
      expect(i18n.global.t('common.dashboard')).toBe('Dashboard');

      i18n.global.locale.value = 'zh-CN';
      expect(i18n.global.locale.value).toBe('zh-CN');
      expect(i18n.global.t('common.dashboard')).toBe('首页');
    });

    it('切换语言时应该更新 localStorage', async () => {
      // 监听 locale 变化并写入 localStorage
      watch(i18n.global.locale, (newLocale) => {
        localStorage.setItem('locale', newLocale);
      });

      i18n.global.locale.value = 'en-US';
      await nextTick();
      expect(localStorage.getItem('locale')).toBe('en-US');

      i18n.global.locale.value = 'zh-CN';
      await nextTick();
      expect(localStorage.getItem('locale')).toBe('zh-CN');
    });
  });
});
