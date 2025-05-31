package repository

import (
	"database/sql"
	"github.com/lib/pq"
	"online_library/backend/internal/models"
)

type CommentRepository interface {
	Create(comment *models.Comment) error
	Update(comment *models.Comment) error
	Delete(id int) error

	GetByID(id int) (*models.Comment, error)
	GetByBookID(bookID int, limit, offset int, statuses []string) ([]models.Comment, error)
	GetByUserID(userID int, limit, offset int) ([]models.Comment, error)
	GetLast(limit int) ([]models.Comment, error)
	SetStatus(id int, status string) error

	CountByBook(bookID int) (int, error)
}

type commentRepo struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepo{db: db}
}

func (r *commentRepo) Create(comment *models.Comment) error {
	query := `
		INSERT INTO comments (book_id, user_id, text, created_at, updated_at, status)
		VALUES ($1, $2, $3, NOW(), NOW(), $4)
		RETURNING id`
	return r.db.QueryRow(query, comment.BookID, comment.UserID, comment.Text, comment.Status).Scan(&comment.ID)
}

func (r *commentRepo) Update(comment *models.Comment) error {
	query := `
		UPDATE comments
		SET text = $1, updated_at = NOW(), status = $2
		WHERE id = $3`
	_, err := r.db.Exec(query, comment.Text, comment.Status, comment.ID)
	return err
}

func (r *commentRepo) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM comments WHERE id = $1`, id)
	return err
}

func (r *commentRepo) GetByID(id int) (*models.Comment, error) {
	var c models.Comment
	query := `SELECT id, book_id, user_id, text, created_at, updated_at, status FROM comments WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&c.ID, &c.BookID, &c.UserID, &c.Text, &c.CreatedAt, &c.UpdatedAt, &c.Status)
	return &c, err
}

func (r *commentRepo) GetByBookID(bookID int, limit, offset int, statuses []string) ([]models.Comment, error) {
	var comments []models.Comment

	query := `SELECT id, book_id, user_id, text, created_at, updated_at, status
	          FROM comments
	          WHERE book_id = $1 AND status = ANY($2)
	          ORDER BY created_at DESC
	          LIMIT $3 OFFSET $4`

	rows, err := r.db.Query(query, bookID, pq.Array(statuses), limit, offset)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var c models.Comment
		if err := rows.Scan(&c.ID, &c.BookID, &c.UserID, &c.Text, &c.CreatedAt, &c.UpdatedAt, &c.Status); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, nil
}

func (r *commentRepo) GetByUserID(userID int, limit, offset int) ([]models.Comment, error) {
	var comments []models.Comment

	query := `SELECT id, book_id, user_id, text, created_at, updated_at, status
	          FROM comments
	          WHERE user_id = $1
	          ORDER BY created_at DESC
	          LIMIT $2 OFFSET $3`

	rows, err := r.db.Query(query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var c models.Comment
		if err := rows.Scan(&c.ID, &c.BookID, &c.UserID, &c.Text, &c.CreatedAt, &c.UpdatedAt, &c.Status); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, nil
}

func (r *commentRepo) GetLast(limit int) ([]models.Comment, error) {
	query := `SELECT id, book_id, user_id, text, created_at, updated_at, status
	          FROM comments
	          WHERE status = $1
	          ORDER BY created_at DESC
	          LIMIT $2`

	rows, err := r.db.Query(query, models.CommentStatusActive, limit)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var comments []models.Comment
	for rows.Next() {
		var c models.Comment
		if err := rows.Scan(&c.ID, &c.BookID, &c.UserID, &c.Text, &c.CreatedAt, &c.UpdatedAt, &c.Status); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, nil
}

func (r *commentRepo) SetStatus(id int, status string) error {
	query := `UPDATE comments SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.Exec(query, status, id)
	return err
}

func (r *commentRepo) CountByBook(bookID int) (int, error) {
	var count int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM comments WHERE book_id = $1 AND status = $2`, bookID, models.CommentStatusActive).Scan(&count)
	return count, err
}
