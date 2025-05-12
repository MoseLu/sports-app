<template>
  <div class="q-pa-md">
    <div class="row justify-between items-center q-mb-md">
      <h5 class="q-ma-none">运动记录</h5>
      <div class="row q-gutter-md">
        <q-select
          v-model="timeRange"
          :options="timeRangeOptions"
          dense
          outlined
          emit-value
          map-options
          class="filter-select"
          label="时间范围"
          :dark="$q.dark.isActive"
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
          :dark="$q.dark.isActive"
        />
      </div>
    </div>

    <q-card class="q-mb-md">
      <q-card-section class="q-pa-sm">
        <div class="text-subtitle2 text-grey-8 q-mb-xs">详细数据</div>
        <div class="table-wrapper">
          <q-table
            :rows="filteredExercises"
            :columns="columns"
            row-key="id"
            :loading="loading"
            v-model:pagination="pagination"
            @request="onRequest"
            separator="cell"
            flat
            bordered
            class="my-sticky-header-last-column-table"
            :dark="$q.dark.isActive"
          >
            <template v-slot:header="props">
              <q-tr :props="props">
                <q-th
                  v-for="col in props.cols"
                  :key="col.name"
                  :props="props"
                  :class="{ 'sticky-column': col.name === 'actions' }"
                >
                  {{ col.label }}
                </q-th>
              </q-tr>
            </template>
            <template v-slot:body-cell-actions="props">
              <q-td :props="props" class="sticky-column">
                <q-btn-group flat>
                  <q-btn flat round color="primary" icon="edit" @click="editExercise(props.row)" />
                  <q-btn
                    flat
                    round
                    color="negative"
                    icon="delete"
                    @click="deleteExercise(props.row)"
                  />
                </q-btn-group>
              </q-td>
            </template>

            <template v-slot:body-cell-image="props">
              <q-td :props="props">
                <q-img
                  v-if="props.row.image_url"
                  :src="getImageUrl(props.row.image_url)"
                  style="width: 50px; height: 50px; cursor: pointer"
                  fit="cover"
                  class="rounded-borders"
                  @error="handleImageError"
                  @click="showImagePreview(props.row.image_url)"
                  :ratio="1"
                >
                  <template v-slot:loading>
                    <q-spinner-dots color="primary" />
                  </template>
                  <template v-slot:error>
                    <div class="text-caption text-grey">
                      <q-icon name="broken_image" size="24px" />
                    </div>
                  </template>
                </q-img>
                <span v-else>无图片</span>
              </q-td>
            </template>

            <template v-slot:body-cell-type="props">
              <q-td :props="props">
                {{ props.row.sport_type?.name || '未知' }}
              </q-td>
            </template>
          </q-table>
        </div>
      </q-card-section>
    </q-card>

    <q-page-sticky position="bottom-right" :offset="[18, 18]">
      <q-btn
        round
        style="z-index: 2006"
        color="primary"
        icon="add"
        size="lg"
        @click="showAddDialog = true"
      />
    </q-page-sticky>

    <q-dialog v-model="showAddDialog" persistent>
      <q-card style="min-width: 350px">
        <q-card-section>
          <div class="text-h6">添加运动记录</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-select
            v-model="newRecord.sport_type_id"
            :options="filteredSportTypes"
            option-label="name"
            option-value="id"
            label="运动类型"
            outlined
            dense
            use-input
            input-debounce="0"
            @filter="filterSportTypes"
            :rules="[(val) => !!val || '请选择运动类型']"
          />
          <q-input
            v-model.number="newRecord.duration"
            label="运动时长(分钟)"
            type="number"
            outlined
            dense
            class="q-mt-sm"
            :rules="[(val) => val > 0 || '请输入有效的运动时长']"
          />
          <q-input
            v-model.number="newRecord.calories"
            label="消耗卡路里"
            type="number"
            outlined
            dense
            class="q-mt-sm"
            :rules="[(val) => val > 0 || '请输入有效的卡路里消耗']"
          />
          <div class="q-mt-sm">
            <ImageUpload
              v-model="newRecord.image_url"
              @upload-success="handleUploadSuccess"
              @upload-error="handleUploadError"
            />
          </div>
        </q-card-section>

        <q-card-actions align="right" class="text-primary">
          <q-btn flat label="取消" v-close-popup />
          <q-btn flat label="确定" @click="addRecord" :disable="!isFormValid" />
        </q-card-actions>
      </q-card>
    </q-dialog>

    <q-dialog v-model="showEditDialog">
      <q-card style="min-width: 350px">
        <q-card-section>
          <div class="text-h6">编辑运动记录</div>
        </q-card-section>

        <q-card-section>
          <ExerciseForm
            v-if="selectedExercise"
            :record="selectedExercise"
            @submit="handleEditSubmit"
            @cancel="showEditDialog = false"
          />
        </q-card-section>
      </q-card>
    </q-dialog>

    <!-- 图片预览对话框 -->
    <q-dialog v-model="showPreviewDialog">
      <q-card style="min-width: 350px">
        <q-card-section class="row items-center q-pb-none">
          <div class="text-h6">图片预览</div>
          <q-space />
          <q-btn icon="close" flat round dense v-close-popup />
        </q-card-section>

        <q-card-section class="q-pa-md">
          <div class="row justify-center">
            <q-img :src="previewImageUrl" style="max-width: 100%; max-height: 70vh" fit="contain">
              <template v-slot:loading>
                <q-spinner-dots color="primary" size="40px" />
              </template>
              <template v-slot:error>
                <div class="text-h6 text-grey">
                  <q-icon name="broken_image" size="40px" />
                  <div class="q-mt-sm">图片加载失败</div>
                </div>
              </template>
            </q-img>
          </div>
        </q-card-section>
      </q-card>
    </q-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useExerciseStore } from 'stores/exercise-store';
