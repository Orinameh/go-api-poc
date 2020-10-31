package routes

import (
	"go-api/controllers"
	"go-api/middleware"

	"github.com/gorilla/mux"
)

// NewRouter function that returns a router
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", controllers.GetHome).Methods("GET")
	r.HandleFunc("/users", middleware.IsAuth(controllers.GetUsers)).Methods("GET")
	r.HandleFunc("/users/{uid}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/users", controllers.PostUsers).Methods("POST")
	r.HandleFunc("/users/{uid}", controllers.PutUser).Methods("PUT")
	r.HandleFunc("/users/{uid}", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/wallets", controllers.GetWallets).Methods("GET")
	r.HandleFunc("/wallets/{public_key}", controllers.GetWallet).Methods("GET")
	r.HandleFunc("/wallets/{public_key}", controllers.PutWallet).Methods("PUT")
	r.HandleFunc("/transactions", controllers.GetTransactions).Methods("GET")
	r.HandleFunc("/transactions/{public_key}", controllers.PostTransaction).Methods("POST")

	r.HandleFunc("/login", controllers.Login).Methods("POST")

	return r
}
