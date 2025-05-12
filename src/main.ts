import { createApp } from 'vue';
import { Quasar, Notify, Dialog } from 'quasar';
import App from './App.vue';
import router from './router';
import { createPinia } from 'pinia';

// 导入 Quasar 样式
import '@quasar/extras/material-icons/material-icons.css';
import '@quasar/extras/fontawesome-v6/fontawesome-v6.css';
import 'quasar/dist/quasar.css';

// 导入全局样式
import './styles/table.scss';

// 创建应用实例
const app = createApp(App);

// 使用 Quasar
app.use(Quasar, {
  plugins: {
    Notify,
    Dialog,
  },
  config: {
    brand: {
      primary: '#1976D2',
      secondary: '#26A69A',
      accent: '#9C27B0',
      dark: '#1d1d1d',
      positive: '#21BA45',
      negative: '#C10015',
      info: '#31CCEC',
      warning: '#F2C037',
    },
  },
});

// 创建 Pinia 实例
const pinia = createPinia();

// 使用 Pinia 状态管理
app.use(pinia);

// 使用路由
app.use(router);

// 挂载应用
app.mount('#app');
