package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	FirstName string `gorm:"size:255;not null" json:"first_name"`
	LastName  string `gorm:"size:255;not null" json:"last_name"`
	Email     string `gorm:"unique;size:255;not null" json:"email"`
	Password  string `gorm:"size:255;not null" json:"password"`
	Address   []*Address
	Roles     []*Role `gorm:"many2many:user_roles;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
