package database

import "github.com/RobertoCostaTupinamba/go-study/internal/entity"

// UserInterface is an interface that defines the methods that a UserDatabase should implement
type UserInterface interface {
	CreateUser(user *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
	GetUserById(id string) (*entity.User, error)
}

// ProductInterface is an interface that defines the methods that a ProductDatabase should implement
type ProductInterface interface {
	CreateProduct(product *entity.Product) error
	FindAll(offset, limit int, sort string) ([]entity.Product, error)
	FindById(id string) (*entity.Product, error)
	UpdateProduct(product *entity.Product) error
	DeleteProduct(id string) error
}
