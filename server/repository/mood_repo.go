package repository

import (
	"fmt"
	"github.com/AntonioTrupac/socialHabitsTracker/graph/customTypes"
	"github.com/AntonioTrupac/socialHabitsTracker/models"
	"gorm.io/gorm"
)

type MoodRepository interface {
	CreateMood(moodInput customTypes.MoodInput) (*models.Mood, error)
	GetMoodsByUserID(id int) ([]*models.Mood, error)
	UpdateMood(moodInput *customTypes.MoodInput, id int) error
}

type MoodService struct {
	DB *gorm.DB
}

var _ MoodRepository = &MoodService{}

func NewMoodService(db *gorm.DB) *MoodService {
	return &MoodService{
		DB: db,
	}
}

func (m MoodService) CreateMood(moodInput customTypes.MoodInput) (*models.Mood, error) {
	mood := &models.Mood{
		Note:      *moodInput.Note,
		Type:      mapMoodTypes(moodInput.Types),
		Intensity: mapMoodIntensity(moodInput.Intensity),
	}

	err := m.DB.Create(&mood).Error

	if err != nil {
		return nil, err
	}

	return mood, nil
}

func (m MoodService) GetMoodsByUserID(id int) ([]*models.Mood, error) {
	//TODO implement me
	panic("implement me")
}

func (m MoodService) UpdateMood(moodInput *customTypes.MoodInput, id int) error {
	//TODO implement me
	panic("implement me")
}

func mapMoodTypes(types customTypes.MoodType) models.MoodType {
	switch types {
	case customTypes.MoodTypeHappy:
		return models.Happy
	case customTypes.MoodTypeSad:
		return models.Sad
	case customTypes.MoodTypeAngry:
		return models.Angry
	case customTypes.MoodTypeCalm:
		return models.Calm
	case customTypes.MoodTypeAnxious:
		return models.Anxious
	case customTypes.MoodTypeExcited:
		return models.Excited
	case customTypes.MoodTypeDisgusted:
		return models.Disgusted
	case customTypes.MoodTypeFearful:
		return models.Fearful
	case customTypes.MoodTypeIrritated:
		return models.Irritated
	case customTypes.MoodTypeNegative:
		return models.Negative
	case customTypes.MoodTypeSurprised:
		return models.Surprised
	case customTypes.MoodTypeTense:
		return models.Tense
	case customTypes.MoodTypeRelaxed:
		return models.Relaxed
	default:
		panic(fmt.Sprintf("unknown mood type %v", types))
	}

}

func mapMoodIntensity(intensity customTypes.MoodIntensity) models.MoodIntensity {
	switch intensity {
	case customTypes.MoodIntensityHigh:
		return models.High
	case customTypes.MoodIntensityMedium:
		return models.Medium
	case customTypes.MoodIntensityLow:
		return models.Low
	default:
		panic(fmt.Sprintf("unknown mood intensity %v", intensity))
	}
}
