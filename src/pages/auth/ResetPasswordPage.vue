<template>
  <div class="reset-page">
    <div class="reset-container">
      <div class="reset-content">
        <div class="reset-header">
          <img :src="isDark ? darkLogoIcon : lightLogoIcon" alt="Re Sports" class="logo" />
          <h1 class="reset-title">重置密码</h1>
          <p class="text-subtitle1 q-mt-sm text-weight-medium text-center">
            请输入您的邮箱地址，我们将向您发送重置密码的链接
          </p>
        </div>

        <q-card class="reset-card q-mt-lg">
          <q-card-section>
            <q-form @submit="onSubmit" class="q-gutter-md">
              <q-input
                v-model="form.email"
                type="email"
                label="邮箱"
                :rules="[
                  (val) => !!val || '请输入邮箱',
                  (val) => /.+@.+\..+/.test(val) || '请输入有效的邮箱',
                ]"
                outlined
                class="input-field"
                @update:model-value="validateEmail"
                clearable
              >
                <template v-slot:prepend>
                  <q-icon name="email" />
                </template>
              </q-input>

              <q-input
                v-model="form.code"
                label="验证码"
                :rules="[(val) => !!val || '请输入验证码']"
                outlined
                class="input-field"
                clearable
              >
                <template v-slot:prepend>
                  <q-icon name="key" />
                </template>
                <template v-slot:append>
                  <q-btn
                    flat
                    dense
                    :label="countdown > 0 ? `${countdown}秒后重试` : '获取验证码'"
                    :disable="countdown > 0"
                    @click="sendCode"
                    class="text-primary"
                  />
                </template>
              </q-input>

              <q-input
                v-model="form.newPassword"
                type="password"
                label="新密码"
                :rules="[
                  (val) => !!val || '请输入新密码',
                  (val) => val.length >= 6 || '密码长度不能少于6位',
                ]"
                outlined
                class="input-field"
                clearable
              >
                <template v-slot:prepend>
                  <q-icon name="lock" />
                </template>
              </q-input>

              <q-input
                v-model="form.confirmPassword"
                type="password"
                label="确认密码"
                :rules="[
                  (val) => !!val || '请确认密码',
                  (val) => val === form.newPassword || '两次输入的密码不一致',
                ]"
                outlined
                class="input-field"
                clearable
              >
                <template v-slot:prepend>
                  <q-icon name="lock" />
                </template>
              </q-input>

              <div class="row justify-center">
                <q-btn
                  label="重置密码"
                  type="submit"
                  color="primary"
                  class="reset-btn"
                  :loading="loading"
                  unelevated
                  rounded
                />
              </div>

              <div class="text-center q-mt-md">
                <q-btn flat dense label="返回登录" to="/login" class="text-primary" />
              </div>
            </q-form>
          </q-card-section>
        </q-card>

        <div class="reset-footer q-mt-xl">
          <p class="text-caption text-weight-medium text-center">
            © 2025 Re Sports. All rights reserved.
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useQuasar } from 'quasar';
import { api } from 'boot/axios';
import { showSuccess, showWarning } from '../../services/error-handler';
import lightLogoIcon from '../../assets/icons/logo/logo-light.svg';
import darkLogoIcon from '../../assets/icons/logo/logo-dark.svg';

const router = useRouter();
const $q = useQuasar();
const isDark = computed(() => $q.dark.isActive);
const loading = ref(false);
const countdown = ref(0);

const form = ref({
  email: '',
  code: '',
  newPassword: '',
  confirmPassword: '',
});

const isValidEmail = ref(false);

const validateEmail = (value: string | number | null) => {
  if (typeof value === 'string') {
    const email = value.toLowerCase();
    if (!/.+@.+\..+/.test(email)) {
      isValidEmail.value = false;
      return;
    }
    // 检查是否是 Outlook 或 Microsoft 365 邮箱
    const isOutlook =
      email.endsWith('@outlook.com') ||
      email.endsWith('@hotmail.com') ||
      email.endsWith('@msn.com') ||
      email.endsWith('@bellis-technology.cn') ||
      email.includes('@microsoft') ||
      email.includes('@office365');
    const isQQ = email.endsWith('@qq.com');
    isValidEmail.value = isOutlook || isQQ;
  } else {
    isValidEmail.value = false;
  }
};

const startCountdown = () => {
  countdown.value = 30;
  const timer = setInterval(() => {
    countdown.value--;
    if (countdown.value <= 0) {
      clearInterval(timer);
    }
  }, 1000);
};

