package jwt_helper

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

func SignJwt(key string, id int) (string, error) {
	claims := &JwtCustomClaims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	sign, errSign := token.SignedString([]byte(key))
	if errSign != nil {
		return "", errSign
	}

	return sign, nil
}
