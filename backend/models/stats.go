package models

type Stats struct {
	TotalDuration   int64   `json:"total_duration"`   // 总运动时长（分钟）
	ExerciseCount   int64   `json:"exercise_count"`   // 运动次数
	AverageDuration float64 `json:"average_duration"` // 平均运动时长（分钟）
	AverageCalories float64 `json:"average_calories"` // 平均消耗卡路里
	DailyDuration   []int64 `json:"daily_duration"`   // 每日运动时长（分钟）
	DailyCount      []int64 `json:"daily_count"`      // 每日运动次数
}