package auth

import (
	"errors"
	"go-api/config"
	"go-api/models"
	"go-api/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Errors
var (
	ErrInvalidPassword = errors.New("Invalid email or password")
)

// Auth structure
type Auth struct {
	User    models.User `json:"user"`
	Token   string      `json:"token"`
	IsValid bool        `json:"is_valid"`
}

var configs = config.LoadConfig()

// SignIn function
func SignIn(user models.User) (Auth, error) {
	password := user.Password
	user, err := models.GetUserByEmail(user.Email)
	if err != nil {
		return Auth{IsValid: false}, err
	}

	err = utils.IsPassword(user.Password, password)
	if err != nil {
		return Auth{IsValid: false}, ErrInvalidPassword
	}
	token, err := GenerateJWT(user)
	if err != nil {
		return Auth{IsValid: false}, err
	}
	return Auth{user, token, true}, nil
}

// GenerateJWT function
func GenerateJWT(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["userId"] = user.UID
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	return token.SignedString(configs.Jwt.SecretKey)

}
