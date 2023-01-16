package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	generated "github.com/AntonioTrupac/socialHabitsTracker/graph"
	"github.com/AntonioTrupac/socialHabitsTracker/graph/customTypes"
	"github.com/AntonioTrupac/socialHabitsTracker/middleware"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input customTypes.UserInput) (*customTypes.User, error) {
	isValid := validation(ctx, input)

	if !isValid {
		return nil, ErrInput
	}

	// check if user email already exists
	userExists, err := r.UserRepository.CheckUserEmail(input.Email)

	if userExists {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "User with this email already exists",
		})
		return nil, err
	}

	user, err := r.UserRepository.CreateUser(&input)

	createdUser := &customTypes.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Address:   generated.MapAddressModelToGqlType(user.Address),
		Role:      generated.ConvertModelRoleToEnum(user.Role),
		ID:        int(user.ID),
	}

	if err != nil {
		return nil, &gqlerror.Error{
			Message: "Could not create user",
		}
	}

	return createdUser, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input customTypes.UserInput) (*customTypes.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*customTypes.User, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input customTypes.LoginInput) (interface{}, error) {
	isValid := validation(ctx, input)

	if !isValid {
		return nil, ErrInput
	}

	return r.UserRepository.Login(ctx, input.Email, input.Password)
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id int) (*customTypes.User, error) {
	userClaims := middleware.GetValFromCtx(ctx)

	if userClaims == nil || userClaims.UserId <= 0 && userClaims.IsLoggedIn == false && userClaims.RoleName != "REGULAR" {
		return nil, &gqlerror.Error{
			Message: "User is not authorized or logged in",
		}
	}

	user, err := r.UserRepository.GetUserById(id)

	userGql := &customTypes.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      generated.ConvertModelRoleToEnum(user.Role),
		ID:        int(user.ID),
		Address:   generated.MapAddressModelToGqlType(user.Address),
	}

	if err != nil {
		return nil, &gqlerror.Error{
			Message: "Could not return user by id",
		}
	}

	return userGql, nil
}

// GetUsers is the resolver for the getUsers field.
func (r *queryResolver) GetUsers(ctx context.Context) ([]*customTypes.User, error) {
	// commented out just for testing purposes

	//userClaims := middleware.GetValFromCtx(ctx)
	//
	//fmt.Println("USER CLAIMS: ", userClaims)
	//
	//if userClaims == nil || userClaims.UserId <= 0 && userClaims.IsLoggedIn == false || userClaims.RoleName != "REGULAR" {
	//	return nil, &gqlerror.Error{
	//		Message: "User is not authorized or logged in",
	//	}
	//}

	var usersGql []*customTypes.User
	usersRepo, err := r.UserRepository.GetUsers()

	for _, u := range usersRepo {
		usersGql = append(usersGql, &customTypes.User{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Role:      generated.ConvertModelRoleToEnum(u.Role),
			ID:        int(u.ID),
			Address:   generated.MapAddressModelToGqlType(u.Address),
		})
	}

	if err != nil {
		return nil, err
	}

	return usersGql, nil
}

// GetRole is the resolver for the getRole field.
func (r *queryResolver) GetRole(ctx context.Context, id int) (customTypes.Role, error) {
	userClaims := middleware.GetValFromCtx(ctx)

	if userClaims == nil || userClaims.UserId <= 0 && userClaims.IsLoggedIn == false {
		return "", &gqlerror.Error{
			Message: "User is not authorized or logged in",
		}
	}

	role, err := r.UserRepository.GetRoleByUserID(id)

	if err != nil {
		return "", err
	}

	var roleGql customTypes.Role

	// convert role name to enum
	switch role {
	case "ADMIN":
		roleGql = customTypes.RoleAdmin
	case "REGULAR":
		roleGql = customTypes.RoleRegular
	case "PREMIUM":
		roleGql = customTypes.RolePremium
	case "TRAINER":
		roleGql = customTypes.RoleTrainer
	}

	return roleGql, nil
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
var (
	ErrInput = errors.New("input errors")
)
