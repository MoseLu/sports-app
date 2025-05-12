import { defineStore } from 'pinia';
import { api } from '../boot/axios';
import type { ExerciseRecord, ExerciseStats, SportType } from 'src/types/exercise';

interface ExerciseState {
  exercises: ExerciseRecord[];
  sportTypes: SportType[];
  stats: ExerciseStats;
  loading: boolean;
}

export const useExerciseStore = defineStore('exercise', {
  state: (): ExerciseState => ({
    exercises: [],
    sportTypes: [],
    stats: {
      total_duration: 0,
      total_calories: 0,
      exercise_count: 0,
      average_duration: 0,
      average_calories: 0,
      daily_duration: [],
      daily_count: [],
    },
    loading: false,
  }),

  getters: {
    getExercises: (state) => state.exercises,
    getSportTypes: (state) => state.sportTypes,
    getStats: (state) => state.stats,
  },

  actions: {
    async fetchExercises() {
      try {
        this.loading = true;
        const response = await api.get('/records');
        this.exercises = response.data;
      } catch (error) {
        console.error('Error fetching exercises:', error);
        throw error;
      } finally {
        this.loading = false;
      }
    },

    async fetchSportTypes() {
      const response = await api.get('/sport-types');

      if (!response.data || !Array.isArray(response.data)) {
        throw new Error('获取运动类型失败：响应数据格式错误');
      }

      this.sportTypes = response.data;
    },

    async fetchStats(timeRange: string = 'week', sportTypeId: number = 0) {
      try {
        const response = await api.get('/records/stats', {
          params: {
            time_range: timeRange,
            sport_type_id: sportTypeId,
          },
        });
        this.stats = response.data;
      } catch (error) {
        console.error('Error fetching stats:', error);
        throw error;
      }
    },

    async createRecord(data: Omit<ExerciseRecord, 'id' | 'sport_type'>) {
      try {
        this.loading = true;
        const response = await api.post('/records', data);
        this.exercises.push(response.data);
        await this.fetchStats();
        return response.data;
      } catch (error) {
        console.error('Error creating record:', error);
        throw error;
      } finally {
        this.loading = false;
      }
    },

    async updateRecord(id: number, data: Partial<ExerciseRecord>) {
      try {
        this.loading = true;
        const response = await api.put(`/records/${id}`, data);
        const index = this.exercises.findIndex((ex) => ex.id === id);
        if (index !== -1) {
          this.exercises[index] = response.data;
        }
        await this.fetchStats();
        return response.data;
      } catch (error) {
        console.error('Error updating record:', error);
        throw error;
      } finally {
        this.loading = false;
      }
    },

    async deleteRecord(id: number) {
      try {
        this.loading = true;
        await api.delete(`/records/${id}`);
        this.exercises = this.exercises.filter((ex) => ex.id !== id);
        await this.fetchStats();
      } catch (error) {
        console.error('Error deleting record:', error);
        throw error;
      } finally {
        this.loading = false;
      }
    },
  },
});
