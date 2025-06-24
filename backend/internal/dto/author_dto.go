package dto

import "online_library/backend/internal/models"

type AuthorDTO struct {
	ID       int     `json:"id"`
	NameRU   string  `json:"name_ru"`
	NameEN   string  `json:"name_en"`
	Bio      *string `json:"bio,omitempty"`
	PhotoURL *string `json:"photo_url,omitempty"`
}

// Преобразование одного автора
func ConvertAuthor(a models.Author) AuthorDTO {
	return AuthorDTO{
		ID:       a.ID,
		NameRU:   a.NameRU,
		NameEN:   a.NameEN,
		Bio:      a.Bio,
		PhotoURL: a.PhotoURL,
	}
}

// Преобразование списка авторов
func ConvertAuthors(authors []models.Author) []AuthorDTO {
	result := make([]AuthorDTO, 0, len(authors))
	for _, a := range authors {
		result = append(result, ConvertAuthor(a))
	}
	return result
}
