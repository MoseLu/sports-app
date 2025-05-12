<template>
  <q-page padding>
    <div class="row q-col-gutter-md">
      <div class="col-12">
        <q-card>
          <q-card-section>
            <div class="text-h6">运动统计</div>
          </q-card-section>
          <q-card-section>
            <div class="row q-col-gutter-md">
              <div class="col-12 col-md-6">
                <q-card>
                  <q-card-section>
                    <div class="text-subtitle2">本周运动时长</div>
                    <div class="text-h4">{{ weeklyDuration }} 分钟</div>
                  </q-card-section>
                </q-card>
              </div>
              <div class="col-12 col-md-6">
                <q-card>
                  <q-card-section>
                    <div class="text-subtitle2">本月运动次数</div>
                    <div class="text-h4">{{ monthlyCount }} 次</div>
                  </q-card-section>
                </q-card>
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>
    </div>
  </q-page>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import { api } from 'boot/axios';
import { handleError } from '../../services/error-handler';
import type { AxiosError } from 'axios';

export default defineComponent({
  name: 'StatisticsPage',
  setup() {
    const weeklyDuration = ref(0);
    const monthlyCount = ref(0);

    const loadStatistics = async () => {
      try {
        const response = await api.get('/statistics');
        weeklyDuration.value = response.data.weeklyDuration;
        monthlyCount.value = response.data.monthlyCount;
      } catch (error) {
        handleError(error as AxiosError);
      }
    };

    onMounted(() => {
      void loadStatistics();
    });

    return {
      weeklyDuration,
      monthlyCount,
    };
  },
});
</script>
