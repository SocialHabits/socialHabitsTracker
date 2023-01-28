package generated

import (
	"github.com/AntonioTrupac/socialHabitsTracker/graph/customTypes"
	"github.com/AntonioTrupac/socialHabitsTracker/models"
)

func MapAddressModelToGqlType(addressesModel []*models.Address) []*customTypes.Address {
	var addresses []*customTypes.Address

	for _, address := range addressesModel {
		addresses = append(addresses, &customTypes.Address{
			Street:  address.Street,
			City:    address.City,
			Country: address.Country,
			ID:      int(address.ID),
			UserID:  int(address.UserID),
		})
	}

	return addresses
}

// ConvertModelRoleToEnum convert role name to graphql enum
func ConvertModelRoleToEnum(roleName models.UserRole) customTypes.Role {
	var role customTypes.Role

	switch roleName {
	case models.Admin:
		role = customTypes.RoleAdmin
	case models.Regular:
		role = customTypes.RoleRegular
	case models.Premium:
		role = customTypes.RolePremium
	case models.Trainer:
		role = customTypes.RoleTrainer
	}

	return role
}

func ConvertMoodTypeToEnum(moodType models.MoodType) customTypes.MoodType {
	var mood customTypes.MoodType

	switch moodType {
	case models.Happy:
		mood = customTypes.MoodTypeHappy
	case models.Sad:
		mood = customTypes.MoodTypeSad
	case models.Irritated:
		mood = customTypes.MoodTypeIrritated
	case models.Angry:
		mood = customTypes.MoodTypeAngry
	case models.Surprised:
		mood = customTypes.MoodTypeSurprised
	case models.Negative:
		mood = customTypes.MoodTypeNegative
	case models.Fearful:
		mood = customTypes.MoodTypeFearful
	case models.Anxious:
		mood = customTypes.MoodTypeAnxious
	case models.Calm:
		mood = customTypes.MoodTypeCalm
	case models.Excited:
		mood = customTypes.MoodTypeExcited
	case models.Disgusted:
		mood = customTypes.MoodTypeDisgusted
	case models.Relaxed:
		mood = customTypes.MoodTypeRelaxed
	case models.Tense:
		mood = customTypes.MoodTypeTense
	}

	return mood
}

func ConvertMoodIntensityToEnum(moodIntensity models.MoodIntensity) customTypes.MoodIntensity {
	var intensity customTypes.MoodIntensity

	switch moodIntensity {
	case models.Low:
		intensity = customTypes.MoodIntensityLow
	case models.Medium:
		intensity = customTypes.MoodIntensityMedium
	case models.High:
		intensity = customTypes.MoodIntensityHigh
	}

	return intensity
}
