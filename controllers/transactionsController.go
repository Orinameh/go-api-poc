package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-api/models"
	"go-api/utils"
	"go-api/validations"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Errors
var (
	ErrInvalidCash = errors.New("Insufficient money")
)

// PostTransaction is a controller function
func PostTransaction(w http.ResponseWriter, r *http.Request) {
	transaction, err := verifyTransaction(r)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	_, err = models.NewTransaction(transaction)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, utils.DefaultResponse{"Transaction successful", http.StatusCreated})
}

func verifyTransaction(r *http.Request) (models.Transaction, error) {
	// receives the public key from the portfolio that will receive the amount
	params := mux.Vars(r)
	targetKey := params["public_key"]

	target, err := models.GetWalletByPublicKey(targetKey)
	if err != nil {
		return models.Transaction{}, err
	}
	// receives the json from the wallet that will send a paid amount and the public key
	body, _ := ioutil.ReadAll(r.Body)
	var origin models.Wallet
	err = json.Unmarshal(body, &origin) //data sent is stored inside the origin pointer
	if err != nil {
		return models.Transaction{}, err
	}

	// Check whether the portfolio exists and return the value of the wallet from the db
	originVerify, err := models.GetWalletByPublicKey(origin.PublicKey)
	if err != nil {
		return models.Transaction{}, err
	}

	if validations.IsEmpty(target.PublicKey) || validations.IsEmpty(originVerify.PublicKey) {
		return models.Transaction{}, models.ErrWalletNotFound
	}
	// checks if the balance to be transferred is greater than the balance of the portfolio
	if origin.Balance > originVerify.Balance || origin.Balance < 0 {
		return models.Transaction{}, ErrInvalidCash
	}

	var transaction models.Transaction
	transaction.Cash = origin.Balance
	transaction.Message = fmt.Sprintf("%s transferred %.2f $, for  %s", originVerify.User.Nickname, origin.Balance, target.User.Nickname)
	transaction.Origin = origin
	transaction.Target = target
	return transaction, nil
}

// GetTransactions controller
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := models.GetTransactions()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	utils.ToJSON(w, transactions)
}