import { useQuasar } from 'quasar';
import ExerciseForm from './ExerciseForm.vue';
import { handleError } from '../services/error-handler';
import type { AxiosError } from 'axios';
import type { ExerciseRecord, SportType } from '../types/exercise';
import ImageUpload from './ImageUpload.vue';

const $q = useQuasar();
const exerciseStore = useExerciseStore();
const exercises = computed(() => exerciseStore.getExercises);
const loading = computed(() => exerciseStore.loading);

const timeRange = ref('week');
const exerciseType = ref('all');
const showAddDialog = ref(false);
const showEditDialog = ref(false);
const selectedExercise = ref<ExerciseRecord | undefined>(undefined);
const showPreviewDialog = ref(false);
const previewImageUrl = ref('');

const newRecord = ref({
  sport_type_id: null as SportType | null,
  duration: 0,
  calories: 0,
  start_time: '',
  end_time: '',
  image_url: '',
});

const isFormValid = computed(() => {
  return (
    newRecord.value.sport_type_id !== null &&
    newRecord.value.duration > 0 &&
    newRecord.value.calories > 0
  );
});

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

const columns = [
  {
    name: 'type',
    label: '运动类型',
    field: (row: ExerciseRecord) => {
      const sportType = exerciseStore.getSportTypes.find((type) => type.id === row.sport_type_id);
      return sportType?.name || '未知';
    },
    align: 'center' as const,
  },
  {
    name: 'duration',
    label: '时长(分钟)',
    field: 'duration',
    align: 'center' as const,
  },
  {
    name: 'calories',
    label: '消耗卡路里',
    field: 'calories',
    align: 'center' as const,
  },
  {
    name: 'image',
    label: '运动图片',
    field: 'image_url',
    align: 'center' as const,
  },
  {
    name: 'actions',
    label: '操作',
    field: 'actions',
    align: 'center' as const,
    style: 'width: 100px',
    classes: 'sticky-column',
  },
];

const pagination = ref({
  sortBy: 'date',
  descending: true,
  page: 1,
  rowsPerPage: 10,
});

const filteredExercises = computed(() => {
  let data = Array.isArray(exercises.value) ? exercises.value : [];

  // 根据时间范围筛选
  if (timeRange.value === 'month') {
    const now = new Date();
    const startOfMonth = new Date(now.getFullYear(), now.getMonth(), 1);
    const endOfMonth = new Date(now.getFullYear(), now.getMonth() + 1, 0);
    data = data.filter((exercise) => {
      const date = new Date(exercise.start_time);
      return date >= startOfMonth && date <= endOfMonth;
    });
  } else if (timeRange.value === 'year') {
    const now = new Date();
    const startOfYear = new Date(now.getFullYear(), 0, 1);
    const endOfYear = new Date(now.getFullYear(), 11, 31);
    data = data.filter((exercise) => {
      const date = new Date(exercise.start_time);
      return date >= startOfYear && date <= endOfYear;
    });
  } else {
    // 本周
    const now = new Date();
    const startOfWeek = new Date(now.setDate(now.getDate() - now.getDay() + 1));
    const endOfWeek = new Date(now.setDate(now.getDate() - now.getDay() + 7));
    data = data.filter((exercise) => {
      const date = new Date(exercise.start_time);
      return date >= startOfWeek && date <= endOfWeek;
    });
  }

  // 根据运动类型筛选
  if (exerciseType.value !== 'all') {
    data = data.filter((exercise) => {
      console.log(
        'Exercise type:',
        exercise.sport_type?.name,
        'Selected type:',
        exerciseType.value,
      );
      return exercise.sport_type?.name === exerciseType.value;
    });
  }

  return data;
});

const filteredSportTypes = ref<SportType[]>([]);

const filterSportTypes = (val: string, update: (callback: () => void) => void) => {
  if (val === '') {
    update(() => {
      filteredSportTypes.value = [...exerciseStore.getSportTypes];
    });
    return;
  }

  update(() => {
    const needle = val.toLowerCase();
    filteredSportTypes.value = exerciseStore.getSportTypes.filter((type) =>
      type.name.toLowerCase().includes(needle),
    );
  });
};

onMounted(async () => {
  try {
    await Promise.all([
      exerciseStore.fetchExercises(),
      exerciseStore.fetchSportTypes(),
      exerciseStore.fetchStats(),
    ]);
  } catch (error) {
    handleError(error as AxiosError);
  }
});

