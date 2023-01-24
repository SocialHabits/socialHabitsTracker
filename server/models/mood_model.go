package models

import (
	"database/sql/driver"
	"fmt"
)

type MoodType string
type MoodIntensity string

const (
	Negative  MoodType = "NEGATIVE"
	Irritated MoodType = "IRRITATED"
	Tense     MoodType = "TENSE"
	Anxious   MoodType = "ANXIOUS"
	Excited   MoodType = "EXCITED"
	Calm      MoodType = "CALM"
	Happy     MoodType = "HAPPY"
	Relaxed   MoodType = "RELAXED"
	Surprised MoodType = "SURPRISED"
	Sad       MoodType = "SAD"
	Angry     MoodType = "ANGRY"
	Disgusted MoodType = "DISGUSTED"
	Fearful   MoodType = "FEARFUL"
)

const (
	High   MoodIntensity = "HIGH"
	Medium MoodIntensity = "MEDIUM"
	Low    MoodIntensity = "LOW"
)

func (u *MoodType) Scan(value interface{}) error {
	if v, ok := value.([]byte); ok {
		*u = MoodType(v)
		return nil
	}
	if v, ok := value.(string); ok {
		*u = MoodType(v)
		return nil
	}
	return fmt.Errorf("unexpected type %T for MoodType", value)
}

func (u MoodType) Value() (driver.Value, error) {
	return string(u), nil
}

func (u *MoodIntensity) Scan(value interface{}) error {
	if v, ok := value.([]byte); ok {
		*u = MoodIntensity(v)
		return nil
	}
	if v, ok := value.(string); ok {
		*u = MoodIntensity(v)
		return nil
	}
	return fmt.Errorf("unexpected type %T for MoodIntensity", value)
}

func (u MoodIntensity) Value() (driver.Value, error) {
	return string(u), nil
}

type Mood struct {
	ID        uint64        `gorm:"primarykey" json:"id"`
	Note      string        `gorm:"size:255;not null" json:"note"`
	Type      MoodType      `gorm:"type:enum('NEGATIVE', 'IRRITATED', 'TENSE', 'ANXIOUS', 'EXCITED', 'CALM', 'HAPPY', 'RELAXED', 'SURPRISED', 'SAD', 'ANGRY', 'DISGUSTED', 'FEARFUL')";"column:type" json:"type"`
	Intensity MoodIntensity `gorm:"type:enum('HIGH', 'MEDIUM', 'LOW')";"column:intensity" json:"intensity"`
	UserId    uint64        `gorm:"not null" json:"user_id"`
}

func (Mood) TableName() string {
	return "moods"
}
