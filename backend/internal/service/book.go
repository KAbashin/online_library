package service

import (
	"fmt"
	"online_library/backend/internal/dto"
	"online_library/backend/internal/models"
	"online_library/backend/internal/pkg/roles"
	"online_library/backend/internal/repository"
)

type BookService interface {
	CreateBook(book *models.Book, userRole string, userID int) (int, error)
	UpdateBook(book *models.Book, userID int, userRole string) error
	DeleteBook(bookID int, userID int, userRole string) error

	GetBookDTO(bookID int, userID int, userRole string) (*dto.BookDTO, error)
	GetBookExtrasDTO(bookID int, userID int, userRole string) (*dto.BookExtrasDTO, error)

	//GetBookByID(bookID int, userID int, userRole string) (*models.BookResponse, error)
	GetBooksByStatuses(userRole string, offset, limit int) ([]models.Book, error)
	GetBooksByAuthor(authorID int, userRole string, offset, limit int) ([]dto.BookPreviewDTO, error)
	GetBooksByTag(tagID int, userRole string, offset, limit int) ([]dto.BookPreviewDTO, error)
	SetBookAuthors(bookID int, authorIDs []int, userID int, userRole string) error
	AddBookAuthor(bookID, authorID int, userID int, userRole string) error
	RemoveBookAuthor(bookID, authorID int, userID int, userRole string) error
	SetBookTags(bookID int, tagIDs []int, userID int, userRole string) error
	AddBookTag(bookID, tagID int, userID int, userRole string) error
	RemoveBookTag(bookID, tagID int, userID int, userRole string) error
	UpdateBookStatus(bookID int, status string, userRole string) error
	SearchBooks(query string, userRole string, limit, offset int) ([]dto.BookPreviewDTO, error)
	GetDuplicateBooks(title string) ([]*models.Book, error)
	GetUserBooks(userID int) ([]dto.BookPreviewDTO, error)
	GetUserFavoriteBooks(userID int, userRole string) ([]dto.BookPreviewDTO, error)
	AddBookToFavorites(userID, bookID int) error
	RemoveBookFromFavorites(userID, bookID int) error
}

type bookService struct {
	repo        repository.BookRepository
	tagRepo     repository.TagRepository
	imageRepo   repository.ImageRepository
	fileRepo    repository.FileRepository
	commentRepo repository.CommentRepository
}

