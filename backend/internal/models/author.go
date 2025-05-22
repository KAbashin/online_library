package models

type Author struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Bio      *string `json:"bio,omitempty"`
	PhotoURL *string `json:"photo_url,omitempty"`
}
