package models

import (
	"time"
)

type Verification struct {
	ID        uint      `gorm:"primarykey"`
	Email     string    `gorm:"size:255;not null"`
	Code      string    `gorm:"size:6;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
} 