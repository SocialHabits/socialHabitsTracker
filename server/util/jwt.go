package util

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"os"
	"time"
)

type Claims struct {
	UserID   int    `json:"userId"`
	TokenID  string `json:"tokenId"`
	Email    string `json:"email"`
	RoleName string `json:"roleName"`
	jwt.RegisteredClaims
}

var accessTokenSecret = []byte(getAccessTokenSecret())
var refreshTokenSecret = []byte(getRefreshTokenSecret())

func getAccessTokenSecret() string {
	secret := os.Getenv("ACCESS_SECRET")

	if secret == "" {
		return ""
	}

	return secret
}

func getRefreshTokenSecret() string {
	secret := os.Getenv("REFRESH_SECRET")

	if secret == "" {
		return ""
	}

	return secret
}

func GenerateAccessToken(userId int, email string, role string) (string, error) {
	tokenId := uuid.New().String()

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserID:   userId,
		TokenID:  tokenId,
		Email:    email,
		RoleName: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	// after the token has been created, sign it with JWT_SECRET to make it secure
	token, err := tokenClaims.SignedString(accessTokenSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateIdToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there has been a problem with the signing method")
		}

		return accessTokenSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("there has been a problem with the claims")
	}

	return claims, nil
}
