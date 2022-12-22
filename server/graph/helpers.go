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
