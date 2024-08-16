package main

import (
	"encoding/json"
	"net/http"

	"github.com/RobertoCostaTupinamba/go-study/configs"
	"github.com/RobertoCostaTupinamba/go-study/internal/dto"
	"github.com/RobertoCostaTupinamba/go-study/internal/entity"
	"github.com/RobertoCostaTupinamba/go-study/internal/infra/database"
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
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	// Start the HTTP server
	http.ListenAndServe(":8000", nil)
}

type ProductHandler struct {
	ProductDB *database.ProductDatabase
}

func NewProductHandler(db *gorm.DB) *ProductHandler {
	return &ProductHandler{
		ProductDB: database.NewProductDatabase(db),
	}
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create a new Product entity instance.
	//TODO: implement useCase to create a product
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = ph.ProductDB.CreateProduct(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
