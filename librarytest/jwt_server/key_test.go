package jwt_server

import "testing"

func TestKey(t *testing.T) {
	err := GenerateKeys()
	if err != nil {
		t.Error(err)
	}
}
