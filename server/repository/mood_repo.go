package repository

import (
	"github.com/AntonioTrupac/socialHabitsTracker/graph/customTypes"
	"github.com/AntonioTrupac/socialHabitsTracker/models"
	"gorm.io/gorm"
)

type MoodRepository interface {
	CreateMood(moodInput *customTypes.MoodInput) (*models.Mood, error)
	GetMoodsByUserID(id int) ([]*models.Mood, error)
	UpdateMood(moodInput *customTypes.MoodInput, id int) error
}

type MoodService struct {
	DB *gorm.DB
}

func (m MoodService) CreateMood(moodInput *customTypes.MoodInput) (*models.Mood, error) {
	//TODO implement me
	panic("implement me")
}

func (m MoodService) GetMoodsByUserID(id int) ([]*models.Mood, error) {
	//TODO implement me
	panic("implement me")
}

func (m MoodService) UpdateMood(moodInput *customTypes.MoodInput, id int) error {
	//TODO implement me
	panic("implement me")
}

var _ MoodRepository = &MoodService{}

func NewMoodService(db *gorm.DB) *MoodService {
	return &MoodService{
		DB: db,
	}
}
