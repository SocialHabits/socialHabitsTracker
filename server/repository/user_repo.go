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

func (u UserService) CreateUser(userInput *customTypes.UserInput) (*models.User, error) {
	var userRoles []*models.UserRoles

	user := &models.User{
		FirstName: userInput.FirstName,
		LastName:  userInput.LastName,
		Email:     userInput.Email,
		Password:  userInput.Password,
		//Address:   mapAddressInputUser(userInput.Address),
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
			fmt.Println(value.ID)
			value.UserID = user.ID
		}

		// add address values to user model (for returning)
		user.Address = address

		for _, role := range userInput.Role {
			role, err := u.GetRoleByName(role.Name)

			if err != nil {
				return err
			}

			userRoles = append(userRoles, &models.UserRoles{
				UserID: user.ID,
				RoleID: role.ID,
			})
		}

		if err := tx.Create(&userRoles).Error; err != nil {
			fmt.Printf("Error creating user roles: %v", err)
			return err
		}

		// add role values to user model (for returning)
		for _, role := range userInput.Role {
			role, err := u.GetRoleByName(role.Name)

			if err != nil {
				return err
			}

			user.Roles = append(user.Roles, role)
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
