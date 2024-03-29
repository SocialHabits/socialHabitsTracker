package graph

import "github.com/AntonioTrupac/socialHabitsTracker/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	BookRepository repository.BookRepository
	UserRepository repository.UserRepository
	MoodRepository repository.MoodRepository
}
