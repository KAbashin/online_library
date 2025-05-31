package service

import (
	"online_library/backend/internal/models"
	"online_library/backend/internal/repository"
)

type CategoryService interface {
	GetCategoryTree() ([]map[string]interface{}, error)
	GetCategoryRoot() ([]*models.Category, error)
	GetCategoryByID(id int) (*models.Category, error)
	GetCategoryChildren(id int) ([]*models.Category, error)
	GetBooksByCategoryIDRecursive(id int) ([]*models.Book, error)
	CreateCategory(category *models.Category) (int, error)
	UpdateCategory(category *models.Category) error
	DeleteCategory(id int) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) GetCategoryTree() ([]map[string]interface{}, error) {
	categories, err := s.repo.GetAllCategories()
	if err != nil {
		return nil, err
	}
	return buildCategoryTree(categories, nil), nil
}

func buildCategoryTree(categories []models.Category, parentID *int) []map[string]interface{} {
	var tree []map[string]interface{}
	for _, cat := range categories {
		if (cat.ParentID == nil && parentID == nil) || (cat.ParentID != nil && parentID != nil && *cat.ParentID == *parentID) {
			children := buildCategoryTree(categories, &cat.ID)
			node := map[string]interface{}{
				"id":          cat.ID,
				"name":        cat.Name,
				"slug":        cat.Slug,
				"children":    children,
				"description": cat.Description,
			}
			tree = append(tree, node)
		}
	}
	return tree
}

func (s *categoryService) GetCategoryRoot() ([]*models.Category, error) {
	return s.repo.GetRootCategories()
}

func (s *categoryService) GetCategoryByID(id int) (*models.Category, error) {
	return s.repo.GetCategoryByID(id)
}

func (s *categoryService) GetCategoryChildren(id int) ([]*models.Category, error) {
	return s.repo.GetCategoryChildren(id)
}

func (s *categoryService) GetBooksByCategoryIDRecursive(categoryID int) ([]*models.Book, error) {
	return s.repo.GetBooksByCategoryIDRecursive(categoryID)
}

func (s *categoryService) CreateCategory(category *models.Category) (int, error) {
	return s.repo.CreateCategory(category)
}

func (s *categoryService) UpdateCategory(category *models.Category) error {
	return s.repo.UpdateCategory(category)
}

func (s *categoryService) DeleteCategory(id int) error {
	return s.repo.DeleteCategory(id)
}
