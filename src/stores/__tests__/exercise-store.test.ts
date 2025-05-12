import { describe, it, expect, beforeEach, vi } from 'vitest';
import { setActivePinia, createPinia } from 'pinia';
import { useExerciseStore } from '../exercise-store';
import { api } from '../../boot/axios';
import type { ExerciseRecord, ExerciseStats, SportType } from 'src/types/exercise';

// Mock axios
vi.mock('../../boot/axios', () => ({
  api: {
    get: vi.fn(),
    post: vi.fn(),
    put: vi.fn(),
    delete: vi.fn(),
  },
}));

describe('Exercise Store', () => {
  beforeEach(() => {
    setActivePinia(createPinia());
    vi.clearAllMocks();
  });

  describe('State', () => {
    it('should initialize with default values', () => {
      const store = useExerciseStore();
      expect(store.exercises).toEqual([]);
      expect(store.sportTypes).toEqual([]);
      expect(store.stats).toEqual({
        total_duration: 0,
        total_calories: 0,
        exercise_count: 0,
        average_duration: 0,
        average_calories: 0,
        daily_duration: [],
        daily_count: [],
      });
      expect(store.loading).toBe(false);
    });
  });

  describe('Actions', () => {
    it('should fetch exercises successfully', async () => {
      const mockExercises: ExerciseRecord[] = [
        {
          id: 1,
          sport_type: { id: 1, name: '跑步' },
          duration: 30,
          calories: 300,
          date: '2024-03-20',
        },
      ];
      vi.mocked(api.get).mockResolvedValueOnce({ data: mockExercises });

      const store = useExerciseStore();
      await store.fetchExercises();

      expect(api.get).toHaveBeenCalledWith('/records');
      expect(store.exercises).toEqual(mockExercises);
      expect(store.loading).toBe(false);
    });

    it('should fetch sport types successfully', async () => {
      const mockSportTypes: SportType[] = [
        { id: 1, name: '跑步' },
        { id: 2, name: '游泳' },
      ];
      vi.mocked(api.get).mockResolvedValueOnce({ data: mockSportTypes });

      const store = useExerciseStore();
      await store.fetchSportTypes();

      expect(api.get).toHaveBeenCalledWith('/sport-types');
      expect(store.sportTypes).toEqual(mockSportTypes);
    });

    it('should fetch stats successfully', async () => {
      const mockStats: ExerciseStats = {
        total_duration: 120,
        total_calories: 1200,
        exercise_count: 4,
        average_duration: 30,
        average_calories: 300,
        daily_duration: [30, 30, 30, 30],
        daily_count: [1, 1, 1, 1],
      };
      vi.mocked(api.get).mockResolvedValueOnce({ data: mockStats });

      const store = useExerciseStore();
      await store.fetchStats('week', 1);

      expect(api.get).toHaveBeenCalledWith('/records/stats', {
        params: {
          time_range: 'week',
          sport_type_id: 1,
        },
      });
      expect(store.stats).toEqual(mockStats);
    });

    it('should create record successfully', async () => {
      const mockRecord: ExerciseRecord = {
        id: 1,
        sport_type: { id: 1, name: '跑步' },
        duration: 30,
        calories: 300,
        date: '2024-03-20',
      };
      vi.mocked(api.post).mockResolvedValueOnce({ data: mockRecord });
      vi.mocked(api.get).mockResolvedValueOnce({ data: {} });

      const store = useExerciseStore();
      const result = await store.createRecord({
        sport_type_id: 1,
        duration: 30,
        calories: 300,
        date: '2024-03-20',
      });

      expect(api.post).toHaveBeenCalledWith('/records', {
        sport_type_id: 1,
        duration: 30,
        calories: 300,
        date: '2024-03-20',
      });
      expect(store.exercises).toContainEqual(mockRecord);
      expect(result).toEqual(mockRecord);
      expect(store.loading).toBe(false);
    });

    it('should update record successfully', async () => {
      const mockRecord: ExerciseRecord = {
        id: 1,
        sport_type: { id: 1, name: '跑步' },
        duration: 45,
        calories: 450,
        date: '2024-03-20',
      };
      vi.mocked(api.put).mockResolvedValueOnce({ data: mockRecord });
      vi.mocked(api.get).mockResolvedValueOnce({ data: {} });

      const store = useExerciseStore();
      store.exercises = [
        {
          id: 1,
          sport_type: { id: 1, name: '跑步' },
          duration: 30,
          calories: 300,
          date: '2024-03-20',
        },
      ];

      const result = await store.updateRecord(1, {
        duration: 45,
        calories: 450,
      });

      expect(api.put).toHaveBeenCalledWith('/records/1', {
        duration: 45,
        calories: 450,
      });
      expect(store.exercises[0]).toEqual(mockRecord);
      expect(result).toEqual(mockRecord);
      expect(store.loading).toBe(false);
    });

    it('should delete record successfully', async () => {
      vi.mocked(api.delete).mockResolvedValueOnce({});
      vi.mocked(api.get).mockResolvedValueOnce({ data: {} });

      const store = useExerciseStore();
      store.exercises = [
        {
          id: 1,
          sport_type: { id: 1, name: '跑步' },
          duration: 30,
          calories: 300,
          date: '2024-03-20',
        },
      ];

      await store.deleteRecord(1);

      expect(api.delete).toHaveBeenCalledWith('/records/1');
      expect(store.exercises).toHaveLength(0);
      expect(store.loading).toBe(false);
    });
  });

  describe('Getters', () => {
    it('should return exercises', () => {
      const store = useExerciseStore();
      const mockExercises: ExerciseRecord[] = [
        {
          id: 1,
          sport_type: { id: 1, name: '跑步' },
          duration: 30,
          calories: 300,
          date: '2024-03-20',
        },
      ];
      store.exercises = mockExercises;

      expect(store.getExercises).toEqual(mockExercises);
    });

    it('should return sport types', () => {
      const store = useExerciseStore();
      const mockSportTypes: SportType[] = [
        { id: 1, name: '跑步' },
        { id: 2, name: '游泳' },
      ];
      store.sportTypes = mockSportTypes;

      expect(store.getSportTypes).toEqual(mockSportTypes);
    });

    it('should return stats', () => {
      const store = useExerciseStore();
      const mockStats: ExerciseStats = {
        total_duration: 120,
        total_calories: 1200,
        exercise_count: 4,
        average_duration: 30,
        average_calories: 300,
        daily_duration: [30, 30, 30, 30],
        daily_count: [1, 1, 1, 1],
      };
      store.stats = mockStats;

      expect(store.getStats).toEqual(mockStats);
    });
  });
});
