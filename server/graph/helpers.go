package generated

import (
	"github.com/AntonioTrupac/socialHabitsTracker/graph/customTypes"
	"github.com/AntonioTrupac/socialHabitsTracker/models"
)

func MapAddressModelToGqlType(addressesModel []*models.Address) []*customTypes.Address {
	var addresses []*customTypes.Address

	for _, address := range addressesModel {
		addresses = append(addresses, &customTypes.Address{
			City:    address.City,
			Country: address.Country,
			Street:  address.Street,
			ID:      int(address.ID),
			UserID:  int(address.UserID),
		})
	}

	return addresses
}

//
//func MapRole()  {
//
//}

//func MapRoleModelToGqlType(roleModel []*models.Role) []*customTypes.Role {
//	var roles []*customTypes.Role
//
//	for _, r := range roleModel {
//		roles = append(roles, &customTypes.Role{
//			ID:   int(r.ID),
//			Name: r.Name,
//		})
//	}
//
//	return roles
//}
