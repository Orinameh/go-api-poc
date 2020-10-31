package validations

import "github.com/badoux/checkmail"

// IsEmpty function
func IsEmpty(param string) bool {
	if param == "" {
		return true
	}
	return false
}

// IsEmail function
func IsEmail(email string) bool {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return false
	}
	return true
}
