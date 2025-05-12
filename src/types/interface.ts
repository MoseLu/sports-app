export interface ExerciseType {
  label: string;
  value: string;
}

export interface SportType {
  id: number;
  name: string;
  description: string;
  icon: string;
}

export interface Exercise {
  id: number;
  sport_type_id: number;
  sport_type: SportType;
  duration: number;
  calories: number;
  start_time: string;
  end_time: string;
}

export interface Stats {
  total_duration: number;
  exercise_count: number;
  average_duration: number;
  average_calories: number;
  daily_duration: number[];
  daily_count: number[];
}

export const exerciseTypes: ExerciseType[] = [
  { label: '跑步', value: 'running' },
  { label: '游泳', value: 'swimming' },
  { label: '骑行', value: 'cycling' },
  { label: '步行', value: 'walking' },
  { label: '健身', value: 'fitness' },
  { label: '瑜伽', value: 'yoga' },
];
