package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"online_library/backend/internal/models"
)

type UserRepository interface {
	GetAllActive(ctx context.Context) ([]models.User, error)
	GetByEmail(email string) (*models.User, error)
	SetNewUser(email string, name string, passwordHash string, bio string) (*models.User, error)
	CheckEmailExists(email string) (bool, error)
	UpdateUserByID(ctx context.Context, id int, input models.UserInput) (*models.User, error)
	SoftDeleteUserByID(ctx context.Context, id int) error
	HardDeleteUserByID(ctx context.Context, id int) error
	AdminUpdateUser(ctx context.Context, id int, input models.AdminUserUpdateInput) (*models.User, error)
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetAllActive(ctx context.Context) ([]models.User, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, email, name, role, bio, registered_at FROM users WHERE is_active = TRUE")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("failed to close rows: %v", err)
		}
	}(rows)

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Email, &u.Name, &u.Role, &u.Bio, &u.RegisteredAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepo) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(`SELECT id, email, password_hash, role, token_version, is_active FROM users WHERE email = $1`, email).
		Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Role, &user.TokenVersion, &user.Is_active)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetByID(ctx context.Context, id int) (*models.User, error) {
	var u models.User
	err := r.db.QueryRowContext(ctx, `
		SELECT id, email, name, role, bio, registered_at
		FROM users WHERE id = $1
	`, id).Scan(&u.ID, &u.Email, &u.Name, &u.Role, &u.Bio, &u.RegisteredAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepo) SetNewUser(email string, name string, passwordHash string, bio string) (*models.User, error) {
	exists, err := r.CheckEmailExists(email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("email already exists")
	}

	_, err = r.db.Exec(`
		INSERT INTO users (email, name, password_hash, role, bio)
		VALUES ($1, $2, $3, $4, $5)
		`, email, name, passwordHash, models.RoleNewUser, bio)

	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	return nil, nil
}

func (r *UserRepo) CheckEmailExists(email string) (bool, error) {
	var count int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM users WHERE email = $1`, email).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check email existence: %w", err)
	}
	return count > 0, nil
}

func (r *UserRepo) UpdateUserByID(ctx context.Context, id int, input models.UserInput) (*models.User, error) {
	_, err := r.db.ExecContext(ctx, `
		UPDATE users
		SET email = $1, name = $2, bio = $3
		WHERE id = $4 AND is_active = TRUE
	`, input.Email, input.Name, input.Bio, id)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	user, err := r.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated user: %w", err)
	}

	return user, nil
}

func (r *UserRepo) SoftDeleteUserByID(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE users
		SET is_active = FALSE
		WHERE id = $1
	`, id)
	return err
}

func (r *UserRepo) HardDeleteUserByID(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, `
		DELETE FROM users
		WHERE id = $1
	`, id)
	return err
}

func (r *UserRepo) AdminUpdateUser(ctx context.Context, id int, input models.AdminUserUpdateInput) (*models.User, error) {
	_, err := r.db.ExecContext(ctx, `
		UPDATE users
		SET email = $1, name = $2, bio = $3, role = $4, token_version = token_version + 1
		WHERE id = $5
	`, input.Email, input.Name, input.Bio, input.Role, id)
	if err != nil {
		return nil, err
	}
	return r.GetByID(ctx, id)
}

// При смене роли у пользователя UPDATE users SET role = 'user', token_version = token_version + 1 WHERE id = ?;
