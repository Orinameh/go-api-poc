package controllers

import (
	"encoding/json"
	"go-api/auth"
	"go-api/models"
	"go-api/utils"
	"io/ioutil"
	"net/http"
)

// Login function
func Login(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnauthorized)
		return
	}

	userAuth, err := auth.SignIn(user)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnauthorized)
		return
	}
	utils.ToJSON(w, userAuth)
}
