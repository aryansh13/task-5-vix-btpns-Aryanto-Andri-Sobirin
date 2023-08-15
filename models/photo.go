package models

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Caption   string
	PhotoUrl  string
	UserID    uint
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
}
