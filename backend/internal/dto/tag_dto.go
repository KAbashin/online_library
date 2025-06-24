package dto

import "online_library/backend/internal/models"

type TagDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color,omitempty"`
}

func ConvertTags(tags []models.Tag) []TagDTO {
	result := make([]TagDTO, 0, len(tags))
	for _, tag := range tags {
		result = append(result, TagDTO{
			ID:    tag.ID,
			Name:  tag.Name,
			Color: tag.Color,
		})
	}
	return result
}

func (t *TagDTO) ToModel() *models.Tag {
	return &models.Tag{
		ID:    t.ID,
		Name:  t.Name,
		Color: t.Color,
	}
}
