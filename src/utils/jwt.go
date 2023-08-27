package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	ID   uint
	Name string
}

var signingKey = []byte(viper.GetString("jwt.signingKey"))

func GenerateToken(id uint, name string) (string, error) {
	expires := time.Duration(viper.GetInt("jwt.expires"))
	claims := CustomClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    viper.GetString("app.name"),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expires * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

func ParseToken(tokenString string) (CustomClaims, error) {
	claims := CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if !token.Valid {
		return claims, errors.New("invalid token")
	}
	return claims, err
}

func IsTokenValid(token string) bool {
	_, err := ParseToken(token)
	if err != nil {
		return false
	}
	return true
}
