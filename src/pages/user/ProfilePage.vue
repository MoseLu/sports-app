<template>
  <q-page padding>
    <div class="row q-col-gutter-md">
      <div class="col-12">
        <q-card>
          <q-card-section>
            <div class="text-h6">个人资料</div>
          </q-card-section>
          <q-card-section>
            <q-form @submit="onSubmit" class="q-gutter-md">
              <q-input
                v-model="user.username"
                label="用户名"
                :rules="[(val) => !!val || '请输入用户名']"
                readonly
              />
              <q-input
                v-model="user.email"
                label="邮箱"
                :rules="[(val) => !!val || '请输入邮箱']"
                readonly
              />
              <q-input
                v-model="user.nickname"
                label="昵称"
                :rules="[(val) => !!val || '请输入昵称']"
              />
              <q-input v-model="user.bio" label="个人简介" type="textarea" autogrow />
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
import { defineComponent, ref, onMounted } from 'vue';
import { showSuccess, handleError } from '../../services/error-handler';
import type { AxiosError } from 'axios';
import { useUserStore } from '../../stores/user';

export default defineComponent({
  name: 'ProfilePage',
  setup() {
    const user = ref({
      username: '',
      email: '',
      nickname: '',
      bio: '',
    });

    const userStore = useUserStore();

    const loadUserProfile = async () => {
      try {
        await userStore.fetchUser();
      } catch (error) {
        handleError(error as AxiosError);
      }
    };

    const onSubmit = async () => {
      try {
        await userStore.updateProfile({
          nickname: user.value.nickname,
          bio: user.value.bio,
        });
        showSuccess('个人资料更新成功');
      } catch (error) {
        handleError(error as AxiosError);
      }
    };

    onMounted(() => {
      void loadUserProfile();
    });

    return {
      user,
      onSubmit,
    };
  },
});
</script>
