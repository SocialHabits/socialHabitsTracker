package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type HabitType string

const (
	Preset HabitType = "PRESET"
	Custom HabitType = "CUSTOM"
)

func (h *HabitType) Scan(value interface{}) error {
	if v, ok := value.([]byte); ok {
		*h = HabitType(v)
		return nil
	}
	if v, ok := value.(string); ok {
		*h = HabitType(v)
		return nil
	}
	return fmt.Errorf("unexpected type %T for HabitType", value)
}

func (h HabitType) Value() (driver.Value, error) {
	return string(h), nil
}

type Habit struct {
	ID        uint      `gorm:"primarykey"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Skipped   int       `gorm:"not null" json:"skipped"`
	Type      HabitType `gorm:"type:enum('PRESET','CUSTOM')" column:"type" json:"type"`
	Completed int       `gorm:"not null" json:"completed"`
	Streak    int       `gorm:"not null" json:"streak"`
	Failed    int       `gorm:"not null" json:"failed"`
	Total     int       `gorm:"not null" json:"total"`
	Goal      int       `gorm:"not null" json:"goal"`
	StartDate time.Time `gorm:"not null" json:"start_date"`
	EndDate   time.Time `gorm:"not null" json:"end_date"`
	UserId    *uint     `json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Habit) TableName() string {
	return "habits"
}
