package database

import (
	"github.com/RobertoCostaTupinamba/go-study/internal/entity"
	"gorm.io/gorm"
)

type UserDatabase struct {
	DB *gorm.DB
}

// NewUser creates a new UserDatabase
func NewUserDatabase(db *gorm.DB) *UserDatabase {
	return &UserDatabase{
		DB: db,
	}
}

// CreateUser creates a new user
func (udb *UserDatabase) CreateUser(user *entity.User) error {
	return udb.DB.Create(user).Error
}

// GetUserByEmail returns a user by email
func (udb *UserDatabase) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := udb.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserById returns a user by id
func (udb *UserDatabase) GetUserById(id string) (*entity.User, error) {
	var user entity.User
	if err := udb.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
