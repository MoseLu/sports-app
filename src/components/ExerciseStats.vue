<template>
  <div class="q-pa-sm">
    <div class="row q-col-gutter-sm">
      <div class="col-6">
        <q-card class="stats-card">
          <q-card-section class="row items-center no-wrap">
            <div class="col">
              <div class="text-subtitle2 text-grey-7">本周运动时长</div>
              <div class="text-h5 text-weight-medium">{{ stats.total_duration }}</div>
              <div class="text-caption text-grey-7">分钟</div>
            </div>
            <q-icon name="schedule" size="2rem" color="primary" />
          </q-card-section>
        </q-card>
      </div>

      <div class="col-6">
        <q-card class="stats-card">
          <q-card-section class="row items-center no-wrap">
            <div class="col">
              <div class="text-subtitle2 text-grey-7">本月运动次数</div>
              <div class="text-h5 text-weight-medium">{{ stats.exercise_count }}</div>
              <div class="text-caption text-grey-7">次</div>
            </div>
            <q-icon name="fitness_center" size="2rem" color="secondary" />
          </q-card-section>
        </q-card>
      </div>

      <div class="col-6">
        <q-card class="stats-card">
          <q-card-section class="row items-center no-wrap">
            <div class="col">
              <div class="text-subtitle2 text-grey-7">平均运动时长</div>
              <div class="text-h5 text-weight-medium">{{ stats.average_duration }}</div>
              <div class="text-caption text-grey-7">分钟/次</div>
            </div>
            <q-icon name="timer" size="2rem" color="accent" />
          </q-card-section>
        </q-card>
      </div>

      <div class="col-6">
        <q-card class="stats-card">
          <q-card-section class="row items-center no-wrap">
            <div class="col">
              <div class="text-subtitle2 text-grey-7">平均消耗卡路里</div>
              <div class="text-h5 text-weight-medium">{{ stats.average_calories }}</div>
              <div class="text-caption text-grey-7">卡路里/次</div>
            </div>
            <q-icon name="local_fire_department" size="2rem" color="positive" />
          </q-card-section>
        </q-card>
      </div>
    </div>

    <div class="row q-mt-sm">
      <q-card class="col-12 stats-card">
        <q-card-section class="q-pa-sm">
          <div class="row justify-between items-center q-mb-xs">
            <div class="text-subtitle2 text-grey-8">近7天运动统计</div>
            <div class="row q-gutter-sm">
              <q-select
                v-model="timeRange"
                :options="timeRangeOptions"
                dense
                outlined
                emit-value
                map-options
                class="filter-select"
                label="时间范围"
                dark
              />
              <q-select
                v-model="exerciseType"
                :options="exerciseTypeOptions"
                dense
                outlined
                emit-value
                map-options
                class="filter-select"
                label="运动类型"
                dark
              />
            </div>
          </div>
          <div ref="chartRef" style="height: 300px"></div>
        </q-card-section>
      </q-card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, onUnmounted, watch } from 'vue';
import { useExerciseStore } from 'stores/exercise-store';
import * as echarts from 'echarts';
import { handleError } from '../services/error-handler';
import type { AxiosError } from 'axios';

const exerciseStore = useExerciseStore();
const stats = computed(() => {
  const storeStats = exerciseStore.getStats;
  return (
    storeStats || {
      total_duration: 0,
      total_calories: 0,
      exercise_count: 0,
      average_duration: 0,
      average_calories: 0,
      daily_duration: [],
      daily_count: [],
    }
  );
});
const chartRef = ref<HTMLElement | null>(null);
let chart: echarts.ECharts | null = null;

const timeRange = ref('week');
const exerciseType = ref('all');

const timeRangeOptions = [
  { label: '本周', value: 'week' },
  { label: '本月', value: 'month' },
  { label: '本年', value: 'year' },
];

const exerciseTypeOptions = [
  { label: '全部运动', value: 'all' },
  { label: '跑步', value: '跑步' },
  { label: '骑行', value: '骑行' },
  { label: '游泳', value: '游泳' },
  { label: '健身', value: '健身' },
];

const initChart = () => {
  if (chartRef.value) {
    chart = echarts.init(chartRef.value);
    updateChart();
  }
};

