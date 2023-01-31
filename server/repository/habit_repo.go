package repository

import (
	"fmt"
	"time"

	"github.com/AntonioTrupac/socialHabitsTracker/graph/customTypes"
	"github.com/AntonioTrupac/socialHabitsTracker/models"
	"gorm.io/gorm"
)

type HabitRepository interface {
	CreateHabit(habitInput customTypes.CreateHabitInput, userId uint64) (*models.Habit, error)
}

type HabitService struct {
	DB *gorm.DB
}

var _ HabitRepository = &HabitService{}

func NewHabitService(db *gorm.DB) *HabitService {
	return &HabitService{
		DB: db,
	}
}

// CreateHabit implements HabitRepository
func (h *HabitService) CreateHabit(habitInput customTypes.CreateHabitInput, userId uint64) (*models.Habit, error) {
	if userId == 0 {
		return nil, fmt.Errorf("user id cannot be 0")
	}

	id := uint(userId)

	habit := &models.Habit{
		Name:      habitInput.Name,
		Skipped:   0,
		Type:      models.HabitType(habitInput.Type),
		Completed: 0,
		Streak:    0,
		Failed:    0,
		Total:     0,
		Goal:      *habitInput.Goal,
		StartDate: time.Time{},
		EndDate:   time.Time{},
		UserId:    &id,
	}

	err := h.DB.Create(&habit)

	if err != nil {
		return nil, err.Error
	}

	return habit, nil
}
