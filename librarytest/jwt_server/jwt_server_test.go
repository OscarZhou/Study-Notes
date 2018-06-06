package jwt_server

import (
	"io/ioutil"
	"log"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
)

func TestJWTServer(t *testing.T) {

	keyData, err := ioutil.ReadFile("private.pem")
	if err != nil {
		t.Error(err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(keyData))
	if err != nil {
		t.Error(err)
	}

	jwt, err := GenerateJWT(key)
	if err != nil {
		t.Error(err)

	}
	log.Println(jwt)
}
