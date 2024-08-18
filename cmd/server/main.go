package main

import (
	"net/http"

	"github.com/RobertoCostaTupinamba/go-study/configs"
	"github.com/RobertoCostaTupinamba/go-study/internal/entity"
	"github.com/RobertoCostaTupinamba/go-study/internal/infra/database"
	"github.com/RobertoCostaTupinamba/go-study/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	// Create a new router
	r := chi.NewRouter()
	// Use the logger middleware
	r.Use(middleware.Logger)
	// Register the handler function
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products", productHandler.FindAllProducts)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	// Start the HTTP server
	http.ListenAndServe(":8000", r)
}
