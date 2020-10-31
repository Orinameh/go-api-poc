package middleware

import (
	"fmt"
	"go-api/config"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var configs = config.LoadConfig()

// IsAuth function checks if a user is authenticated when performing an action
func IsAuth(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header != "" {

			bearerToken := strings.Split(header, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					_, ok := token.Method.(*jwt.SigningMethodHMAC)
					if !ok {
						return nil, fmt.Errorf("Failure to Authenticate")
					}
					return configs.Jwt.SecretKey, nil
				})
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Unauthorized"))
					return
				}
				if token.Valid {
					handler(w, r)

				}
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized")
		}

	})
}
