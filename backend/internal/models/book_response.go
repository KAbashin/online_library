package models

type BookResponse struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description *string  `json:"description"`
	PublishYear *int     `json:"publish_year,omitempty"`
	Pages       *int     `json:"pages,omitempty"`
	Language    *string  `json:"language,omitempty"`
	Publisher   *string  `json:"publisher,omitempty"`
	Type        *string  `json:"type,omitempty"`
	Rating      int      `json:"rating"`
	CoverURL    *string  `json:"cover_url,omitempty"`
	Status      string   `json:"status"`
	Authors     []Author `json:"authors"`
	Tags        []Tag    `json:"tags"`
	// IsFavorite  bool     `json:"is_favorite"`
}
