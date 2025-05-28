package service

import "online_library/backend/internal/models"

func (s *BookService) SearchBooks(tagIDs []int, limit, offset int, sort string) ([]models.Book, error) {
	if len(tagIDs) == 0 {
		return s.repo.GetAllBooks(limit, offset, sort)
	}
	return s.repo.GetBooksByTags(tagIDs, limit, offset, sort)
}
