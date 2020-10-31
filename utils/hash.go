package utils

import (
	"crypto/md5"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Md5 function for creating a hash
func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

// Bcrypt function for hashing password
func Bcrypt(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// IsPassword function
func IsPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