const updateChart = () => {
  if (!chart || !stats.value) return;

  // 获取最近7天的日期
  const dates = Array.from({ length: 7 }, (_, i) => {
    const date = new Date();
    date.setDate(date.getDate() - (6 - i));
    return date.toLocaleDateString('zh-CN', { month: 'numeric', day: 'numeric' });
  });

  const durations = stats.value.daily_duration || [];
  const counts = stats.value.daily_count || [];

  chart.setOption({
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'cross',
        crossStyle: {
          color: '#999',
        },
      },
    },
    legend: {
      data: ['运动时长', '运动次数'],
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true,
    },
    xAxis: {
      type: 'category',
      data: dates,
      axisPointer: {
        type: 'shadow',
      },
    },
    yAxis: [
      {
        type: 'value',
        name: '时长(分钟)',
        min: 0,
        axisLabel: {
          formatter: '{value}',
        },
      },
      {
        type: 'value',
        name: '次数',
        min: 0,
        axisLabel: {
          formatter: '{value}',
        },
      },
    ],
    series: [
      {
        name: '运动时长',
        type: 'bar',
        data: durations,
        itemStyle: {
          color: '#1976D2',
          borderRadius: [4, 4, 0, 0],
        },
      },
      {
        name: '运动次数',
        type: 'line',
        yAxisIndex: 1,
        data: counts,
        itemStyle: {
          color: '#26A69A',
        },
      },
    ],
  });
};

// 监听时间范围和运动类型的变化
watch([timeRange, exerciseType], async ([newTimeRange, newExerciseType]) => {
  try {
    await exerciseStore.fetchStats(
      newTimeRange,
      newExerciseType === 'all' ? 0 : parseInt(newExerciseType),
    );
    updateChart();
  } catch (error) {
    handleError(error as AxiosError);
  }
});

// 监听 stats 的变化
watch(
  () => stats.value,
  () => {
    updateChart();
  },
  { deep: true },
);

onMounted(async () => {
  try {
    await exerciseStore.fetchStats();
    initChart();
  } catch (error) {
    handleError(error as AxiosError);
  }
});

onUnmounted(() => {
  if (chart) {
    chart.dispose();
    chart = null;
  }
});
</script>

<style lang="scss" scoped>
.stats-card {
  transition: all 0.3s ease;
  border-radius: 8px;
  overflow: hidden;
  height: 100%;

  &:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  }

  .q-card__section {
    padding: 12px;
  }
}

.text-h5 {
  font-weight: 500;
  margin: 0;
  line-height: 1.2;
}

.text-subtitle2 {
  font-weight: 500;
}

.text-caption {
  margin-top: 2px;
}

.filter-select {
  min-width: 120px;
  max-width: 160px;
}

:deep(.q-table--light) {
  background: transparent;

  thead tr {
    background: #f5f5f5;

    th {
      color: #1a1a1a;
      font-weight: 600;
      border-bottom: 2px solid #e0e0e0;
    }
  }

  tbody tr {
    background: #fff;

    &:nth-child(even) {
      background: #fafafa;
    }

    &:hover {
      background: #f0f0f0;
    }

    td {
      color: #1a1a1a;
      border-bottom: 1px solid #e0e0e0;
    }
  }
}

:deep(.q-table--dark) {
  thead tr {
    background: rgba(0, 0, 0, 0.2);

    th {
      color: rgba(255, 255, 255, 0.9);
      border-bottom: 2px solid rgba(255, 255, 255, 0.1);
    }
  }

  tbody tr {
    background: transparent;

    &:nth-child(even) {
      background: rgba(0, 0, 0, 0.1);
    }

    &:hover {
      background: rgba(0, 0, 0, 0.2);
    }

    td {
      color: rgba(255, 255, 255, 0.8);
      border-bottom: 1px solid rgba(255, 255, 255, 0.1);
    }
  }
}

@media (max-width: 599px) {
  .filter-select {
    min-width: 100px;
    max-width: 140px;
  }

  .row.justify-between {
    flex-direction: column;
    align-items: flex-start !important;
    gap: 8px;
  }

  .row.q-gutter-sm {
    width: 100%;
    justify-content: space-between;
  }

  .col-6 {
    padding: 4px;
  }

  .stats-card {
    .q-card__section {
      padding: 8px;
    }

    .text-h5 {
      font-size: 1.1rem;
    }

    .text-subtitle2 {
      font-size: 0.9rem;
    }

    .text-caption {
      font-size: 0.8rem;
    }

    .q-icon {
      font-size: 1.5rem;
    }
  }
}
</style>
