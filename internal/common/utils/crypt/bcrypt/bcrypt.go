package bcrypt

import "golang.org/x/crypto/bcrypt"

func Encode(password []byte) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return string(bytes), err
}

func Matches(password []byte, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
