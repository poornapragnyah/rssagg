package main

import (
	"github.com/google/uuid"
	"github.com/poornapragnyah/rssagg/internal/database"
)

// User represents a user in the system
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Name      string    `json:"name"`
}

func databaseUserToUser(u database.User) User {
	return User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
		Name:      u.Name,
	}
} 