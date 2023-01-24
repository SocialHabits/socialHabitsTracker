package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"
	"github.com/AntonioTrupac/socialHabitsTracker/middleware"

	generated "github.com/AntonioTrupac/socialHabitsTracker/graph"
	"github.com/AntonioTrupac/socialHabitsTracker/graph/customTypes"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateMood is the resolver for the createMood field.
func (r *mutationResolver) CreateMood(ctx context.Context, input customTypes.MoodInput) (*customTypes.Mood, error) {
	repoMood, err := r.MoodRepository.CreateMood(input)

	if err != nil {
		return nil, &gqlerror.Error{
			Message: "Could not create mood",
		}
	}

	mood := &customTypes.Mood{
		ID:        int(repoMood.ID),
		Note:      &repoMood.Note,
		Types:     generated.ConvertMoodTypeToEnum(repoMood.Type),
		Intensity: generated.ConvertMoodIntensityToEnum(repoMood.Intensity),
	}

	return mood, nil
}

// UpdateMood is the resolver for the updateMood field.
func (r *mutationResolver) UpdateMood(ctx context.Context, id int, input *customTypes.MoodInput) (*customTypes.Mood, error) {
	panic(fmt.Errorf("not implemented: UpdateMood - updateMood"))
}

// DeleteMood is the resolver for the deleteMood field.
func (r *mutationResolver) DeleteMood(ctx context.Context, id int) (bool, error) {
	err := r.MoodRepository.DeleteMood(id)

	if err != nil {
		return false, &gqlerror.Error{
			Message: "Could not delete mood",
		}
	}

	return true, nil
}

// GetMoods is the resolver for the getMoods field.
func (r *queryResolver) GetMoods(ctx context.Context, userID *int) ([]*customTypes.Mood, error) {
	userClaims := middleware.GetValFromCtx(ctx)

	moods, err := r.MoodRepository.GetMoodsByUserID(userClaims.UserId)

	if err != nil {
		return nil, &gqlerror.Error{
			Message: "Could not get moods",
		}
	}

	var customMoods []*customTypes.Mood

	for _, mood := range moods {
		customMoods = append(customMoods, &customTypes.Mood{
			ID:        int(mood.ID),
			Note:      &mood.Note,
			Types:     generated.ConvertMoodTypeToEnum(mood.Type),
			Intensity: generated.ConvertMoodIntensityToEnum(mood.Intensity),
		})
	}

	return customMoods, nil
}

// GetMood is the resolver for the getMood field.
func (r *queryResolver) GetMood(ctx context.Context, id int) (*customTypes.Mood, error) {
	mood, err := r.MoodRepository.GetMoodByID(id)

	if err != nil {
		return nil, &gqlerror.Error{
			Message: "Could not get mood",
		}
	}

	return &customTypes.Mood{
		ID:        int(mood.ID),
		Note:      &mood.Note,
		Types:     generated.ConvertMoodTypeToEnum(mood.Type),
		Intensity: generated.ConvertMoodIntensityToEnum(mood.Intensity),
	}, nil
}
