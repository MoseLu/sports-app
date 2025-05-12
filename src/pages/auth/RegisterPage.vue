<template>
  <div class="register-page">
    <div class="register-container">
      <div class="register-content">
        <div class="register-header">
          <img :src="isDark ? darkLogoIcon : lightLogoIcon" alt="Re Sports" class="logo" />
          <h1 class="text-h4 q-mt-md">创建账号</h1>
          <p class="text-subtitle1 q-mt-sm text-grey-7">加入 Re Sports 社区</p>
        </div>

        <q-card class="register-card q-mt-lg">
          <q-card-section>
            <q-form @submit="onSubmit" class="q-gutter-md">
              <q-input
                v-model="form.username"
                label="用户名"
                :rules="[(val) => !!val || '请输入用户名']"
                outlined
                class="input-field"
              >
                <template v-slot:prepend>
                  <q-icon name="person" />
                </template>
                <template v-slot:hint>
                  <span class="text-caption">用户名将用于登录和显示</span>
                </template>
              </q-input>

              <q-input
                v-model="form.email"
                type="email"
                label="邮箱"
                :rules="[
                  (val) => !!val || '请输入邮箱',
                  (val) => /.+@.+\..+/.test(val) || '请输入有效的邮箱',
                  (val) =>
                    /@(qq\.com|outlook\.com|hotmail\.com|msn\.com|bellis-technology\.cn|microsoft\.|office365\.)/.test(
                      val,
                    ) || '请输入有效的邮箱地址（支持 QQ 邮箱、Outlook 邮箱和 Microsoft 365 邮箱）',
                ]"
                outlined
                class="input-field"
              >
                <template v-slot:prepend>
                  <q-icon name="email" />
                </template>
                <template v-slot:hint>
                  <span class="text-caption">用于接收通知和找回密码</span>
                </template>
              </q-input>

              <q-input
                v-model="form.password"
                type="password"
                label="密码"
                :rules="[
                  (val) => !!val || '请输入密码',
                  (val) => val.length >= 6 || '密码长度至少为6位',
                ]"
                outlined
                class="input-field"
              >
                <template v-slot:prepend>
                  <q-icon name="lock" />
                </template>
                <template v-slot:hint>
                  <span class="text-caption">密码长度至少6位，建议使用字母、数字和符号组合</span>
                </template>
              </q-input>

              <q-input
                v-model="form.confirmPassword"
                type="password"
                label="确认密码"
                :rules="[
                  (val) => !!val || '请确认密码',
                  (val) => val === form.password || '两次输入的密码不一致',
                ]"
                outlined
                class="input-field"
              >
                <template v-slot:prepend>
                  <q-icon name="lock" />
                </template>
              </q-input>

              <div class="row justify-between items-center q-mt-md">
                <q-checkbox v-model="agreeTerms" label="我已阅读并同意" />
                <q-btn flat dense label="隐私协议" class="text-primary" />
                <q-btn flat dense label="服务条款" class="text-primary" />
              </div>

              <q-btn
                type="submit"
                color="primary"
                label="注册"
                class="full-width q-mt-lg register-btn"
                :loading="loading"
                :disable="!agreeTerms"
              />

              <div class="text-center q-mt-md">
                <span class="text-grey-7">已有账号？</span>
                <q-btn flat dense label="立即登录" to="/login" class="text-primary" />
              </div>
            </q-form>
          </q-card-section>
        </q-card>

        <div class="register-footer q-mt-xl">
          <p class="text-caption text-grey-7">© 2025 Re Sports. All rights reserved.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onUnmounted } from 'vue';
import { useRouter } from 'vue-router';
import { useQuasar } from 'quasar';
import { useUserStore } from 'stores/user';
import { showSuccess, showWarning, handleError } from '../../services/error-handler';
import type { AxiosError } from 'axios';
import lightLogoIcon from '../../assets/icons/logo/logo-light.svg';
import darkLogoIcon from '../../assets/icons/logo/logo-dark.svg';

const router = useRouter();
const userStore = useUserStore();
const $q = useQuasar();

const isDark = computed(() => $q.dark.isActive);
const loading = computed(() => userStore.isLoading);
const agreeTerms = ref(false);

const form = ref({
  email: localStorage.getItem('register_email') || '',
  username: '',
  password: '',
  confirmPassword: '',
});

// 在组件卸载时清除存储的邮箱
onUnmounted(() => {
  localStorage.removeItem('register_email');
});

async function onSubmit() {
  if (form.value.password !== form.value.confirmPassword) {
    showWarning('两次输入的密码不一致');
    return;
  }

  if (!agreeTerms.value) {
    showWarning('请阅读并同意服务条款');
    return;
  }

  try {
    await userStore.register({
      username: form.value.username,
      email: form.value.email,
      password: form.value.password,
    });
    showSuccess('注册成功');
    await router.push('/login');
  } catch (error) {
    handleError(error as AxiosError);
  }
}
</script>

<style scoped>
.register-page {
  height: 100vh;
  background: linear-gradient(135deg, var(--q-primary) 0%, var(--q-secondary) 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 16px;
}

.register-container {
  width: 100%;
  max-width: 420px;
  max-height: calc(100vh - 32px);
}

.register-content {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 16px;
  padding: 24px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.register-header {
  text-align: center;
  margin-bottom: 16px;
}

.logo {
  width: 40px;
  height: 40px;
  margin-bottom: 8px;
}

.register-card {
  background: transparent;
  box-shadow: none;
}

.input-field {
  transition: all 0.3s ease;
}

.input-field:hover {
  transform: translateY(-2px);
}

.register-btn {
  height: 40px;
  font-size: 16px;
  font-weight: 500;
  border-radius: 8px;
  transition: all 0.3s ease;
}

.register-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.register-footer {
  text-align: center;
  margin-top: 16px;
}

/* 深色模式适配 */
:deep(.q-dark) .register-content {
  background: rgba(30, 30, 30, 0.95);
}

:deep(.q-dark) .register-card {
  background: transparent;
}

/* 移动端适配 */
@media (max-height: 600px) {
  .register-content {
    padding: 16px;
  }

  .logo {
    width: 32px;
    height: 32px;
    margin-bottom: 4px;
  }

  .register-header {
    margin-bottom: 12px;
  }

  .register-footer {
    margin-top: 12px;
  }

  .q-gutter-md {
    --q-gutter-md: 12px;
  }
}
</style>
