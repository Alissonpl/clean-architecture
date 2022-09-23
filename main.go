package main

import (
	"clean-go/src/main/factories/application/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() {
	

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_DATABASE := os.Getenv("DB_DATABASE")

	dbURL := "postgres://" + DB_USERNAME + ":" + DB_PASSWORD + "@" + DB_HOST + ":" + DB_PORT + "/" + DB_DATABASE

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Error DB", err.Error())
	}

	println("DB", db)
}

func main() {
    err := godotenv.Load()
       
    if err != nil {
		log.Fatal("Error loading .env file")
	}

    initDB()
    
    r := mux.NewRouter()
    controllers.MakeUserHandlers(r)
    
    http.ListenAndServe(":8080", r)
}