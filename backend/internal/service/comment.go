package service

import (
	"fmt"
	"online_library/backend/internal/models"
	"online_library/backend/internal/pkg/roles"
	"online_library/backend/internal/repository"
	"time"
)

type CommentService interface {
	Create(comment *models.Comment) error
	Update(comment *models.Comment, userID int, userRole string) error
	Delete(id, userID int, userRole string) error

	GetByID(id int) (*models.Comment, error)
	GetByBookID(bookID int, limit, offset int, statuses []string) ([]models.Comment, error)
	GetByUserID(userID int, limit, offset int) ([]models.Comment, error)
	GetLast(limit int) ([]models.Comment, error)

	SetStatus(id int, status, userRole string) error
	CountByBook(bookID int) (int, error)
}

type commentService struct {
	repo repository.CommentRepository
	//logger *zap.SugaredLogger
}

func NewCommentService(repo repository.CommentRepository) CommentService { // logger *zap.SugaredLogger
	return &commentService{repo: repo} // , logger: logger
}

func (s *commentService) Create(comment *models.Comment) error {
	comment.Status = models.CommentStatusActive // либо "pending", если будет модерация
	return s.repo.Create(comment)
}

func (s *commentService) Update(comment *models.Comment, userID int, userRole string) error {
	existing, err := s.repo.GetByID(comment.ID)
	if err != nil {
		return err
	}

	isOwner := existing.UserID == userID
	isAdmin := userRole == roles.RoleAdmin || userRole == roles.RoleSuperAdmin

	if !isOwner && !isAdmin {
		return fmt.Errorf("access denied: not owner or admin")
	}

	existing.Text = comment.Text
	existing.UpdatedAt = time.Now()

	return s.repo.Update(existing)
}

func (s *commentService) Delete(id, userID int, userRole string) error {
	comment, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	isOwner := comment.UserID == userID
	isAdmin := userRole == roles.RoleAdmin || userRole == roles.RoleSuperAdmin

	if !isOwner && !isAdmin {
		return fmt.Errorf("access denied: cannot delete")
	}

	return s.repo.SetStatus(id, models.CommentStatusDeleted)
}

func (s *commentService) GetByID(id int) (*models.Comment, error) {
	return s.repo.GetByID(id)
}

func (s *commentService) GetByBookID(bookID int, limit, offset int, statuses []string) ([]models.Comment, error) {
	if len(statuses) == 0 {
		statuses = []string{models.CommentStatusActive} // по умолчанию только активные
	}

	return s.repo.GetByBookID(bookID, limit, offset, statuses)
}

func (s *commentService) GetByUserID(userID int, limit, offset int) ([]models.Comment, error) {
	return s.repo.GetByUserID(userID, limit, offset)
}

func (s *commentService) GetLast(limit int) ([]models.Comment, error) {
	return s.repo.GetLast(limit)
}

func (s *commentService) SetStatus(id int, status, userRole string) error {
	if userRole != roles.RoleAdmin && userRole != roles.RoleSuperAdmin {
		return fmt.Errorf("access denied: only admin can change status")
	}

	return s.repo.SetStatus(id, status)
}

func (s *commentService) CountByBook(bookID int) (int, error) {
	return s.repo.CountByBook(bookID)
}
