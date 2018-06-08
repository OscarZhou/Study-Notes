package jwt_server

import (
	"errors"
	"io/ioutil"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(key interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": "1122141",
		"pic": "dfshf",
		"rol": "mb",
		"exp": time.Now().Add(time.Duration(3600) * time.Second).Unix(),
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

type JWTDecoder struct {
	Subject   string `json:"sub"`
	ExpiresAt int64  `json:"exp"`
	Role      string `json:"rol"`
	Picture   string `json:"pic"`
}

func ParseJWT(token string) (*JWTDecoder, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() == jwt.SigningMethodRS256.Name {
			keyData, err := ioutil.ReadFile("ca.pem")
			if err != nil {
				return nil, err
			}

			key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
			if err != nil {
				return key, err
			}
			return key, nil
		}
		return nil, errors.New("unexpected signing method")
	})
	if err != nil {
		return nil, err
	}

	var (
		jwtDecoder JWTDecoder
		ok         bool
	)

	if claims, claimOk := jwtToken.Claims.(jwt.MapClaims); claimOk && jwtToken.Valid {
		jwtDecoder.ExpiresAt, ok = claims["exp"].(int64)
		if !ok {
			return nil, errors.New("exp is null")
		}
		jwtDecoder.Subject, ok = claims["sub"].(string)
		if !ok {
			return nil, errors.New("sub is null")
		}
		jwtDecoder.Role, ok = claims["rol"].(string)
		if !ok {
			return nil, errors.New("rol")
		}
		jwtDecoder.Role, ok = claims["pic"].(string)
		if !ok {
			return nil, errors.New("pic")
		}

		if jwtDecoder.Subject == "" {
			return nil, errors.New("sub is empty")
		}
		now := time.Now().Unix()
		if now > jwtDecoder.ExpiresAt {
			return nil, errors.New("token has been expired")
		}
	}

	return &jwtDecoder, nil
}
