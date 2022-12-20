package repository

import (
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
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetUsers() ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
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

func (u UserService) CreateUser(userInput *customTypes.UserInput) (*models.User, error) {
	user := &models.User{
		FirstName: userInput.FirstName,
		LastName:  userInput.LastName,
		Email:     userInput.Email,
		Password:  userInput.Password,
		Address:   mapAddressInput(userInput.Address),
	}

	err := u.Db.Transaction(func(tx *gorm.DB) error {
		user.Password = util.HashPassword(userInput.Password)

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
