package models

import (
	"database/sql"
	"fmt"
	"go-api/config"
	"log"

	// This is a blank import for postgres driver
	_ "github.com/lib/pq"
)

// constants
// const (
// 	USER    = "davidevhade"
// 	PASS    = "0704502"
// 	DBNAME  = "blockcoin"
// 	SSLMODE = "disable"
// )
var configs = config.LoadConfig()

// Connect function that returns db connection url
func Connect() *sql.DB {
	URL := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", configs.Database.User, configs.Database.Pass, configs.Database.Name, "disable")
	db, err := sql.Open("postgres", URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

// TestConnection for testing the db connection
func TestConnection() {
	con := Connect()
	defer con.Close()
	err := con.Ping()
	if err != nil {
		fmt.Println(fmt.Errorf("%s", err.Error()))
		return
	}
	fmt.Println("Database connected")
}
