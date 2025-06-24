package dto

import (
	"online_library/backend/internal/models"
	"time"
)

type UserDTO struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	Name         *string   `json:"name,omitempty"`
	Role         string    `json:"role"`
	Bio          *string   `json:"bio,omitempty"`
	RegisteredAt time.Time `json:"registered_at"`
	// Можно добавить, например, поле активен ли пользователь
	IsActive bool `json:"is_active"`
}

func ConvertToUserDTO(u *models.User) UserDTO {
	return UserDTO{
		ID:           u.ID,
		Email:        u.Email,
		Name:         u.Name,
		Role:         u.Role,
		Bio:          u.Bio,
		RegisteredAt: u.RegisteredAt,
		IsActive:     u.Is_active,
	}
}

func ConvertToUserDTOs(users []models.User) []UserDTO {
	dtos := make([]UserDTO, len(users))
	for i, u := range users {
		dtos[i] = ConvertToUserDTO(&u)
	}
	return dtos
}
