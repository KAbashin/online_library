package models

type User struct {
	ID           int     `json:"id"`
	Email        string  `json:"email"`
	PasswordHash string  `json:"-"`
	Name         *string `json:"name,omitempty"`
	Role         string  `json:"role"`
}
