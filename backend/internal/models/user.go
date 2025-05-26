package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Name         *string   `json:"name,omitempty"`
	Role         string    `json:"role"`
	Bio          *string   `json:"bio,omitempty"`
	RegisteredAt time.Time `json:"registered_at"`
	TokenVersion int       `json:"-"`
	Is_active    bool      `json:"is_active"`
}

type UserInput struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Bio      string `json:"bio"`
	Password string `json:"password"`
}

type AdminUserUpdateInput struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Bio   string `json:"bio"`
	Role  string `json:"role"` // user, admin, superadmin
}
