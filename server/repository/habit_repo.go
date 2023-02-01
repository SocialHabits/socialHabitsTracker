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
		Type:      models.HabitType(habitInput.Type),
		Completed: 0,
		Streak:    0,
		Failed:    0,
		Total:     0,
		Goal:      habitInput.Goal,
		StartDate: habitInput.StartDate,
		EndDate:   habitInput.EndDate,
		UserId:    &id,
	}

	err := h.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("name = ?", habitInput.Name).Find(&habit).Error

		// check if habit already exists
		if err == nil {
			return fmt.Errorf("habit already exists")
		}

		if habitInput.Type == "PRESET" {
			fmt.Printf("cannot create a habit of type PRESET")
			return fmt.Errorf("cannot create a habit of type PRESET")
		}

		if err := tx.Create(&habit).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return habit, nil
}
