package service

import (
	"errors"
	"online_library/backend/internal/models"
	"online_library/backend/internal/pkg/translit"
	"online_library/backend/internal/repository"
)

type AuthorServiceInterface interface {
	CreateAuthor(author *models.Author) error
	UpdateAuthor(author *models.Author) error
	DeleteAuthor(id int) error
	GetAuthorByID(id int) (*models.Author, error)
	SearchAuthors(query string, limit, offset int) ([]*models.Author, int, error)
	GetAllAuthors(limit, offset int) ([]models.Author, error)
}

type AuthorService struct {
	repo repository.AuthorRepository
}

func NewAuthorService(repo repository.AuthorRepository) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) CreateAuthor(author *models.Author) error {

	if author.NameRU == "" && author.NameEN == "" {
		return errors.New("at least one of NameRU or NameEN must be provided")
	}

	if author.NameEN == "" && author.NameRU != "" {
		author.NameEN = translit.ToLatin(author.NameRU)
	}

	exists, err := s.repo.AuthorExists(author.NameRU, author.NameEN, 0)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("author with this name already exists")
	}

	return s.repo.CreateAuthor(author)
}

func (s *AuthorService) UpdateAuthor(author *models.Author) error {
	if author.NameRU == "" && author.NameEN == "" {
		return errors.New("at least one of NameRU or NameEN must be provided")
	}

	if author.NameEN == "" && author.NameRU != "" {
		author.NameEN = translit.ToLatin(author.NameRU)
	}

	exists, err := s.repo.AuthorExists(author.NameRU, author.NameEN, author.ID) // исключаем самого себя по ID
	if err != nil {
		return err
	}
	if exists {
		return errors.New("author with this name already exists")
	}

	return s.repo.UpdateAuthor(author)
}

func (s *AuthorService) DeleteAuthor(id int) error {
	return s.repo.DeleteAuthor(id)
}

func (s *AuthorService) GetAuthorByID(id int) (*models.Author, error) {
	return s.repo.GetAuthorByID(id)
}

func (s *AuthorService) SearchAuthors(query string, limit, offset int) ([]*models.Author, int, error) {
	authors, err := s.repo.SearchAuthorByName(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.repo.CountAuthors(query)
	if err != nil {
		return nil, 0, err
	}

	return authors, count, nil
}

func (s *AuthorService) GetAllAuthors(limit, offset int) ([]models.Author, error) {
	return s.repo.GetAllAuthors(limit, offset)
}
