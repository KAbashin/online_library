package dto

import "online_library/backend/internal/models"

type CategoryDTO struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Slug     *string `json:"slug,omitempty"`
	ParentID *int    `json:"parent_id,omitempty"`
}

type BreadcrumbCategoryDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
type CreateCategoryDTO struct {
	Name        string  `json:"name" binding:"required"`
	ParentID    *int    `json:"parent_id,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (dto CreateCategoryDTO) ToModel() *models.Category {
	return &models.Category{
		Name:        dto.Name,
		ParentID:    dto.ParentID,
		Slug:        dto.Slug,
		Description: dto.Description,
	}
}

func ConvertCategory(cat models.Category) CategoryDTO {
	return CategoryDTO{
		ID:       cat.ID,
		Name:     cat.Name,
		Slug:     cat.Slug,
		ParentID: cat.ParentID,
	}
}

func ConvertCategories(cats []models.Category) []CategoryDTO {
	result := make([]CategoryDTO, 0, len(cats))
	for _, c := range cats {
		result = append(result, ConvertCategory(c))
	}
	return result
}

func BuildBreadcrumbs(startID int, fetch func(id int) (models.Category, error)) ([]BreadcrumbCategoryDTO, error) {
	var breadcrumbs []BreadcrumbCategoryDTO

	currentID := startID
	for {
		category, err := fetch(currentID)
		if err != nil {
			return nil, err
		}

		bdto := BreadcrumbCategoryDTO{
			ID:   category.ID,
			Name: category.Name,
		}
		if category.Slug != nil {
			bdto.Slug = *category.Slug
		}

		breadcrumbs = append([]BreadcrumbCategoryDTO{bdto}, breadcrumbs...) // prepend

		if category.ParentID == nil {
			break
		}
		currentID = *category.ParentID
	}

	return breadcrumbs, nil
}
