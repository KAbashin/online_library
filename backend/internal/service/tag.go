package service

import (
	"errors"
	"online_library/backend/internal/models"
	"online_library/backend/internal/repository"
)

type TagService interface {
	GetAllTags() ([]models.Tag, error)
	GetTagByID(id int) (models.Tag, error)

	CreateTag(tag *models.Tag) error
	UpdateTag(tag *models.Tag) error
	DeleteTag(id int) error

	GetTagsByBookID(bookID int) ([]models.Tag, error)
	AssignTagToBook(bookTag *models.BookTag) error
	RemoveTagFromBook(bookID, tagID int) error
}

type tagService struct {
	tagRepo repository.TagRepository
}

func NewTagService(tagRepo repository.TagRepository) TagService {
	return &tagService{tagRepo: tagRepo}
}

func (s *tagService) GetAllTags() ([]models.Tag, error) {
	return s.tagRepo.GetAllTags()
}

func (s *tagService) GetTagByID(id int) (models.Tag, error) {
	tag, err := s.tagRepo.GetTagByID(id)
	if err != nil {
		return models.Tag{}, err
	}
	return tag, nil
}

func (s *tagService) CreateTag(tag *models.Tag) error {
	if tag.Name == "" {
		return errors.New("tag name is required")
	}
	return s.tagRepo.CreateTag(tag)
}

func (s *tagService) UpdateTag(tag *models.Tag) error {
	if tag.ID == 0 {
		return errors.New("tag ID is required")
	}
	return s.tagRepo.UpdateTag(tag)
}

func (s *tagService) DeleteTag(id int) error {
	if id == 0 {
		return errors.New("tag ID is required")
	}
	return s.tagRepo.DeleteTag(id)
}

func (s *tagService) GetTagsByBookID(bookID int) ([]models.Tag, error) {
	return s.tagRepo.GetTagsByBookID(bookID)
}

func (s *tagService) AssignTagToBook(bt *models.BookTag) error {
	if bt.BookID == 0 || bt.TagID == 0 {
		return errors.New("book ID and tag ID are required")
	}
	return s.tagRepo.AssignTagToBook(bt)
}

func (s *tagService) RemoveTagFromBook(bookID, tagID int) error {
	if bookID == 0 || tagID == 0 {
		return errors.New("book ID and tag ID are required")
	}
	return s.tagRepo.RemoveTagFromBook(bookID, tagID)
}