interface RequestProps {
  pagination: {
    page: number;
    rowsPerPage: number;
    sortBy: string;
    descending: boolean;
  };
}

async function onRequest(props: RequestProps) {
  const { page, rowsPerPage, sortBy, descending } = props.pagination;
  await exerciseStore.fetchExercises();
  pagination.value.page = page;
  pagination.value.rowsPerPage = rowsPerPage;
  pagination.value.sortBy = sortBy;
  pagination.value.descending = descending;
}

function editExercise(exercise: ExerciseRecord) {
  const sportType = exerciseStore.getSportTypes.find((type) => type.id === exercise.sport_type_id);
  if (!sportType) {
    $q.notify({
      type: 'negative',
      message: '找不到对应的运动类型',
    });
    return;
  }
  selectedExercise.value = {
    ...exercise,
    sport_type: sportType,
    image_url: exercise.image_url || '',
    img_url_list: exercise.img_url_list || '[]',
  };
  showEditDialog.value = true;
}

const deleteExercise = (exercise: ExerciseRecord) => {
  $q.dialog({
    title: '确认删除',
    message: '确定要删除这条运动记录吗？',
    cancel: true,
    persistent: true,
  }).onOk(() => {
    exerciseStore
      .deleteRecord(exercise.id)
      .then(() => exerciseStore.fetchExercises())
      .catch((error) => {
        console.error('Error deleting exercise:', error);
        handleError(error as AxiosError);
      });
  });
};

const addRecord = async () => {
  try {
    const startTime = new Date();
    const endTime = new Date(startTime.getTime() + newRecord.value.duration * 60000);

    // 确保 sport_type_id 是有效的数字
    if (!newRecord.value.sport_type_id) {
      $q.notify({
        type: 'negative',
        message: '请选择运动类型',
        position: 'top',
        timeout: 3000,
      });
      return;
    }

    const data = {
      sport_type_id: Number(newRecord.value.sport_type_id.id),
      duration: Number(newRecord.value.duration),
      calories: Number(newRecord.value.calories),
      start_time: startTime.toISOString(),
      end_time: endTime.toISOString(),
      image_url: newRecord.value.image_url || '',
      img_url_list: '[]', // 设置一个有效的空 JSON 数组
      notes: '',
    };

    await exerciseStore.createRecord(data);
    showAddDialog.value = false;
    newRecord.value = {
      sport_type_id: null,
      duration: 0,
      calories: 0,
      start_time: '',
      end_time: '',
      image_url: '',
    };
  } catch (error) {
    console.error('Error adding record:', error);
    handleError(error as AxiosError);
  }
};

const handleEditSubmit = async () => {
  try {
    await exerciseStore.fetchExercises();
    showEditDialog.value = false;
  } catch (error) {
    console.error('Error updating exercise:', error);
    handleError(error as AxiosError);
  }
};

const handleUploadSuccess = (url: string) => {
  newRecord.value.image_url = url;
  $q.notify({
    type: 'positive',
    message: '图片上传成功',
  });
};

const handleUploadError = (error: string) => {
  $q.notify({
    type: 'negative',
    message: error,
  });
};

const getImageUrl = (url: string) => {
  if (!url) return '';
  return url;
};

const handleImageError = (error: Event) => {
  const imgElement = error.target as HTMLImageElement;
  if (imgElement) {
    imgElement.style.display = 'none';
    const errorDiv = imgElement.parentElement?.querySelector('.text-caption');
    if (errorDiv instanceof HTMLElement) {
      errorDiv.style.display = 'flex';
    }
  }
};

const showImagePreview = (url: string) => {
  if (!url) return;
  previewImageUrl.value = getImageUrl(url);
  showPreviewDialog.value = true;
};
</script>

<style lang="scss" scoped>
.filter-select {
  min-width: 140px;
  max-width: 180px;
}

.table-wrapper {
  position: relative;
  overflow-x: auto;
  max-height: calc(100vh - 200px);
}

.sticky-column {
  position: sticky !important;
  right: 0;
  z-index: 5;
  background-color: var(--q-card-bg);
}

.q-dialog {
  z-index: 2005;
}

.my-sticky-header-last-column-table {
  height: auto;
  min-height: 200px;
  max-height: calc(100vh - 200px);
  max-width: 100%;

  td:last-child {
    background-color: rgb(236, 240, 237);
    position: sticky;
    right: 0;
    z-index: 1;
  }

  tr th {
    position: sticky;
    z-index: 2;
    background: rgb(26, 175, 212);
  }

  thead tr:last-child th {
    top: 48px;
    z-index: 3;
  }

  thead tr:first-child th {
    top: 0;
    z-index: 1;
  }

  tr:last-child th:last-child {
    z-index: 3;
  }

  /** 固定最后一列 */
  td:last-child,
  th:last-child {
    position: sticky;
    right: 0;
  }

  tbody {
    scroll-margin-top: 48px;
  }
}

// 确保悬浮按钮在最上层
.q-page-sticky {
  z-index: 2006 !important;
}
</style>
