package validator

import "regexp"

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// IsEmail validates email address
func (v *Validator) IsEmail(field, email string) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if !emailRegexp.MatchString(email) {
		v.Errors[field] = "Invalid email address"

		return false
	}

	return true
}
