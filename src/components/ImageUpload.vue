<template>
  <div class="image-upload">
    <q-file
      v-model="file"
      label="上传图片"
      accept=".jpg,.jpeg,.png,.gif"
      :max-file-size="5 * 1024 * 1024"
      @update:model-value="handleFileChange"
      outlined
      :loading="uploading"
      :disable="uploading"
    >
      <template v-slot:prepend>
        <q-icon name="image" />
      </template>
      <template v-slot:append>
        <q-icon name="close" @click.stop="clearFile" class="cursor-pointer" />
      </template>
    </q-file>

    <div v-if="previewUrl" class="preview-container q-mt-sm">
      <q-img :src="previewUrl" :ratio="16 / 9" class="preview-image" spinner-color="primary">
        <template v-slot:loading>
          <q-spinner-dots color="primary" />
        </template>
      </q-img>
      <q-btn flat round dense icon="delete" class="delete-btn" @click="clearFile" />
    </div>

    <q-banner v-if="error" class="bg-negative text-white q-mt-sm" rounded>
      {{ error }}
    </q-banner>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { api } from '../boot/axios';
import type { AxiosError } from 'axios';

const props = defineProps<{
  modelValue: string;
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void;
  (e: 'upload-success', url: string): void;
  (e: 'upload-error', error: string): void;
}>();

const file = ref<File | null>(null);
const previewUrl = ref(props.modelValue || '');
const uploading = ref(false);
const error = ref('');

// 监听 modelValue 变化
watch(
  () => props.modelValue,
  (newValue) => {
    previewUrl.value = newValue || '';
  },
);

// 处理文件选择
const handleFileChange = async (newFile: File | null) => {
  if (!newFile) {
    clearFile();
    return;
  }

  // 验证文件类型
  const validTypes = ['image/jpeg', 'image/png', 'image/gif'];
  if (!validTypes.includes(newFile.type)) {
    error.value = '只支持 JPG、PNG、GIF 格式的图片';
    clearFile();
    return;
  }

  // 验证文件大小
  if (newFile.size > 5 * 1024 * 1024) {
    error.value = '图片大小不能超过 5MB';
    clearFile();
    return;
  }

  error.value = '';
  previewUrl.value = URL.createObjectURL(newFile);
  await uploadFile(newFile);
};

// 上传文件
const uploadFile = async (file: File) => {
  try {
    uploading.value = true;
    const formData = new FormData();
    formData.append('image', file);

    const response = await api.post('/upload/image', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });

    emit('upload-success', response.data.url);
  } catch (err) {
    const error = err as AxiosError<{ error: string }>;
    const errorMessage = error.response?.data?.error || '上传失败';
    emit('upload-error', errorMessage);
  } finally {
    uploading.value = false;
  }
};

// 清除文件
const clearFile = () => {
  file.value = null;
  previewUrl.value = '';
  error.value = '';
  emit('update:modelValue', '');
};
</script>

<style lang="scss" scoped>
.image-upload {
  .preview-container {
    position: relative;
    width: 100%;
    max-width: 300px;
    border-radius: 8px;
    overflow: hidden;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);

    .preview-image {
      width: 100%;
      height: auto;
    }

    .delete-btn {
      position: absolute;
      top: 8px;
      right: 8px;
      background: rgba(0, 0, 0, 0.5);
      color: white;
    }
  }
}
</style>
