package parseJWKS

import (
	"fmt"
	"testing"
)

func TestParseJWKS(t *testing.T) {
	jwksURL := "https://oscarhealth.au.auth0.com/.well-known/jwks.json"
	s, err := ParseJWKS(jwksURL, "RS256")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(s)
}
