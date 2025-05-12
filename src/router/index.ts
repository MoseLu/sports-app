import {
  createMemoryHistory,
  createRouter,
  createWebHashHistory,
  createWebHistory,
} from 'vue-router';
import routes from './routes';
import { useUserStore } from 'stores/user';

/*
 * If not building with SSR mode, you can
 * directly export the Router instantiation;
 *
 * The function below can be async too; either use
 * async/await or return a Promise which resolves
 * with the Router instance.
 */

const createHistory = process.env.SERVER
  ? createMemoryHistory
  : process.env.VUE_ROUTER_MODE === 'history'
    ? createWebHistory
    : createWebHashHistory;

const Router = createRouter({
  scrollBehavior: () => ({ left: 0, top: 0 }),
  routes,

  // Leave this as is and make changes in quasar.conf.js instead!
  // quasar.conf.js -> build -> vueRouterMode
  // quasar.conf.js -> build -> publicPath
  history: createHistory(process.env.VUE_ROUTER_BASE),
});

// 路由守卫
Router.beforeEach(async (to, from, next) => {
  const userStore = useUserStore();

  // 初始化用户状态
  if (!userStore.isAuthenticated) {
    await userStore.init();
  }

  const isAuthenticated = userStore.isAuthenticated;

  // 需要认证的路由
  if (to.meta.requiresAuth && !isAuthenticated) {
    next({ path: '/login', query: { redirect: to.fullPath } });
    return;
  }

  // 访客路由（已登录用户不能访问）
  if (to.meta.guest && isAuthenticated) {
    next('/');
    return;
  }

  // 记录页面访问历史
  if (from.name) {
    userStore.addToHistory({
      name: from.name as string,
      path: from.path,
      timestamp: new Date().toISOString(),
    });
  }

  next();
});

export default Router;
