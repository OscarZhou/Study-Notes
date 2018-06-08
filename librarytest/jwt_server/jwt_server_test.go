package jwt_server

import (
	"fmt"
	"io/ioutil"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
)

func TestJWTServer(t *testing.T) {

	keyData, err := ioutil.ReadFile("ca-key.pem")
	if err != nil {
		t.Error(err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		t.Error(err)
	}

	jwt, err := GenerateJWT(key)
	if err != nil {
		t.Error(err)

	}

	fmt.Println(jwt)
	jwtDecoder, err := ParseJWT(jwt)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(jwtDecoder)
}
