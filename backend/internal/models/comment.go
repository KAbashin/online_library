package models

import (
	"time"
)

type Comment struct {
	ID        int       `json:"id"`
	BookID    int       `json:"book_id"`
	UserID    int       `json:"user_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"` // Добавляем для отслеживания изменений
	Status    string    `json:"status"`     // "active", "hidden", "deleted", "pending"...
}
