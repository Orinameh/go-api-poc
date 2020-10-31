package validations

import (
	"errors"
	"go-api/models"
)

// Global Variables
var (
	ErrEmptyFields  = errors.New("One or more fields are empty")
	ErrInvalidEmail = errors.New("Email is invalid")
)

// ValidateNewUser function to validate user fields
func ValidateNewUser(user models.User) (models.User, error) {
	if IsEmpty(user.Nickname) || IsEmpty(user.Email) || IsEmpty(user.Password) {
		return models.User{}, ErrEmptyFields
	}
	if !IsEmail(user.Email) {
		return models.User{}, ErrInvalidEmail
	}
	return user, nil
}
