package controllers

import (
	"go-api/utils"
	"net/http"
)

// GetHome controller
func GetHome(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type", "application/json")
	// json.NewEncoder(w).Encode(struct {
	// 	Message string `json:"message"`
	// }{
	// 	Message: "Go RESTful Api",
	// })
	utils.ToJSON(w, struct {
		Message string `json:"message"`
	}{
		Message: "Go RESTful Api",
	})
}
