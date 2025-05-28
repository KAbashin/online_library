package models

type Author struct {
	ID       int     `json:"id"`
	NameRU   string  `json:"name_ru"`
	NameEN   string  `json:"name_en"`
	Bio      *string `json:"bio,omitempty"`
	PhotoURL *string `json:"photo_url,omitempty"`
}
