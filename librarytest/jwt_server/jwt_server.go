package jwt_server

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(key interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"foo": "oscar",
		"nbf": time.Date(2018, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