func NewBookService(repo repository.BookRepository,
	tagRepo repository.TagRepository,
	imageRepo repository.ImageRepository,
	fileRepo repository.FileRepository,
	commentRepo repository.CommentRepository,
) BookService {
	return &bookService{repo: repo, tagRepo: tagRepo, imageRepo: imageRepo, fileRepo: fileRepo, commentRepo: commentRepo}
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

func (s *bookService) GetBookDTO(bookID int, userID int, userRole string) (*dto.BookDTO, error) {
	statuses := getViewableStatuses(userRole)

	book, err := s.repo.GetBookByID(bookID, statuses)
	if err != nil {
		return nil, err
	}

	// Проверка: доступен ли контент этому пользователю
	if (book.Status == models.StatusBookPrivate || book.Status == models.StatusBookQuarantine) &&
		!(book.CreatedBy == userID || roles.IsAdmin(userRole)) {
		return nil, fmt.Errorf("access denied to this book")
	}

	authors, err := s.repo.GetAuthorsByBookID(bookID)
	if err != nil {
		return nil, err
	}

	tags, err := s.tagRepo.GetTagsByBookID(bookID)
	if err != nil {
		return nil, err
	}

	bookTags, err := s.tagRepo.GetBookTags(bookID)
	if err != nil {
		return nil, err
	}

	files, err := s.fileRepo.GetFilesByBookID(bookID)
	if err != nil {
		return nil, err
	}

	images, err := s.imageRepo.GetImagesByBookID(bookID)
	if err != nil {
		return nil, err
	}

	return &dto.BookDTO{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		PublishYear: book.PublishYear,
		Pages:       book.Pages,
		Language:    book.Language,
		Publisher:   book.Publisher,
		Type:        book.Type,
		CoverURL:    book.CoverURL,
		Authors:     dto.ConvertAuthors(authors),
		Tags:        dto.MergeTagsWithWeight(tags, bookTags),
		Files:       dto.ConvertBookFiles(files),
		Images:      dto.ConvertBookImages(images),
	}, nil
}

func (s *bookService) GetBookExtrasDTO(bookID int, userID int, userRole string) (*dto.BookExtrasDTO, error) {

	inFav, err := s.repo.IsBookInFavorites(bookID, userID)
	if err != nil {
		return nil, err
	}

	comments, err := s.commentRepo.GetCommentsByBookID(bookID)
	if err != nil {
		return nil, err
	}

	/* TODO
	 relations, err :=s.repo.GetRelationsByBookID(bookID)
	if err != nil {
		return nil, err
	}

	/*  TODO
	rating, err := s.repo.GetRatingByBookID(bookID)
	if err != nil {
		return nil, err
	}
	*/

	return &dto.BookExtrasDTO{
		InFavorites: inFav,
		Comments:    dto.ConvertBookComments(comments),
		// Relations:   dto.ConvertBookRelations(relations), // TODO
		//  Rating:      rating,  // TODO
	}, nil
}

func (s *bookService) GetBooksByStatuses(userRole string, offset, limit int) ([]models.Book, error) {
	statuses := getViewableStatuses(userRole)
	return s.repo.GetBooksByStatuses(statuses, offset, limit)
}

func (s *bookService) GetBooksByAuthor(authorID int, userRole string, offset, limit int) ([]dto.BookPreviewDTO, error) {
	statuses := getViewableStatuses(userRole)
	books, err := s.repo.GetBooksByAuthor(authorID, statuses, limit, offset)

	if err != nil {
		return nil, err
	}

	return dto.ConvertBooksToPreviewDTOs(
		books,
		s.repo.GetAuthorsByBookID,
		s.imageRepo.GetImagesByBookID,
	)
}

func (s *bookService) GetBooksByTag(tagID int, userRole string, offset, limit int) ([]dto.BookPreviewDTO, error) {
	statuses := getViewableStatuses(userRole)
	books, err := s.repo.GetBooksByTag(tagID, statuses, limit, offset)
	if err != nil {
		return nil, err
	}
	return dto.ConvertBooksToPreviewDTOs(
		books,
		s.repo.GetAuthorsByBookID,
		s.imageRepo.GetImagesByBookID,
	)
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

func (s *bookService) SearchBooks(query string, userRole string, limit, offset int) ([]dto.BookPreviewDTO, error) {
	statuses := getViewableStatuses(userRole)
	books, err := s.repo.SearchBooks(query, statuses, limit, offset)
	if err != nil {
		return nil, err
	}

	return dto.ConvertBooksToPreviewDTOs(
		books,
		s.repo.GetAuthorsByBookID,
		s.imageRepo.GetImagesByBookID,
	)
}

func (s *bookService) GetDuplicateBooks(title string) ([]*models.Book, error) {
	return s.repo.GetDuplicateBooks(title)
}

func (s *bookService) GetUserBooks(userID int) ([]dto.BookPreviewDTO, error) {
	books, err := s.repo.GetUserBooks(userID)
	if err != nil {
		return nil, err
	}

	return dto.ConvertBooksToPreviewDTOs(
		books,
		s.repo.GetAuthorsByBookID,
		s.imageRepo.GetImagesByBookID,
	)
}

func (s *bookService) GetUserFavoriteBooks(userID int, userRole string) ([]dto.BookPreviewDTO, error) {
	statuses := getViewableStatuses(userRole)
	books, err := s.repo.GetUserFavoriteBooks(userID, statuses)
	if err != nil {
		return nil, err
	}

	return dto.ConvertBooksToPreviewDTOs(
		books,
		s.repo.GetAuthorsByBookID,
		s.imageRepo.GetImagesByBookID,
	)
}

func (s *bookService) AddBookToFavorites(userID, bookID int) error {
	return s.repo.AddBookToFavorites(userID, bookID)
}

func (s *bookService) RemoveBookFromFavorites(userID, bookID int) error {
	return s.repo.RemoveBookFromFavorites(userID, bookID)
}
