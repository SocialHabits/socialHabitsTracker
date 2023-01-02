package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/AntonioTrupac/socialHabitsTracker/graph/customTypes"
	"github.com/AntonioTrupac/socialHabitsTracker/middleware"
	"github.com/AntonioTrupac/socialHabitsTracker/models"
	"github.com/AntonioTrupac/socialHabitsTracker/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	GetUserById(id int) (*models.User, error)
	GetUsers() ([]*models.User, error)
	CreateUser(userInput *customTypes.UserInput) (*models.User, error)
	GetRoles() ([]*models.User, error)
	GetRoleByName(name string) (*models.User, error)
	Login(ctx context.Context, email, password string) (interface{}, error)
	CheckUserEmail(email string) (bool, error)
	// UpdateUser(userInput *customTypes.UserInput, id int) error
}

type UserService struct {
	Db *gorm.DB
}

var _ UserRepository = &UserService{}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		Db: db,
	}
}

func (u UserService) GetUserById(id int) (*models.User, error) {
	var user models.User

	err := u.Db.Model(&models.User{}).Preload("Address").Where("id = ?", id).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("User with id %d not found", id)
	}

	return &user, err
}

func (u UserService) GetUsers() ([]*models.User, error) {
	var users []*models.User

	err := u.Db.Model(&models.User{}).Preload("Address").Preload("Roles").Find(&users).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("No users found")
	}

	return users, err
}

func mapAddressInput(addressInput []*customTypes.AddressInput, id uint) []*models.Address {
	var addresses []*models.Address

	for _, address := range addressInput {
		addresses = append(addresses, &models.Address{
			Street:  address.Street,
			City:    address.City,
			Country: address.Country,
			UserID:  id,
		})
	}

	return addresses
}

// CheckUserEmail check if user email already exists
func (u UserService) CheckUserEmail(email string) (bool, error) {
	var user models.User

	err := u.Db.Model(&models.User{}).Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}

	return true, nil
}

type Result struct {
	ID   uint
	Name string
}

func mapUserInputRoleToUserRole(role *customTypes.Role) models.UserRole {
	return models.UserRole(role.String())
}

func (u UserService) CreateUser(userInput *customTypes.UserInput) (*models.User, error) {
	//var userRoles []*models.UserRoles

	user := &models.User{
		FirstName: userInput.FirstName,
		LastName:  userInput.LastName,
		Email:     userInput.Email,
		Password:  userInput.Password,
		Role:      mapUserInputRoleToUserRole(userInput.Role),
	}

	err := u.Db.Transaction(func(tx *gorm.DB) error {
		user.Password = util.HashPassword(userInput.Password)

		if err := tx.Omit(clause.Associations).Create(&user).Error; err != nil {
			return err
		}

		address := mapAddressInput(userInput.Address, user.ID)

		if err := tx.Create(&address).Error; err != nil {
			fmt.Printf("Error creating address: %v", err)
			return err
		}

		for _, value := range address {
			value.UserID = user.ID
		}

		user.Address = address

		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserService) GetRoles() ([]*models.User, error) {
	var roles []*models.User

	err := u.Db.Model(&models.User{}).Find(&roles).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("No roles found")
	}

	return roles, err
}

func (u UserService) GetRoleByName(name string) (*models.User, error) {
	var role models.User

	err := u.Db.Model(&models.User{}).Select("id, name").Where("name = ?", name).Find(&role).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("Role with name %s not found", name)
	}

	return &role, err
}

func (u UserService) Login(ctx context.Context, email, password string) (interface{}, error) {
	var user *models.User

	if err := u.Db.Model(&user).Preload("Address").Preload("Roles").Where("email LIKE ?", email).Take(&user).Error; err != nil {
		// if user not found
		if err == gorm.ErrRecordNotFound {
			return nil, &gqlerror.Error{
				Message: "User with this email not found",
			}
		}

		return nil, err

	}

	if err := util.ComparePassword(password, user.Password); err != nil {
		return nil, &gqlerror.Error{
			Message: "Incorrect password",
		}
	}

	accessToken, err := util.GenerateAccessToken(int(user.ID), user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	CA := middleware.GetCookieAccess(ctx)
	CA.SetToken(accessToken)
	CA.UserId = uint64(user.ID)
	CA.RoleName = user.Role

	return map[string]interface{}{
		"accessToken": accessToken,
	}, nil
}
