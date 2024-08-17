package main

import (
	"net/http"

	"github.com/RobertoCostaTupinamba/go-study/configs"
	"github.com/RobertoCostaTupinamba/go-study/internal/entity"
	"github.com/RobertoCostaTupinamba/go-study/internal/infra/database"
	"github.com/RobertoCostaTupinamba/go-study/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Load the application configuration
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	// Open a database connection
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Migrate the schema
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	// Create a new ProductDatabase instance
	productDB := database.NewProductDatabase(db)
	// Create a new ProductHandler instance
	productHandler := handlers.NewProductHandler(productDB)

	// Register the handler function
	http.HandleFunc("/products", productHandler.CreateProduct)
	// Start the HTTP server
	http.ListenAndServe(":8000", nil)
}
