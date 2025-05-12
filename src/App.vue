<template>
  <router-view />
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useQuasar } from 'quasar';
import { CodePushService } from './services/code-push';

const $q = useQuasar();
const codePush = CodePushService.getInstance();

const loading = ref<{ show: (message?: string) => void; hide: () => void }>({
  show: (message?: string) => {
    $q.loading.show({
      message: message || '正在更新...',
      spinnerColor: 'primary',
    });
  },
  hide: () => {
    $q.loading.hide();
  },
});

onMounted(() => {
  // 检查更新
  void checkForUpdates();
});

async function checkForUpdates() {
  const hasUpdate = await codePush.checkForUpdate();

  if (hasUpdate) {
    $q.dialog({
      title: '发现新版本',
      message: '是否立即更新？',
      cancel: true,
      persistent: true,
    }).onOk(() => {
      // 在更新过程中显示进度
      loading.value.show('正在下载更新...');

      // 同步更新
      void codePush
        .syncUpdate(
          (status) => {
            // 更新状态变化
            if (status === 0) {
              // UP_TO_DATE
              loading.value.hide();
              $q.notify({
                type: 'positive',
                message: '更新完成，请重启应用',
              });
            }
          },
          (progress) => {
            // 更新下载进度
            loading.value.show(
              `正在更新... ${Math.round((progress.receivedBytes / progress.totalBytes) * 100)}%`,
            );
          },
        )
        .catch(() => {
          loading.value.hide();
          $q.notify({
            type: 'negative',
            message: '更新失败，请稍后重试',
          });
        });
    });
  }
}
</script>

<style lang="scss">
html,
body,
#q-app {
  height: 100%;
  margin: 0;
}

/* 顶部导航 */
.q-header {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  /* header 内容区高 56px + 顶部安全区（iOS 会取系统值，Android fallback 24px） */
  height: calc(56px + env(safe-area-inset-top, 24px));
  padding-top: env(safe-area-inset-top, 24px);
  background: white;
  z-index: 3000;
}

/* 底部导航 */
.q-footer {
  position: fixed;
  left: 0;
  right: 0;
  bottom: 0;
  /* 图标区高 96px + 底部安全区（iOS 会取系统值，Android fallback 56px） */
  height: calc(96px + env(safe-area-inset-bottom, 56px));
  padding-bottom: env(safe-area-inset-bottom, 56px);
  background: white;
  z-index: 2000;

  .q-tab {
    padding: 8px 0;

    &__icon {
      font-size: 24px;
      margin-bottom: 4px;
    }

    &__label {
      font-size: 12px;
      line-height: 1.2;
    }

    &--active {
      color: var(--q-primary);
    }

    &:not(.q-tab--active) {
      color: #666;
    }
  }
}

/* 主内容区，顶底都要预留空间 */
.q-page {
  padding-top: calc(56px + env(safe-area-inset-top, 24px));
  padding-bottom: calc(96px + env(safe-area-inset-bottom, 56px));
}
</style>
