package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	Admin   UserRole = "ADMIN"
	Regular UserRole = "REGULAR"
	Premium UserRole = "PREMIUM"
	Trainer UserRole = "TRAINER"
)

func (u *UserRole) Scan(value interface{}) error {
	if v, ok := value.([]byte); ok {
		*u = UserRole(v)
		return nil
	}
	if v, ok := value.(string); ok {
		*u = UserRole(v)
		return nil
	}
	return fmt.Errorf("unexpected type %T for UserRole", value)
}

func (u UserRole) Value() (driver.Value, error) {
	return string(u), nil
}

type User struct {
	ID        uint   `gorm:"primarykey"`
	FirstName string `gorm:"size:255;not null" json:"first_name"`
	LastName  string `gorm:"size:255;not null" json:"last_name"`
	Email     string `gorm:"size:255;not null" json:"email"`
	Password  string `gorm:"size:255;not null" json:"password"`
	Address   []*Address
	Role      UserRole `gorm:"type:enum('ADMIN','REGULAR','PREMIUM','TRAINER')"; column:"role" json:"role"`
	Mood      []*Mood
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
	return "users"
}
