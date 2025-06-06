package models

type Category struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	ParentID    *int    `json:"parent_id,omitempty"` // nil для корня
	Slug        *string `json:"slug,omitempty"`
	Description *string `json:"description,omitempty"`
}
