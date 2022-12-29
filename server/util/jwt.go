package util

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"os"
	"time"
)

type Claims struct {
	UserID  int    `json:"userId"`
	TokenID string `json:"tokenId"`
	Email   string `json:"email"`
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

func GenerateAccessToken(userId int, email string) (string, error) {
	tokenId := uuid.New().String()

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserID:  userId,
		TokenID: tokenId,
		Email:   email,
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
