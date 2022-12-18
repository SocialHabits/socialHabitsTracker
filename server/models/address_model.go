package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Street  string `gorm:"size:255;not null" json:"street"`
	City    string `gorm:"size:255;not null" json:"city"`
	Country string `gorm:"size:255;not null" json:"country"`
	UserID  uint   `json:"user_id"`
}
