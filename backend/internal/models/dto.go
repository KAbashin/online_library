package models

type CreateBookRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description,omitempty"`
	PublishYear *int    `json:"publish_year,omitempty"`
	Pages       *int    `json:"pages,omitempty"`
	Language    *string `json:"language,omitempty"`
	Publisher   *string `json:"publisher,omitempty"`
	Type        *string `json:"type,omitempty"`
	CoverURL    *string `json:"cover_url,omitempty"`
}

type BookResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	PublishYear *int    `json:"publish_year,omitempty"`
	Pages       *int    `json:"pages,omitempty"`
	Language    *string `json:"language,omitempty"`
	Publisher   *string `json:"publisher,omitempty"`
	Type        *string `json:"type,omitempty"`
	CoverURL    *string `json:"cover_url,omitempty"`
	CreatedAt   string  `json:"created_at"`
}
