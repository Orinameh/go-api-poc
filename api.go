package main

import (
	"fmt"
	"go-api/models"
	"go-api/routes"
	"log"
	"net/http"
)

func main() {
	port := "8000"
	models.TestConnection()
	fmt.Printf("Api running on port %s\n", port)
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
