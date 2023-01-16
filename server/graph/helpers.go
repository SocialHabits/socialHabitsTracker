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
