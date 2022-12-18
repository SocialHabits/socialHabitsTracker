package models

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"size:255;not null" json:"first_name"`
	LastName  string `gorm:"size:255;not null" json:"last_name"`
	Email     string `gorm:"size:255;not null" json:"email"`
	Password  string `gorm:"size:255;not null" json:"password"`
	Address   []*Address
	Roles     []Role `gorm:"many2many:user_roles;"`
}

// Claims : struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	UserID string `json:"userid"`
	jwt.RegisteredClaims
}
