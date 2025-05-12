<template>
  <div class="login-page">
    <div class="login-container">
      <div class="login-content">
        <div class="login-header">
          <img :src="isDark ? darkLogoIcon : lightLogoIcon" alt="Re Sports" class="logo" />
          <h1 class="text-h4 q-mt-md">欢迎回来</h1>
          <p class="text-subtitle1 q-mt-sm text-grey-7">请登录您的账号</p>
        </div>

        <q-card class="login-card q-mt-lg">
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
              </q-input>

              <q-input
                v-model="form.password"
                label="密码"
                type="password"
                :rules="[(val) => !!val || '请输入密码']"
                outlined
                class="input-field"
              >
                <template v-slot:prepend>
                  <q-icon name="lock" />
                </template>
              </q-input>

              <div class="row justify-between items-center q-mt-md">
                <q-checkbox v-model="rememberMe" label="记住我" />
                <q-btn flat dense label="忘记密码？" to="/reset-password" class="text-primary" />
              </div>

              <q-btn
                label="登录"
                type="submit"
                color="primary"
                class="full-width q-mt-lg login-btn"
                :loading="loading"
              />

              <div class="text-center q-mt-md">
                <span class="text-grey-7">还没有账号？</span>
                <q-btn flat dense label="立即注册" to="/register" class="text-primary" />
              </div>
            </q-form>
          </q-card-section>
        </q-card>

        <div class="login-footer q-mt-xl">
          <p class="text-caption text-grey-7">© 2025 Re Sports. All rights reserved.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onUnmounted, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useQuasar } from 'quasar';
import { useUserStore } from 'stores/user';
import { showSuccess } from '../../services/error-handler';
import lightLogoIcon from '../../assets/icons/logo/logo-light.svg';
import darkLogoIcon from '../../assets/icons/logo/logo-dark.svg';

const router = useRouter();
const userStore = useUserStore();
const $q = useQuasar();

const isDark = computed(() => $q.dark.isActive);
const loading = computed(() => userStore.isLoading);
const rememberMe = ref(localStorage.getItem('rememberMe') === 'true');

interface LoginForm {
  username: string;
  password: string;
}

const form = ref<LoginForm>({
  username: '',
  password: '',
});

// 在组件挂载时读取保存的用户名
onMounted(() => {
  const rememberedUsername = localStorage.getItem('rememberedUsername');
  if (rememberedUsername) {
    form.value.username = rememberedUsername;
  }
});

// 在组件卸载时清除保存的用户名
onUnmounted(() => {
  // 只有在成功登录后才清除保存的用户名
  if (form.value.username) {
    localStorage.removeItem('rememberedUsername');
  }
});

const onSubmit = async () => {
  try {
    await userStore.login({
      username: form.value.username,
      password: form.value.password,
    });

    // 如果选择了"记住我"，保存用户名
    if (rememberMe.value) {
      localStorage.setItem('rememberMe', 'true');
      localStorage.setItem('rememberedUsername', form.value.username);
    } else {
      localStorage.removeItem('rememberMe');
      localStorage.removeItem('rememberedUsername');
    }

    showSuccess('登录成功');
    await router.push('/dashboard');
  } catch {
    // 如果登录失败，确保清除任何可能存在的无效 token
    localStorage.removeItem('token');
    // 错误处理已经在 userStore.login 中完成，这里不需要额外处理
  }
};
</script>

<style scoped>
.login-page {
  min-height: 100vh;
  background: linear-gradient(135deg, var(--q-primary) 0%, var(--q-secondary) 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
}

.login-container {
  width: 100%;
  max-width: 420px;
}

.login-content {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-radius: 16px;
  padding: 40px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.login-header {
  text-align: center;
}

.logo {
  width: 64px;
  height: 64px;
  margin-bottom: 16px;
}

.login-card {
  background: transparent;
  box-shadow: none;
}

.input-field {
  transition: all 0.3s ease;
}

.input-field:hover {
  transform: translateY(-2px);
}

.login-btn {
  height: 48px;
  font-size: 16px;
  font-weight: 500;
  border-radius: 8px;
  transition: all 0.3s ease;
}

.login-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.login-footer {
  text-align: center;
}

/* 深色模式适配 */
:deep(.q-dark) .login-content {
  background: rgba(30, 30, 30, 0.95);
}

:deep(.q-dark) .login-card {
  background: transparent;
}
</style>
