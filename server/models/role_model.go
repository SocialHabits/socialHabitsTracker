package models

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	Users     []*User        `gorm:"many2many:user_roles;"`
}
