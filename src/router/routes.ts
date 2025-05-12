import type { RouteRecordRaw } from 'vue-router';
import MainLayout from 'layouts/MainLayout.vue';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: MainLayout,
    meta: { requiresAuth: true },
    children: [
      {
        path: '',
        redirect: '/dashboard',
      },
      {
        path: 'dashboard',
        component: () => import('pages/exercise/DashboardPage.vue'),
        meta: { requiresAuth: true },
      },
      {
        path: 'profile',
        component: () => import('pages/user/ProfilePage.vue'),
        meta: { requiresAuth: true },
      },
      {
        path: 'community',
        component: () => import('pages/community/CommunityPage.vue'),
        meta: { requiresAuth: true },
      },
    ],
  },
  {
    path: '/login',
    component: () => import('pages/auth/LoginPage.vue'),
    meta: { guest: true },
  },
  {
    path: '/reset-password',
    component: () => import('pages/auth/ResetPasswordPage.vue'),
    meta: { guest: true },
  },
  {
    path: '/register',
    component: () => import('pages/auth/RegisterPage.vue'),
    meta: { guest: true },
  },
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/error/ErrorNotFound.vue'),
  },
];

export default routes;
