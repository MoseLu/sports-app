<template>
  <q-layout view="hHh lpR fFf">
    <!-- —— 顶部 Header —— -->
    <q-header elevated class="bg-primary text-white app-header">
      <q-toolbar>
        <q-toolbar-title>
          <img
            :src="isDark ? darkLogoIcon : lightLogoIcon"
            alt="Re Sports"
            class="logo-icon"
          />
        </q-toolbar-title>
        <q-btn dense flat round @click="toggleLanguage">
          <q-icon><img :src="languageIcon" /></q-icon>
        </q-btn>
        <q-btn dense flat round @click="toggleDarkMode">
          <q-icon>
            <img :src="isDark ? lightModeIcon : darkModeIcon" />
          </q-icon>
        </q-btn>
        <q-btn dense flat round @click="logout">
          <q-icon><img :src="logoutIcon" /></q-icon>
        </q-btn>
      </q-toolbar>
    </q-header>

    <!-- —— 页面内容 —— -->
    <q-page-container>
      <router-view />
    </q-page-container>

    <!-- —— 底部导航 —— -->
    <q-footer elevated class="bg-white app-footer">
      <q-tabs
        v-model="activeTab"
        dense
        background-color="white"
        text-color="grey-6"
        active-color="primary"
        indicator-color="transparent"
      >
        <q-route-tab
          name="dashboard"
          to="/dashboard"
          class="q-pa-sm"
        >
          <q-icon>
            <img
              :src="activeTab === 'dashboard'
                ? dashboardActiveIcon
                : dashboardInactiveIcon"
            />
          </q-icon>
          <div class="q-tab__label">统计</div>
        </q-route-tab>

        <q-route-tab
          name="community"
          to="/community"
          class="q-pa-sm"
        >
          <q-icon>
            <img
              :src="activeTab === 'community'
                ? communityActiveIcon
                : communityInactiveIcon"
            />
          </q-icon>
          <div class="q-tab__label">社区</div>
        </q-route-tab>

        <q-route-tab
          name="profile"
          to="/profile"
          class="q-pa-sm"
        >
          <q-icon>
            <img
              :src="activeTab === 'profile'
                ? profileActiveIcon
                : profileInactiveIcon"
            />
          </q-icon>
          <div class="q-tab__label">个人资料</div>
        </q-route-tab>
      </q-tabs>
    </q-footer>
  </q-layout>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useQuasar } from 'quasar';
import { useUserStore } from 'stores/user';

import languageIcon from '../assets/icons/common/language.svg';
import lightModeIcon from '../assets/icons/common/light_mode.svg';
import darkModeIcon from '../assets/icons/common/dark_mode.svg';
import logoutIcon from '../assets/icons/common/logout.svg';
import dashboardActiveIcon from '../assets/icons/tabbar/dashboard-active.svg';
import dashboardInactiveIcon from '../assets/icons/tabbar/dashboard-inactive.svg';
import profileActiveIcon from '../assets/icons/tabbar/profile-active.svg';
import profileInactiveIcon from '../assets/icons/tabbar/profile-inactive.svg';
import communityActiveIcon from '../assets/icons/tabbar/community-active.svg';
import communityInactiveIcon from '../assets/icons/tabbar/community-inactive.svg';
import lightLogoIcon from '../assets/icons/logo/logo-light.svg';
import darkLogoIcon from '../assets/icons/logo/logo-dark.svg';

const $q = useQuasar();
const router = useRouter();
const route = useRoute();
const userStore = useUserStore();

const isDark = ref(false);
const currentLocale = ref('zh-CN');
const activeTab = ref('dashboard');

const toggleLanguage = () => {
  const newLocale = currentLocale.value === 'zh-CN' ? 'en-US' : 'zh-CN';
  currentLocale.value = newLocale;
  localStorage.setItem('locale', newLocale);
};

const toggleDarkMode = () => {
  isDark.value = !isDark.value;
  $q.dark.set(isDark.value);
};

const logout = async () => {
  try {
    await userStore.logout();
    await router.push('/login');
  } catch (err) {
    console.error('Logout failed:', err);
  }
};

onMounted(() => {
  isDark.value = $q.dark.isActive;
  currentLocale.value = localStorage.getItem('locale') || 'zh-CN';

  const p = route.path;
  if (p.includes('dashboard')) activeTab.value = 'dashboard';
  else if (p.includes('community')) activeTab.value = 'community';
  else if (p.includes('profile')) activeTab.value = 'profile';
});
</script>

<style scoped>
/* —— 全局 reset —— */
html,
body,
#q-app {
  height: 100%;
  margin: 0;
}

/* —— 顶部 Header 安全区 + fallback —— */
.app-header {
  position: fixed;
  top: 0; left: 0; right: 0;
  height: calc(56px + env(safe-area-inset-top, 24px));
  padding-top: env(safe-area-inset-top, 24px);
}

/* —— 底部 Footer 安全区 + fallback —— */
.app-footer {
  position: fixed;
  left: 0; right: 0; bottom: 0;
  padding-bottom: env(safe-area-inset-bottom, 56px);
}

/* —— Tab 样式 —— */
.q-tab {
  min-height: 50px;
  flex-direction: column;
  padding: 8px 0;
}

.q-tab__label {
  font-size: 12px;
  margin-top: 4px;
}

.q-icon img {
  width: 24px;
  height: 24px;
}

/* 在 MainLayout.vue <style scoped> 里 */

.app-footer .q-icon img {
  /* 例 ①：中灰，大约 #777 */
  filter: invert(60%);

  /* 以下几个可任选写一组，或者自己 mix 调：
  filter: invert(60%) brightness(0.9);
  filter: invert(50%) brightness(0.8);
  filter: invert(70%) brightness(1.2);
  filter: invert(60%) sepia(10%) brightness(0.9);
  */
}

/* —— 如果想让“激活态”图标用主色（#027be3 举例） —— */
.app-footer .q-tab--active .q-icon img {
  /* 先反色到灰，再给一点主色调的 sepia+hue-rotate */
  filter:
    invert(50%)
    sepia(80%)
    saturate(300%)
    hue-rotate(200deg)
    brightness(1);
}


/* —— Logo 图标 —— */
.logo-icon {
  width: 32px;
  height: 32px;
  vertical-align: middle;
}
</style>
