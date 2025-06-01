package service

import (
	"fmt"
	"online_library/backend/internal/models"
	"online_library/backend/internal/pkg/roles"
	"online_library/backend/internal/repository"
)

type BookService interface {
	CreateBook(book *models.Book, userRole string, userID int) (int, error)
	UpdateBook(book *models.Book, userID int, userRole string) error
	DeleteBook(bookID int, userID int, userRole string) error
	GetBookByID(bookID int, userRole string) (*models.Book, error)
	GetBooksByStatuses(userRole string, offset, limit int) ([]models.Book, error)
	GetBooksByAuthor(authorID int, userRole string, offset, limit int) ([]models.Book, error)
	GetBooksByTag(tagID int, userRole string, offset, limit int) ([]models.Book, error)
	SetBookAuthors(bookID int, authorIDs []int, userID int, userRole string) error
	AddBookAuthor(bookID, authorID int, userID int, userRole string) error
	RemoveBookAuthor(bookID, authorID int, userID int, userRole string) error
	SetBookTags(bookID int, tagIDs []int, userID int, userRole string) error
	AddBookTag(bookID, tagID int, userID int, userRole string) error
	RemoveBookTag(bookID, tagID int, userID int, userRole string) error
	UpdateBookStatus(bookID int, status string, userRole string) error
	SearchBooks(query string, userRole string, limit, offset int) ([]*models.Book, error)
	GetDuplicateBooks(title string) ([]*models.Book, error)
	GetUserBooks(userID int) ([]*models.Book, error)
	GetUserFavoriteBooks(userID int, userRole string) ([]*models.Book, error)
	AddBookToFavorites(userID, bookID int) error
	RemoveBookFromFavorites(userID, bookID int) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func getViewableStatuses(userRole string) []string {
	if roles.IsAdmin(userRole) {
		return []string{
			models.StatusBookVisible,
			models.StatusBookArchived,
			models.StatusBookQuarantine,
			models.StatusBookPrivate,
		}
	}

	if userRole == roles.RoleUser {
		return []string{
			models.StatusBookVisible,
			models.StatusBookQuarantine,
		}
	}

	if userRole == roles.RoleNewUser {
		return []string{}
	}

	// Для всех остальных
	return []string{models.StatusBookVisible}
}

func (s *bookService) checkBookOwnership(bookID, userID int, userRole string) error {
	if roles.IsAdmin(userRole) {
		return nil
	}

	book, err := s.repo.GetBookMeta(bookID)
	if err != nil {
		return fmt.Errorf("failed to check ownership: %w", err)
	}

	if book.CreatedBy != userID {
		return fmt.Errorf("permission denied: user is not the creator of the book")
	}

	return nil
}

func (s *bookService) CreateBook(book *models.Book, userRole string, userID int) (int, error) {
	if roles.IsAdmin(userRole) {
		book.Status = models.StatusBookVisible
	} else {
		book.Status = models.StatusBookQuarantine
	}
	book.CreatedBy = userID
	return s.repo.CreateBook(book)
}

func (s *bookService) UpdateBook(book *models.Book, userID int, userRole string) error {
	if err := s.checkBookOwnership(book.ID, userID, userRole); err != nil {
		return err
	}
	return s.repo.UpdateBook(book)
}

func (s *bookService) DeleteBook(bookID int, userID int, userRole string) error {
	if err := s.checkBookOwnership(bookID, userID, userRole); err != nil {
		return err
	}
	return s.repo.DeleteBook(bookID)
}

func (s *bookService) GetBookByID(bookID int, userRole string) (*models.Book, error) {
	statuses := getViewableStatuses(userRole)
	return s.repo.GetBookByID(bookID, statuses)
}

func (s *bookService) GetBooksByStatuses(userRole string, offset, limit int) ([]models.Book, error) {
	statuses := getViewableStatuses(userRole)
	return s.repo.GetBooksByStatuses(statuses, offset, limit)
}

func (s *bookService) GetBooksByAuthor(authorID int, userRole string, offset, limit int) ([]models.Book, error) {
	statuses := getViewableStatuses(userRole)
	return s.repo.GetBooksByAuthor(authorID, statuses, limit, offset)
}

func (s *bookService) GetBooksByTag(tagID int, userRole string, offset, limit int) ([]models.Book, error) {
	statuses := getViewableStatuses(userRole)
	return s.repo.GetBooksByTag(tagID, statuses, limit, offset)
}

func (s *bookService) SetBookAuthors(bookID int, authorIDs []int, userID int, userRole string) error {
	if err := s.checkBookOwnership(bookID, userID, userRole); err != nil {
		return err
	}
	return s.repo.SetBookAuthors(bookID, authorIDs)
}

func (s *bookService) AddBookAuthor(bookID, authorID int, userID int, userRole string) error {
	if err := s.checkBookOwnership(bookID, userID, userRole); err != nil {
		return err
	}
	return s.repo.AddBookAuthor(bookID, authorID)
}

func (s *bookService) RemoveBookAuthor(bookID, authorID int, userID int, userRole string) error {
	if err := s.checkBookOwnership(bookID, userID, userRole); err != nil {
		return err
	}
	return s.repo.RemoveBookAuthor(bookID, authorID)
}

func (s *bookService) SetBookTags(bookID int, tagIDs []int, userID int, userRole string) error {
	if err := s.checkBookOwnership(bookID, userID, userRole); err != nil {
		return err
	}
	return s.repo.SetBookTags(bookID, tagIDs)
}

func (s *bookService) AddBookTag(bookID, tagID int, userID int, userRole string) error {
	if err := s.checkBookOwnership(bookID, userID, userRole); err != nil {
		return err
	}
	return s.repo.AddBookTag(bookID, tagID)
}

func (s *bookService) RemoveBookTag(bookID, tagID int, userID int, userRole string) error {
	if err := s.checkBookOwnership(bookID, userID, userRole); err != nil {
		return err
	}
	return s.repo.RemoveBookTag(bookID, tagID)
}

func (s *bookService) UpdateBookStatus(bookID int, status string, userRole string) error {
	if !roles.IsAdmin(userRole) {
		return fmt.Errorf("permission denied: only admin can update status")
	}
	return s.repo.UpdateBookStatus(bookID, status)
}

func (s *bookService) SearchBooks(query string, userRole string, limit, offset int) ([]*models.Book, error) {
	statuses := getViewableStatuses(userRole)
	return s.repo.SearchBooks(query, statuses, limit, offset)
}

func (s *bookService) GetDuplicateBooks(title string) ([]*models.Book, error) {
	return s.repo.GetDuplicateBooks(title)
}

func (s *bookService) GetUserBooks(userID int) ([]*models.Book, error) {
	return s.repo.GetUserBooks(userID)
}

func (s *bookService) GetUserFavoriteBooks(userID int, userRole string) ([]*models.Book, error) {
	statuses := getViewableStatuses(userRole)
	return s.repo.GetUserFavoriteBooks(userID, statuses)
}

func (s *bookService) AddBookToFavorites(userID, bookID int) error {
	return s.repo.AddBookToFavorites(userID, bookID)
}

func (s *bookService) RemoveBookFromFavorites(userID, bookID int) error {
	return s.repo.RemoveBookFromFavorites(userID, bookID)
}
