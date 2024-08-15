package database

import (
	"testing"

	"github.com/RobertoCostaTupinamba/go-study/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestProductDatabase_CreateProduct(t *testing.T) {
	// Open an in-memory SQLite database connection
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err) // Fail the test if there's an error opening the database
	}

	// Automatically migrate the schema to keep it up to date
	db.AutoMigrate(&entity.Product{})

	// Create a new product entity instance
	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)

	// Create a new ProductDatabase instance using the in-memory database
	productDB := NewProductDatabase(db)

	// Call the CreateProduct method to insert the product into the database
	err = productDB.CreateProduct(product)
	assert.Nil(t, err) // Assert that there were no errors during product creation

	// Declare a variable to hold the retrieved product
	var productFound entity.Product

	// Query the database to find the product by ID
	err = db.First(&productFound, "id = ?", product.ID).Error
	assert.Nil(t, err)                                 // Assert that there were no errors during retrieval
	assert.Equal(t, product.ID, productFound.ID)       // Assert that the retrieved product's ID matches the original
	assert.Equal(t, product.Name, productFound.Name)   // Assert that the names match
	assert.Equal(t, product.Price, productFound.Price) // Assert that the prices match
}

func TestProductDatabase_FindById(t *testing.T) {
	// Open an in-memory SQLite database connection
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err) // Fail the test if there's an error opening the database
	}

	// Automatically migrate the schema to keep it up to date
	db.AutoMigrate(&entity.Product{})

	// Create a new product entity instance
	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)

	// Create a new ProductDatabase instance using the in-memory database
	productDB := NewProductDatabase(db)

	// Call the CreateProduct method to insert the product into the database
	err = productDB.CreateProduct(product)
	assert.Nil(t, err) // Assert that there were no errors during product creation

	// Call the FindById method to retrieve the product by ID
	productFound, err := productDB.FindById(product.ID.String())
	assert.Nil(t, err)                                 // Assert that there were no errors during retrieval
	assert.Equal(t, product.ID, productFound.ID)       // Assert that the retrieved product's ID matches the original
	assert.Equal(t, product.Name, productFound.Name)   // Assert that the names match
	assert.Equal(t, product.Price, productFound.Price) // Assert that the prices match
}

// TestProductDatabase_UpdateProduct tests the UpdateProduct method of the ProductDatabase
func TestProductDatabase_UpdateProduct(t *testing.T) {
	// Open an in-memory SQLite database connection
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err) // Fail the test if there's an error opening the database
	}

	// Automatically migrate the schema to keep it up to date
	db.AutoMigrate(&entity.Product{})

	// Create a new product entity instance
	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)

	// Create a new ProductDatabase instance using the in-memory database
	productDB := NewProductDatabase(db)

	// Call the CreateProduct method to insert the product into the database
	err = productDB.CreateProduct(product)
	assert.Nil(t, err) // Assert that there were no errors during product creation

	// Update the product's name and price
	product.Name = "Product 2"
	product.Price = 20.0

	// Call the UpdateProduct method to update the product in the database
	err = productDB.UpdateProduct(product)
	assert.Nil(t, err) // Assert that there were no errors during product update

	// Declare a variable to hold the retrieved product
	var productFound entity.Product

	// Query the database to find the product by ID
	err = db.First(&productFound, "id = ?", product.ID).Error
	assert.Nil(t, err)                                 // Assert that there were no errors during retrieval
	assert.Equal(t, product.ID, productFound.ID)       // Assert that the retrieved product's ID matches the original
	assert.Equal(t, product.Name, productFound.Name)   // Assert that the names match
	assert.Equal(t, product.Price, productFound.Price) // Assert that the prices match
}

// TestProductDatabase_DeleteProduct tests the DeleteProduct method of the ProductDatabase
func TestProductDatabase_DeleteProduct(t *testing.T) {
	// Open an in-memory SQLite database connection
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err) // Fail the test if there's an error opening the database
	}

	// Automatically migrate the schema to keep it up to date
	db.AutoMigrate(&entity.Product{})

	// Create a new product entity instance
	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)

	// Create a new ProductDatabase instance using the in-memory database
	productDB := NewProductDatabase(db)

	// Call the CreateProduct method to insert the product into the database
	err = productDB.CreateProduct(product)
	assert.Nil(t, err) // Assert that there were no errors during product creation

	// Call the DeleteProduct method to delete the product from the database
	err = productDB.DeleteProduct(product.ID.String())
	assert.Nil(t, err) // Assert that there were no errors during product deletion

	// Declare a variable to hold the retrieved product
	var productFound entity.Product

	// Query the database to find the product by ID
	err = db.First(&productFound, "id = ?", product.ID).Error
	assert.NotNil(t, err)                        // Assert that there was an error during retrieval
	assert.Equal(t, gorm.ErrRecordNotFound, err) // Assert that the error is a record not found error
}

// TestProductDatabase_FindAll tests the FindAll method of the ProductDatabase
func TestProductDatabase_FindAll(t *testing.T) {
	// Open an in-memory SQLite database connection
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err) // Fail the test if there's an error opening the database
	}

	// Automatically migrate the schema to keep it up to date
	db.AutoMigrate(&entity.Product{})

	// Create a new product entity instance
	product1, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)

	// Create a new product entity instance
	product2, err := entity.NewProduct("Product 2", 20.0)
	assert.NoError(t, err)

	// Create a new ProductDatabase instance using the in-memory database
	productDB := NewProductDatabase(db)

	// Call the CreateProduct method to insert the products into the database
	err = productDB.CreateProduct(product1)
	assert.Nil(t, err) // Assert that there were no errors during product creation

	// Call the CreateProduct method to insert the products into the database
	err = productDB.CreateProduct(product2)
	assert.Nil(t, err) // Assert that there were no errors during product creation

	// Call the FindAll method to retrieve all products from the database
	products, err := productDB.FindAll(1, 0, "name")
	assert.Nil(t, err)         // Assert that there were no errors during retrieval
	assert.Len(t, products, 2) // Assert that there are two products in the database
}
