package entity

import (
	"errors"
	"time"

	"github.com/RobertoCostaTupinamba/go-study/pkg/entity"
)

var (
	ErrorIDRequired    = errors.New("id is required")
	ErrorInvalidID     = errors.New("id is invalid")
	ErrorNameRequired  = errors.New("name is required")
	ErrorPriceRequired = errors.New("price is required")
	ErrorPriceInvalid  = errors.New("price is invalid")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	if err := product.Validate(); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrorIDRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrorInvalidID
	}
	if p.Name == "" {
		return ErrorNameRequired
	}
	if p.Price == 0 {
		return ErrorPriceRequired
	}
	if p.Price < 0 {
		return ErrorPriceInvalid
	}
	return nil
}
