package main

import "testing"

func TestSha512Salt(t *testing.T) {
	password := `_Bvav&+^'LL%?'+\`

	salt, hashedPassword := HashAndSalt(password, 16)

	err := CompareHashedPassword(password, salt, hashedPassword)
	if err != nil {
		t.Error(err)
	}

	t.Log("testing successfully!")
}
