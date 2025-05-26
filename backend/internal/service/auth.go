package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"online_library/backend/internal/pkg/auth"
	"online_library/backend/internal/repository"
)

type AuthService struct {
	Repo        *repository.UserRepo
	UserService UserService
}

func NewAuthService(repo *repository.UserRepo, userService UserService) *AuthService {
	return &AuthService{
		Repo:        repo,
		UserService: userService,
	}
}

func (s *AuthService) Register(email, name, password string, bio string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = s.Repo.SetNewUser(email, name, string(hash), bio)
	if err != nil {
		return err
	}
	return err
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.Repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if user.Is_active == false {
		return "", errors.New("invalid credentials")
	}

	token, err := auth.GenerateToken(user.ID, user.Role, user.TokenVersion)
	if err != nil {
		return "", err
	}
	return token, nil
}
