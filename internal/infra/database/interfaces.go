package database

import "github.com/RobertoCostaTupinamba/go-study/internal/entity"

// UserInterface is an interface that defines the methods that a UserDatabase should implement
type UserInterface interface {
	CreateUser(user *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
	GetUserById(id string) (*entity.User, error)
}
