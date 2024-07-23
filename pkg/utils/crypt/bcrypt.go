package crypt

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Encode(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		err = fmt.Errorf("error while hashing password: %s", password)
	}

	return bytes, err
}

func Matches(password []byte, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
