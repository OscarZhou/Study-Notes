package parseJWKS

import (
	"io/ioutil"
	"testing"
)

func TestParseJWKS(t *testing.T) {
	jwksURL := "https://xxxx.auth0.com/.well-known/jwks.json"
	expected, err := ParseJWKS(jwksURL, "RS256")
	if err != nil {
		t.Error(err)
	}

	actual, err := ioutil.ReadFile("jwt_certificate.txt")
	if err != nil {
		t.Error(err)
	}

	if len(expected) != len(actual) {
		t.Errorf("expected length is %d, actual length is %d", len(expected), len(actual))
	}

	e := []byte(expected)
	for i := 0; i < len(expected); i++ {
		if e[i] != actual[i] {
			t.Errorf("%d: expected is %v, actual is %v", i, e[i], actual[i])
		}
	}

	if expected != string(actual) {
		t.Errorf("expected:%s\nactual:%s", expected, string(actual))
	}
}
