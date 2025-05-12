package models

import "time"

type Exercise struct {
    ID          int64     `json:"id" gorm:"primaryKey"`
    UserID      int64     `json:"user_id"`
    SportTypeID int64     `json:"sport_type_id"`
    Duration    int       `json:"duration"`
    Distance    float64   `json:"distance"`
    Calories    int       `json:"calories"`
    Date        time.Time `json:"date"`
    ImageURL    string    `json:"image_url" gorm:"size:255"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    User        User      `json:"user" gorm:"foreignKey:UserID"`
    SportType   SportType `json:"sport_type" gorm:"foreignKey:SportTypeID"`
}

func (Exercise) TableName() string {
    return "exercises"
} 