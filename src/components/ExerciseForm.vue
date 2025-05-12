<template>
  <q-form @submit="onSubmit" class="q-gutter-md">
    <q-select
      v-model="form.sport_type_id"
      :options="filteredSportTypes"
      option-label="name"
      :option-value="(option) => option.id"
      label="运动类型"
      use-input
      input-debounce="0"
      @filter="filterSportTypes"
      @update:model-value="(val) => console.log('选择的运动类型:', val)"
      :rules="[(val) => !!val || '请选择或输入运动类型']"
    />

    <div class="row q-col-gutter-md">
      <div class="col-12 col-md-6">
        <q-input
          v-model="form.duration"
          label="运动时长（分钟）"
          type="number"
          :rules="[(val) => !!val || '请输入运动时长']"
          outlined
        />
      </div>
      <div class="col-12 col-md-6">
        <q-input
          v-model="form.calories"
          label="消耗卡路里"
          type="number"
          :rules="[(val) => !!val || '请输入消耗卡路里']"
          outlined
        />
      </div>
    </div>

    <div class="row q-col-gutter-md">
      <div class="col-12">
        <q-input v-model="form.notes" label="备注" type="textarea" outlined autogrow />
      </div>
    </div>

    <div class="row q-col-gutter-md">
      <div class="col-12">
        <ImageUpload
          v-model="form.image_url"
          @upload-success="handleUploadSuccess"
          @upload-error="handleUploadError"
        />
      </div>
    </div>

    <div class="row q-col-gutter-md">
      <div class="col-12">
        <q-input
          v-model="form.start_time"
          label="开始时间"
          type="datetime-local"
          :rules="[(val) => !!val || '请选择开始时间']"
          outlined
        />
      </div>
    </div>

    <div class="row justify-end q-mt-md">
      <q-btn label="取消" flat class="q-mr-sm" @click="$emit('cancel')" />
      <q-btn label="保存" type="submit" color="primary" :loading="loading" />
    </div>
  </q-form>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue';
import { useQuasar } from 'quasar';
import { useExerciseStore } from 'stores/exercise-store';
import ImageUpload from './ImageUpload.vue';
import type { ExerciseRecord, SportType } from '../types/exercise';

const props = defineProps<{
  record?: ExerciseRecord;
}>();

const emit = defineEmits<{
  (e: 'submit'): void;
  (e: 'cancel'): void;
}>();

const $q = useQuasar();
const exerciseStore = useExerciseStore();
const loading = ref(false);
const sportTypes = computed(() => {
  console.log('sportTypes computed:', exerciseStore.getSportTypes);
  return exerciseStore.getSportTypes;
});

const form = ref({
  sport_type_id: null as SportType | null,
  duration: props.record?.duration?.toString() || '',
  calories: props.record?.calories?.toString() || '',
  notes: props.record?.notes || '',
  image_url: props.record?.image_url || '',
  img_url_list: props.record?.img_url_list || '[]',
  start_time: props.record?.start_time
    ? new Date(props.record.start_time).toISOString().slice(0, 16)
    : new Date().toISOString().slice(0, 16),
  end_time: props.record?.end_time
    ? new Date(props.record.end_time).toISOString().slice(0, 16)
    : new Date().toISOString().slice(0, 16),
});

const filteredSportTypes = ref<SportType[]>([]);

// 获取运动类型
const fetchSportTypes = async () => {
  try {
    console.log('开始获取运动类型...');
    await exerciseStore.fetchSportTypes();
    filteredSportTypes.value = [...sportTypes.value];
    console.log('获取到的运动类型:', sportTypes.value);

    // 如果是编辑模式，使用现有的运动类型ID
    if (props.record?.sport_type_id) {
      const sportType = sportTypes.value.find((type) => type.id === props.record?.sport_type_id);
      if (sportType) {
        form.value.sport_type_id = sportType;
        console.log('编辑模式：使用现有运动类型:', sportType.name, 'ID:', sportType.id);
      }
    }
    // 否则使用第一个运动类型的ID
    else if (sportTypes.value && sportTypes.value.length > 0) {
      const firstSportType = sportTypes.value[0];
      if (firstSportType && firstSportType.id) {
        form.value.sport_type_id = firstSportType;
        console.log('新增模式：设置默认运动类型:', firstSportType.name, 'ID:', firstSportType.id);
      }
    }
  } catch (error) {
    console.error('获取运动类型失败:', error);
    $q.notify({
      type: 'negative',
      message: '获取运动类型失败',
    });
  }
};

onMounted(async () => {
  console.log('组件挂载，开始获取运动类型...');
  await fetchSportTypes();
});

// 监听 record 属性变化
watch(
  () => props.record,
  (newRecord) => {
    if (newRecord) {
      const sportType = exerciseStore.getSportTypes.find(
        (type) => type.id === newRecord.sport_type_id,
      );
      form.value = {
        sport_type_id: sportType || null,
        duration: newRecord.duration.toString(),
        calories: newRecord.calories.toString(),
        notes: newRecord.notes || '',
        image_url: newRecord.image_url || '',
        img_url_list: newRecord.img_url_list || '[]',
        start_time: new Date(newRecord.start_time).toISOString().slice(0, 16),
        end_time: new Date(newRecord.end_time).toISOString().slice(0, 16),
      };
    }
  },
  { immediate: true },
);

const handleUploadSuccess = (url: string) => {
  form.value.image_url = url;
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

const onSubmit = async () => {
  try {
    if (!form.value.sport_type_id) {
      $q.notify({
        type: 'negative',
        message: '请选择运动类型',
      });
      return;
    }

    loading.value = true;
    const startTime = new Date(form.value.start_time);
    const endTime = new Date(startTime.getTime() + parseInt(form.value.duration) * 60000);

    const data = {
      sport_type_id: form.value.sport_type_id.id,
      duration: parseInt(form.value.duration),
      calories: parseInt(form.value.calories),
      notes: form.value.notes,
      image_url: form.value.image_url,
      img_url_list: form.value.img_url_list,
      start_time: startTime.toISOString(),
      end_time: endTime.toISOString(),
    };

    if (props.record) {
      await exerciseStore.updateRecord(props.record.id, data);
    } else {
      await exerciseStore.createRecord(data);
    }

    emit('submit');
  } catch (error) {
    console.error('提交失败:', error);
    $q.notify({
      type: 'negative',
      message: '保存失败，请重试',
      position: 'top',
    });
  } finally {
    loading.value = false;
  }
};

const filterSportTypes = (val: string, update: (callback: () => void) => void) => {
  if (val === '') {
    update(() => {
      filteredSportTypes.value = [...sportTypes.value];
    });
    return;
  }

  update(() => {
    const needle = val.toLowerCase();
    filteredSportTypes.value = sportTypes.value.filter((type) =>
      type.name.toLowerCase().includes(needle),
    );
  });
};
</script>
