package models

import "time"

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"`
	PublishYear *int      `json:"publish_year,omitempty"`
	Pages       *int      `json:"pages,omitempty"`
	Language    *string   `json:"language,omitempty"`
	Publisher   *string   `json:"publisher,omitempty"`
	Type        *string   `json:"type,omitempty"`
	Rating      int       `json:"rating"`
	CoverURL    *string   `json:"cover_url,omitempty"`
	Status      string    `json:"status"` // "visible", "archived", "quarantine", "adult"
	CreatedBy   int       `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}
