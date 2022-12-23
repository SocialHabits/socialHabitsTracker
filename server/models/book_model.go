package models

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Title     string         `json:"title"`
	Author    string         `json:"author"`
	Publisher string         `json:"publisher"`
}
