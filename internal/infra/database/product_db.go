package database

import (
	"strings"

	"github.com/RobertoCostaTupinamba/go-study/internal/entity"
	"gorm.io/gorm"
)

type ProductDatabase struct {
	DB *gorm.DB
}

func NewProductDatabase(db *gorm.DB) *ProductDatabase {
	return &ProductDatabase{
		DB: db,
	}
}

// CreateProduct creates a new product
func (p *ProductDatabase) CreateProduct(product *entity.Product) error {
	return p.DB.Create(product).Error
}

// FindById returns a product by its id
func (p *ProductDatabase) FindById(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, id).Error
	return &product, err
}

// UpdateProduct updates a product
func (p *ProductDatabase) UpdateProduct(product *entity.Product) error {
	// Validate if the product exists
	_, err := p.FindById(product.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(product).Error
}

// DeleteProduct deletes a product
func (p *ProductDatabase) DeleteProduct(id string) error {
	// Validate if the product exists
	_, err := p.FindById(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(&entity.Product{}, id).Error
}

// FindAll returns all products with support for multiple sort fields
func (p *ProductDatabase) FindAll(offset, limit int, sort string) ([]entity.Product, error) {
	// Validate offset: ensure it is non-negative
	if offset < 0 {
		offset = 0
	}
	// Validate limit: ensure it is non-negative
	if limit < 0 {
		limit = 0
	}

	// Define valid sort fields
	validSorts := map[string]string{
		"name":       "name",
		"price":      "price",
		"created_at": "created_at",
	}

	// Default sorting order if no valid sort is provided
	defaultSort := "created_at desc"

	var sortFields []string

	// Split the sort string by commas to allow multiple fields
	sortParts := strings.Split(sort, ",")
	for _, part := range sortParts {
		// Trim spaces and split by space to separate field and direction
		part = strings.TrimSpace(part)
		sortField := strings.Split(part, " ")

		field := sortField[0] // Extract the field name
		direction := "asc"    // Default to ascending order

		// If direction is provided, use it
		if len(sortField) > 1 {
			direction = sortField[1]
		}

		// Check if the field is valid
		if validField, exists := validSorts[field]; exists {
			// If valid, add the field and direction to the sortFields slice
			sortFields = append(sortFields, validField+" "+direction)
		}
	}

	// If no valid sort fields were provided, use the default sort
	if len(sortFields) == 0 {
		sortFields = append(sortFields, defaultSort)
	}

	// Join the sort fields with commas for the GORM Order clause
	finalSort := strings.Join(sortFields, ", ")

	var products []entity.Product
	// Query the database with pagination and sorting
	err := p.DB.Offset(offset).Limit(limit).Order(finalSort).Find(&products).Error
	if err != nil {
		// If there's an error, return it along with a nil product slice
		return nil, err
	}

	// Return the products and a nil error (indicating success)
	return products, nil
}
