package utils

import (
	"fmt"
	"restaurant-system/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyCustomClaims struct {
	ID   int
	Name string
	Role string
	jwt.RegisteredClaims
}
var secret_key = config.LocalConfig.JWTSecret

func GenerateToken(u_Id int, u_Name, u_Role string) (string, error) {

	claims := MyCustomClaims{
		ID:   u_Id,
		Name: u_Name,
		Role: u_Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "MyTaskManager",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret_key))
}

func VerifyByJWT(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret_key), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(*MyCustomClaims)

	if !ok {
		return "", fmt.Errorf("could not parse claims")
	}
	return claims.Role, nil
}
