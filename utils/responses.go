package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// DefaultResponse structure
type DefaultResponse struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

// ErrorResponse for display error
func ErrorResponse(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	ToJSON(w, struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	})
}

// ToJSON func
func ToJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}
