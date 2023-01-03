package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	generated "github.com/AntonioTrupac/socialHabitsTracker/graph"
	"github.com/AntonioTrupac/socialHabitsTracker/graph/customTypes"
	"github.com/AntonioTrupac/socialHabitsTracker/middleware"
	"github.com/AntonioTrupac/socialHabitsTracker/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input customTypes.UserInput) (*customTypes.User, error) {
	// check email
	util.CheckEmail(input.Email)

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
		Role:      input.Role,
		ID:        int(user.ID),
	}

	if err != nil {
		return nil, err
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
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (interface{}, error) {
	return r.UserRepository.Login(ctx, email, password)
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id int) (*customTypes.User, error) {
	userClaims := middleware.GetValFromCtx(ctx)

	if userClaims == nil || userClaims.UserId <= 0 && userClaims.IsLoggedIn == false {
		return nil, &gqlerror.Error{
			Message: "User is not authorized or logged in",
		}
	}

	user, err := r.UserRepository.GetUserById(id)

	userGql := &customTypes.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      nil,
		ID:        int(user.ID),
		Address:   generated.MapAddressModelToGqlType(user.Address),
	}

	if err != nil {
		return nil, err
	}

	return userGql, nil
}

// GetUsers is the resolver for the getUsers field.
func (r *queryResolver) GetUsers(ctx context.Context) ([]*customTypes.User, error) {
	userClaims := middleware.GetValFromCtx(ctx)

	if userClaims == nil || userClaims.UserId <= 0 && userClaims.IsLoggedIn == false || userClaims.RoleName != "regular" {
		return nil, &gqlerror.Error{
			Message: "User is not authorized or logged in",
		}
	}

	var usersGql []*customTypes.User
	usersRepo, err := r.UserRepository.GetUsers()

	for _, u := range usersRepo {
		usersGql = append(usersGql, &customTypes.User{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			//Role:      generated.MapRoleModelToGqlType(u.Roles),
			ID:      int(u.ID),
			Address: generated.MapAddressModelToGqlType(u.Address),
		})
	}

	if err != nil {
		return nil, err
	}

	return usersGql, nil
}

// GetRoles is the resolver for the getRoles field.
func (r *queryResolver) GetRoles(ctx context.Context) ([]customTypes.Role, error) {
	userClaims := middleware.GetValFromCtx(ctx)

	if userClaims == nil || userClaims.UserId <= 0 && userClaims.IsLoggedIn == false {
		return nil, &gqlerror.Error{
			Message: "User is not authorized or logged in",
		}
	}

	//roles, err := r.UserRepository.GetRoles()
	//
	//var rolesGql []*customTypes.Role
	//
	//for _, r := range roles {
	//	rolesGql = append(rolesGql, &customTypes.Role{
	//		ID:   int(r.ID),
	//		Name: r.Name,
	//	})
	//}
	//
	//if err != nil {
	//	return nil, err
	//}

	return nil, nil
}

// GetRole is the resolver for the getRole field.
func (r *queryResolver) GetRole(ctx context.Context, name customTypes.Role) (customTypes.Role, error) {
	userClaims := middleware.GetValFromCtx(ctx)

	if userClaims == nil || userClaims.UserId <= 0 && userClaims.IsLoggedIn == false {
		return "", &gqlerror.Error{
			Message: "User is not authorized or logged in",
		}
	}

	role, err := r.UserRepository.GetRoleByName(name)

	var roleGql customTypes.Role

	roleGql = customTypes.Role(role)

	return , nil
}
