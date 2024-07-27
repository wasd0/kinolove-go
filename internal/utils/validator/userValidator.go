package validator

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
