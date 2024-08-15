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
	err := p.DB.First(&product, "id = ?", id).Error
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
	return p.DB.Delete(&entity.Product{}, "id = ?", id).Error
}

// FindAll returns all products with support for multiple sort fields and pagination by page
func (p *ProductDatabase) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	// Validate page: ensure it is non-negative and start from 1
	if page < 1 {
		page = 1
	}
	// Validate limit: ensure it is non-negative
	if limit < 1 {
		limit = 10 // Default limit if not provided or invalid
	}

	// Calculate offset based on page and limit
	offset := (page - 1) * limit

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
