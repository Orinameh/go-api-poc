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
	r.HandleFunc("/users/{uid}", middleware.IsAuth(controllers.GetUser)).Methods("GET")
	r.HandleFunc("/users", middleware.IsAuth(controllers.PostUsers)).Methods("POST")
	r.HandleFunc("/users/{uid}", middleware.IsAuth(controllers.PutUser)).Methods("PUT")
	r.HandleFunc("/users/{uid}", middleware.IsAuth(controllers.DeleteUser)).Methods("DELETE")
	r.HandleFunc("/wallets", middleware.IsAuth(controllers.GetWallets)).Methods("GET")
	r.HandleFunc("/wallets/{public_key}", middleware.IsAuth(controllers.GetWallet)).Methods("GET")
	r.HandleFunc("/wallets/{public_key}", middleware.IsAuth(controllers.PutWallet)).Methods("PUT")
	r.HandleFunc("/transactions", middleware.IsAuth(controllers.GetTransactions)).Methods("GET")
	r.HandleFunc("/transactions/{public_key}", middleware.IsAuth(controllers.PostTransaction)).Methods("POST")

	r.HandleFunc("/login", controllers.Login).Methods("POST")

	return r
}
