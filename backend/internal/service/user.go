package service

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"online_library/backend/internal/models"
	"online_library/backend/internal/repository"
)

type UserService interface {
	GetAllUsers(ctx context.Context) ([]models.User, error)
	CreateUser(ctx context.Context, input models.UserInput) (*models.User, error)
	UpdateUser(ctx context.Context, id int, input models.UserInput) (*models.User, error)
	UpdateUserByAdmin(ctx context.Context, id int, input models.AdminUserUpdateInput) (*models.User, error)
	CheckEmailExists(email string) (bool, error)
	SoftDeleteUser(ctx context.Context, id int) error
	HardDeleteUser(ctx context.Context, id int) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAllActive(ctx)
}

func (s *userService) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	exists, err := s.repo.CheckEmailExists(input.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("email already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return s.repo.SetNewUser(input.Email, input.Name, string(hashed), input.Bio)
}

func (s *userService) UpdateUser(ctx context.Context, id int, input models.UserInput) (*models.User, error) {
	return s.repo.UpdateUserByID(ctx, id, input)
}

func (s *userService) UpdateUserByAdmin(ctx context.Context, id int, input models.AdminUserUpdateInput) (*models.User, error) {
	// возможно, тут стоит проверить права вызывающего (не в этом методе, а выше)
	return s.repo.AdminUpdateUser(ctx, id, input)
}

func (s *userService) CheckEmailExists(email string) (bool, error) {
	return s.repo.CheckEmailExists(email)
}

func (s *userService) SoftDeleteUser(ctx context.Context, id int) error {
	return s.repo.SoftDeleteUserByID(ctx, id)
}

func (s *userService) HardDeleteUser(ctx context.Context, id int) error {
	return s.repo.HardDeleteUserByID(ctx, id)
}
