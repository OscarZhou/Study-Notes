package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"math/rand"
	"time"
)

//
func HashAndSalt(password string, n int) (salt, hashedPassword string) {
	salt = RandString(n)
	h := hmac.New(sha512.New, []byte(salt))
	h.Write([]byte(password))
	hashedPassword = hex.EncodeToString(h.Sum(nil))
	return salt, hashedPassword
}

func CompareHashedPassword(password, salt, hashedPassword string) error {
	hmac := hmac.New(sha512.New, []byte(salt))
	hmac.Write([]byte(password))
	newPassword := hex.EncodeToString(hmac.Sum(nil))
	if newPassword == hashedPassword {
		return nil
	}
	return errors.New("password is not correct")
}

func RandString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
