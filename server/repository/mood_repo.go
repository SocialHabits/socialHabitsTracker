package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/AntonioTrupac/socialHabitsTracker/graph/customTypes"
	"github.com/AntonioTrupac/socialHabitsTracker/models"
	"gorm.io/gorm"
)

type MoodRepository interface {
	CreateMood(moodInput customTypes.MoodInput, userId uint64) (*models.Mood, error)
	GetMoodsByUserID(userId uint64) ([]*models.Mood, error)
	UpdateMood(moodInput *customTypes.UpdateMoodInput, id int) error
	GetMoodByID(userId int) (*models.Mood, error)
	DeleteMood(id int) error
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

func (m MoodService) CreateMood(moodInput customTypes.MoodInput, userId uint64) (*models.Mood, error) {
	if userId == 0 {
		return nil, fmt.Errorf("user id cannot be 0")
	}

	mood := &models.Mood{
		Note:      *moodInput.Note,
		Type:      mapMoodTypes(moodInput.Types),
		Intensity: mapMoodIntensity(moodInput.Intensity),
		UserId:    userId,
	}

	err := m.DB.Create(&mood).Error

	if err != nil {
		return nil, err
	}

	return mood, nil
}

func (m MoodService) GetMoodsByUserID(userId uint64) ([]*models.Mood, error) {
	var moods []*models.Mood

	err := m.DB.Where("user_id = ?", userId).Find(&moods).Error

	if err != nil {
		return nil, err
	}

	return moods, nil
}

func (m MoodService) UpdateMood(moodInput *customTypes.UpdateMoodInput, id int) error {
	updateMood := models.Mood{
		ID:        uint64(id),
		Note:      *moodInput.Note,
		UpdatedAt: time.Now(),
	}

	err := m.DB.Model(models.Mood{}).Where("id=?", id).Updates(&updateMood).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println(err)
		return fmt.Errorf("mood with id %d not found", id)
	}

	return err
}

func (m MoodService) GetMoodByID(userId int) (*models.Mood, error) {
	var mood models.Mood

	err := m.DB.First(&mood, userId).Error

	if err != nil {
		return nil, err
	}

	return &mood, nil
}

func (m MoodService) DeleteMood(id int) error {
	mood, err := m.GetMoodByID(id)

	if err != nil {
		return err
	}

	err = m.DB.Delete(mood).Error

	if err != nil {
		return err
	}

	return nil
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
