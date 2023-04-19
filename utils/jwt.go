package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

type JwtCustomClaims struct {
	ID   int
	Name string
	jwt.RegisteredClaims
}

var jwtSignKey = []byte(viper.GetString("jwt.key"))

// 签发token
func GenerateToken(id int, name string) (string, error) {
	claims := JwtCustomClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.expire") * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "parrot",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSignKey)
}

// 解析token
func ParseToken(tokenString string) (JwtCustomClaims, error) {
	claims := JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSignKey, nil
	})
	if err == nil && !token.Valid {
		return claims, errors.New("token已经失效")
	}
	return claims, err
}

// token是否有效
func IsTokenValid(tokenString string) bool {
	_, err := ParseToken(tokenString)
	return err != nil
}
