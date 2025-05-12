export interface SportType {
  id: number;
  name: string;
  description: string;
  icon: string;
}

export interface ExerciseRecord {
  id: number;
  sport_type_id: number;
  sport_type: SportType;
  duration: number;
  calories: number;
  start_time: string;
  end_time: string;
  notes?: string;
  image_url: string;
  img_url_list: string;
}

export interface ExerciseStats {
  total_duration: number;
  total_calories: number;
  exercise_count: number;
  average_duration: number;
  average_calories: number;
  daily_duration: number[];
  daily_count: number[];
}
