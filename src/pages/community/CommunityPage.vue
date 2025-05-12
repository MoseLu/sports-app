<template>
  <q-page padding>
    <div class="row q-col-gutter-md">
      <div class="col-12">
        <q-card>
          <q-card-section>
            <div class="text-h6">社区动态</div>
          </q-card-section>
          <q-card-section>
            <q-list separator>
              <q-item v-for="post in posts" :key="post.id" class="q-mb-md">
                <q-item-section avatar>
                  <q-avatar>
                    <img :src="post.userAvatar" :alt="post.username" />
                  </q-avatar>
                </q-item-section>
                <q-item-section>
                  <q-item-label>{{ post.username }}</q-item-label>
                  <q-item-label caption>{{ post.content }}</q-item-label>
                  <q-item-label caption>{{ formatDate(post.createdAt) }}</q-item-label>
                </q-item-section>
              </q-item>
            </q-list>
          </q-card-section>
        </q-card>
      </div>
    </div>
  </q-page>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import { api } from 'boot/axios';
import { format } from 'date-fns';
import { handleError } from '../../services/error-handler';
import type { AxiosError } from 'axios';

interface Post {
  id: number;
  username: string;
  userAvatar: string;
  content: string;
  createdAt: string;
}

export default defineComponent({
  name: 'CommunityPage',
  setup() {
    const posts = ref<Post[]>([]);

    const loadPosts = async () => {
      try {
        const response = await api.get('/posts');
        posts.value = response.data;
      } catch (error) {
        handleError(error as AxiosError);
      }
    };

    const formatDate = (date: string) => {
      return format(new Date(date), 'yyyy-MM-dd HH:mm');
    };

    onMounted(async () => {
      await loadPosts();
    });

    return {
      posts,
      formatDate,
    };
  },
});
</script>
