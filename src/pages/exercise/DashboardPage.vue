<template>
  <q-page class="q-pa-md">
    <div class="row q-col-gutter-md">
      <div class="col-12">
        <ExerciseStats />
      </div>
      <div class="col-12">
        <ExerciseList />
      </div>
    </div>
  </q-page>
</template>

<script setup lang="ts">
import ExerciseStats from 'src/components/ExerciseStats.vue';
import ExerciseList from 'src/components/ExerciseList.vue';
</script>
