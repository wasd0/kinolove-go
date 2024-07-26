package validator

import "kinolove/pkg/utils/crypt"

const (
	minUsernameLen = 5
	minPasswordLen = 8
)

func ValidateUser(username string, password string) bool {
	return ValidateUsername(username) && ValidatePassword(password)
}
func ValidatePassword(password string) bool {
	return len(password) >= minPasswordLen
}

func ValidateUsername(username string) bool {
	return len(username) >= minUsernameLen
}

func IsPasswordsMatches(password string, hash []byte) bool {
	pasHash, err := crypt.Encode(password)

	if err != nil {
		return false
	}

	return crypt.Matches(pasHash, hash)
}
