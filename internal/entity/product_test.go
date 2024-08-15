package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewProduct tests the NewProduct function
func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 10.5)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 10.5, product.Price)
}

// Test when the product name is empty
func TestProduct_Validate_EmptyName(t *testing.T) {
	product, err := NewProduct("", 10.5)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrorNameRequired, err)
}

// Test when the product price is zero
func TestProduct_Validate_ZeroPrice(t *testing.T) {
	product, err := NewProduct("Product 1", 0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrorPriceRequired, err)
}

// Test when the product price is negative
func TestProduct_Validate_NegativePrice(t *testing.T) {
	product, err := NewProduct("Product 1", -10.5)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrorPriceInvalid, err)
}
