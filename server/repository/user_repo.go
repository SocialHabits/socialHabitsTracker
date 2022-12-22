package repository

import (
	"errors"
	"fmt"
	"github.com/AntonioTrupac/socialHabitsTracker/graph/customTypes"
	"github.com/AntonioTrupac/socialHabitsTracker/models"
	"github.com/AntonioTrupac/socialHabitsTracker/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	GetUserById(id int) (*models.User, error)
	GetUsers() ([]*models.User, error)
	CreateUser(userInput *customTypes.UserInput) (*models.User, error)
	GetRoles() ([]*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	CreateRole(roleInput *customTypes.RoleInput) (*models.Role, error)
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

	err := u.Db.Model(&models.User{}).Preload("Address").Find(&users).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("No users found")
	}

	return users, err
}

func mapAddressInput(addressInput []*customTypes.AddressInput) []*models.Address {
	var addresses []*models.Address

	for _, address := range addressInput {
		addresses = append(addresses, &models.Address{
			Street:  address.Street,
			City:    address.City,
			Country: address.Country,
		})
	}

	return addresses
}

func mapRolesInput(roleInput []*customTypes.RoleInput) []models.Role {
	var roles []models.Role

	for _, role := range roleInput {
		roles = append(roles, models.Role{
			Name: role.Name,
		})
	}

	return roles
}

func (u UserService) CreateUser(userInput *customTypes.UserInput) (*models.User, error) {
	user := &models.User{
		FirstName: userInput.FirstName,
		LastName:  userInput.LastName,
		Email:     userInput.Email,
		Password:  userInput.Password,
		Address:   mapAddressInput(userInput.Address),
		Roles:     mapRolesInput(userInput.Role),
	}

	err := u.Db.Transaction(func(tx *gorm.DB) error {
		user.Password = util.HashPassword(userInput.Password)

		// insert into users table
		if err := tx.Model(&user).Association("Roles").Append(user.Roles); err != nil {
			return err
		}

		if err := tx.Omit(clause.Associations).Create(user).Error; err != nil {
			return err
		}

		for _, value := range user.Address {
			value.UserID = user.ID
		}

		if err := tx.CreateInBatches(user.Address, 100).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserService) GetRoles() ([]*models.Role, error) {
	var roles []*models.Role

	err := u.Db.Model(&models.Role{}).Find(&roles).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("No roles found")
	}

	return roles, err
}

func (u UserService) GetRoleByName(name string) (*models.Role, error) {
	var role models.Role

	err := u.Db.Model(&models.Role{}).Where("name = ?", name).First(&role).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("Role with name %s not found", name)
	}

	return &role, err
}

func (u UserService) CreateRole(roleInput *customTypes.RoleInput) (*models.Role, error) {
	role := &models.Role{
		Name: roleInput.Name,
	}

	err := u.Db.Create(role).Error

	if err != nil {
		return nil, err
	}

	return role, nil
}
