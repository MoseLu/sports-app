<template>
  <q-page padding>
    <div class="row q-col-gutter-md">
      <div class="col-12">
        <q-card>
          <q-card-section>
            <div class="text-h6">设置</div>
          </q-card-section>
          <q-card-section>
            <q-form @submit="onSubmit" class="q-gutter-md">
              <q-input
                v-model="form.notificationTime"
                label="提醒时间"
                type="time"
                :rules="[(val) => !!val || '请选择提醒时间']"
              />
              <q-toggle v-model="form.enableNotifications" label="启用提醒" />
              <div>
                <q-btn label="保存" type="submit" color="primary" />
              </div>
            </q-form>
          </q-card-section>
        </q-card>
      </div>
    </div>
  </q-page>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { api } from 'boot/axios';
import { showSuccess, handleError } from '../../services/error-handler';
import type { AxiosError } from 'axios';

export default defineComponent({
  name: 'SettingsPage',
  setup() {
    const form = ref({
      notificationTime: '19:00',
      enableNotifications: true,
    });

    const onSubmit = async () => {
      try {
        await api.put('/settings', form.value);
        showSuccess('设置已保存');
      } catch (error) {
        handleError(error as AxiosError);
      }
    };

    return {
      form,
      onSubmit,
    };
  },
});
</script>
