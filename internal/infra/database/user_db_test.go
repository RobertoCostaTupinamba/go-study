package database

import (
	"testing"

	"github.com/RobertoCostaTupinamba/go-study/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TestUserDatabase_CreateUser tests the CreateUser method
func TestUserDatabase_CreateUser(t *testing.T) {
	// Open an in-memory SQLite database connection
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err) // Fail the test if there's an error opening the database
	}

	// Automatically migrate the schema to keep it up to date
	db.AutoMigrate(&entity.User{})

	// Create a new user entity instance
	user, _ := entity.NewUser("Jhon Doe", "j@j.com", "123456")

	// Create a new UserDatabase instance using the in-memory database
	userDB := NewUserDatabase(db)

	// Call the CreateUser method to insert the user into the database
	err = userDB.CreateUser(user)
	assert.Nil(t, err) // Assert that there were no errors during user creation

	// Declare a variable to hold the retrieved user
	var userFound entity.User

	// Query the database to find the user by ID
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)                                 // Assert that there were no errors during retrieval
	assert.Equal(t, user.ID, userFound.ID)             // Assert that the retrieved user's ID matches the original
	assert.Equal(t, user.Username, userFound.Username) // Assert that the usernames match
	assert.Equal(t, user.Email, userFound.Email)       // Assert that the emails match

	// Validate that the retrieved user's password matches the original password
	assert.True(t, user.ValidatePassword("123456"))
}

// TestUserDatabase_GetUserByEmail tests the GetUserByEmail method
func TestUserDatabase_GetUserByEmail(t *testing.T) {
	// Open an in-memory SQLite database connection
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err) // Fail the test if there's an error opening the database
	}

	// Automatically migrate the schema to keep it up to date
	db.AutoMigrate(&entity.User{})

	// Create a new user entity instance
	user, _ := entity.NewUser("Jhon Doe", "j@j.com", "123456")

	// Create a new UserDatabase instance using the in-memory database
	userDB := NewUserDatabase(db)

	// Call the CreateUser method to insert the user into the database
	err = userDB.CreateUser(user)
	assert.Nil(t, err) // Assert that there were no errors during user creation

	// Retrieve the user by email using the GetUserByEmail method
	userFound, err := userDB.GetUserByEmail("j@j.com")
	assert.Nil(t, err)                                 // Assert that there were no errors during retrieval
	assert.Equal(t, user.ID, userFound.ID)             // Assert that the retrieved user's ID matches the original
	assert.Equal(t, user.Username, userFound.Username) // Assert that the usernames match
	assert.Equal(t, user.Email, userFound.Email)       // Assert that the emails match
}

// TestUserDatabase_GetUserById tests the GetUserById method
func TestUserDatabase_GetUserById(t *testing.T) {
	// Open an in-memory SQLite database connection
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err) // Fail the test if there's an error opening the database
	}

	// Automatically migrate the schema to keep it up to date
	db.AutoMigrate(&entity.User{})

	// Create a new user entity instance
	user, _ := entity.NewUser("Jhon Doe", "j@j.com", "123456")

	// Create a new UserDatabase instance using the in-memory database
	userDB := NewUserDatabase(db)

	// Call the CreateUser method to insert the user into the database
	err = userDB.CreateUser(user)
	assert.Nil(t, err) // Assert that there were no errors during user creation

	// Retrieve the user by ID using the GetUserById method
	userFound, err := userDB.GetUserById(user.ID.String())
	assert.Nil(t, err)                                 // Assert that there were no errors during retrieval
	assert.Equal(t, user.ID, userFound.ID)             // Assert that the retrieved user's ID matches the original
	assert.Equal(t, user.Username, userFound.Username) // Assert that the usernames match
	assert.Equal(t, user.Email, userFound.Email)       // Assert that the emails match
}
