package repository

import (
	"fmt"

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

	id := uint64(userId)

	habit := &models.Habit{
		Name:      habitInput.Name,
		Skipped:   0,
		Type:      models.UserCreated,
		Completed: 0,
		Streak:    0,
		Failed:    0,
		Total:     0,
		Goal:      habitInput.Goal,
		StartDate: habitInput.StartDate,
		EndDate:   habitInput.EndDate,
		UserId:    &id,
	}

	if habit.Type != models.UserCreated {
		return nil, fmt.Errorf("habit type is not user created: %v", habit.Type)
	}

	err := h.DB.Create(&habit).Error

	if err != nil {
		return nil, err
	}

	return habit, nil
}
