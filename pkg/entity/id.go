package entity

import "github.com/google/uuid"

type ID = uuid.UUID

// NewID creates a new ID
func NewID() ID {
	return ID(uuid.New())
}

// ParseID parses a string into an ID
func ParseID(id string) (ID, error) {
	parsedID, err := uuid.Parse(id)
	return ID(parsedID), err
}
