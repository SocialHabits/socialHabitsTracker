package customTypes

import "github.com/AntonioTrupac/socialHabitsTracker/validator"

func (u UserInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("firstName", u.FirstName)
	v.Required("lastName", u.LastName)

	v.Required("email", u.Email)
	v.IsEmail("email", u.Email)

	v.Required("password", u.Password)
	v.MinLength("password", u.Password, 8)

	if u.Address == nil || len(u.Address) == 0 {
		v.Required("address", u.Address)
	}

	for _, address := range u.Address {
		v.Required("street", address.Street)
		v.Required("city", address.City)
		v.Required("country", address.Country)
	}

	return v.IsValid(), v.Errors
}

func (u LoginInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("email", u.Email)
	v.IsEmail("email", u.Email)

	v.Required("password", u.Password)
	v.MinLength("password", u.Password, 8)

	return v.IsValid(), v.Errors
}
