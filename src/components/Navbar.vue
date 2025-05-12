<template>
  <q-header elevated>
    <q-toolbar>
      <q-btn flat dense round icon="menu" aria-label="Menu" @click="$emit('toggle-drawer')" />

      <q-toolbar-title> Sports App </q-toolbar-title>

      <q-space />

      <q-btn-dropdown flat round dense icon="person">
        <q-list>
          <q-item clickable v-close-popup to="/login">
            <q-item-section avatar>
              <q-icon name="login" />
            </q-item-section>
            <q-item-section> 登录 </q-item-section>
          </q-item>

          <q-item clickable v-close-popup to="/register">
            <q-item-section avatar>
              <q-icon name="person_add" />
            </q-item-section>
            <q-item-section> 注册 </q-item-section>
          </q-item>

          <q-item clickable v-close-popup @click="handleLogout">
            <q-item-section avatar>
              <q-icon name="logout" />
            </q-item-section>
            <q-item-section> 退出 </q-item-section>
          </q-item>
        </q-list>
      </q-btn-dropdown>
    </q-toolbar>
  </q-header>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { useUserStore } from 'src/stores/user';
import { showSuccess, handleError } from '../services/error-handler';
import type { AxiosError } from 'axios';

defineOptions({
  name: 'AppNavbar',
});

const router = useRouter();
const userStore = useUserStore();

async function handleLogout() {
  try {
    await userStore.logout();
    showSuccess('已退出登录');
    await router.push('/login');
  } catch (error) {
    handleError(error as AxiosError);
  }
}

defineEmits(['toggle-drawer']);
</script>