const sendCode = () => {
  if (!form.value.email) {
    showWarning('请输入邮箱地址');
    return;
  }

  if (!isValidEmail.value) {
    showWarning('请输入有效的邮箱地址（支持 QQ 邮箱、Outlook 邮箱和 Microsoft 365 邮箱）');
    return;
  }

  try {
    loading.value = true;
    showSuccess('验证码已发送到您的邮箱');
    startCountdown();

    api.post('/api/auth/send-reset-code', { email: form.value.email }).catch((error) => {
      console.error('发送验证码失败:', error);
      if (error && typeof error === 'object' && 'response' in error) {
        const axiosError = error as { response?: { data?: { error?: string } } };
        if (axiosError.response?.data?.error === '该邮箱未注册') {
          showWarning('该邮箱尚未注册，请先注册');
          localStorage.setItem('register_email', form.value.email);
          setTimeout(() => {
            void router.push('/register');
          }, 2000);
        }
      }
    });
  } finally {
    loading.value = false;
  }
};

const onSubmit = async () => {
  if (!form.value.code) {
    showWarning('请输入验证码');
    return;
  }

  if (!form.value.newPassword) {
    showWarning('请输入新密码');
    return;
  }

  if (form.value.newPassword !== form.value.confirmPassword) {
    showWarning('两次输入的密码不一致');
    return;
  }

  loading.value = true;
  try {
    const response = await api.post('/api/auth/reset-password', {
      email: form.value.email,
      code: form.value.code,
      newPassword: form.value.newPassword,
    });

    console.log('Reset password response:', response.data);

    if (response.status === 200 && response.data) {
      const username = response.data.data?.username || response.data.username;
      if (username) {
        console.log('Saving username:', username);
        localStorage.setItem('rememberedUsername', username);
      }
      showSuccess('密码重置成功');
      await router.push('/login');
    } else {
      showWarning('重置密码失败，请稍后重试');
    }
  } catch (error: unknown) {
    console.error('Reset password error:', error);
    if (error && typeof error === 'object' && 'response' in error) {
      const axiosError = error as { response?: { status?: number; data?: { error?: string } } };
      if (axiosError.response?.status === 400) {
        showWarning(axiosError.response.data?.error || '验证码错误或已过期');
      } else if (axiosError.response?.status === 500) {
        showWarning('服务器错误，请稍后重试');
      } else {
        showWarning('重置密码失败，请稍后重试');
      }
    } else {
      showWarning('重置密码失败，请稍后重试');
    }
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
.reset-page {
  min-height: 100vh;
  height: 100vh;
  background: linear-gradient(135deg, var(--q-primary) 0%, var(--q-secondary) 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
  overflow: hidden;
}

.reset-container {
  width: 100%;
  max-width: 480px;
  margin: auto;
}

.reset-content {
  text-align: center;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 16px;
  padding: 32px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.logo {
  width: 48px;
  height: 48px;
  margin-bottom: 16px;
}

.reset-card {
  background: transparent;
  box-shadow: none;
  padding: 0;
}

.input-field {
  margin-bottom: 12px;
  width: 100%;
}

.input-field :deep(.q-field__control) {
  height: 40px;
}

.input-field :deep(.q-field__native) {
  font-size: 14px;
  padding: 0 12px;
}

.input-field :deep(.q-field__label) {
  font-size: 14px;
  top: 10px;
}

.input-field :deep(.q-field__append) {
  height: 40px;
}

.input-field :deep(.q-field__prepend) {
  height: 40px;
  padding-left: 12px;
}

.reset-header h1 {
  color: var(--q-primary);
  margin-bottom: 12px;
  font-size: 1.75rem;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.reset-header p {
  color: var(--q-dark);
  opacity: 0.8;
  max-width: 320px;
  margin: 0 auto;
  font-size: 14px;
}

.reset-title {
  font-size: 2rem;
  font-weight: 700;
  color: var(--q-primary);
  margin: 16px 0;
  text-align: center;
  letter-spacing: 1px;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.reset-btn {
  min-width: 180px;
  height: 40px;
  font-size: 14px;
  font-weight: 500;
  letter-spacing: 1px;
  transition: all 0.3s ease;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0 24px;
  margin-top: 12px;
}

.reset-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.reset-btn:active {
  transform: translateY(0);
}

/* 深色模式适配 */
:deep(.q-dark) .reset-header h1 {
  color: var(--q-primary);
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

:deep(.q-dark) .reset-header p,
:deep(.q-dark) .reset-footer p {
  color: var(--q-light);
  opacity: 0.8;
}

:deep(.q-dark) .reset-btn {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
}

:deep(.q-dark) .reset-content {
  background: rgba(30, 30, 30, 0.95);
}

:deep(.q-dark) .reset-title {
  color: var(--q-primary);
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.reset-footer {
  margin-top: 24px;
}

.reset-footer p {
  font-size: 12px;
}
</style>
