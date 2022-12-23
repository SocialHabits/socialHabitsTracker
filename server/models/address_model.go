package models

import (
	"gorm.io/gorm"
	"time"
)

type Address struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Street    string         `gorm:"size:255;not null" json:"street"`
	City      string         `gorm:"size:255;not null" json:"city"`
	Country   string         `gorm:"size:255;not null" json:"country"`
	UserID    uint           `json:"user_id"`
}
