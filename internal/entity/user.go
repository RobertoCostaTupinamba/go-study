package entity

import (
	"github.com/RobertoCostaTupinamba/go-study/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user entity
type User struct {
	ID       entity.ID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

// NewUser creates a new user
func NewUser(username, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewID(),
		Username: username,
		Email:    email,
		Password: string(hash),
	}, nil
}

// ComparePassword compares the user password with the provided password
func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
