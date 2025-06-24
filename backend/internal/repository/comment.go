package repository

import (
	"database/sql"
	"github.com/lib/pq"
	"online_library/backend/internal/models"
)

// CommentRepository определяет интерфейс для работы с комментариями к книгам.
type CommentRepository interface {
	// Create добавляет новый комментарий к книге.
	Create(comment *models.BookComment) error

	// Update обновляет текст и статус существующего комментария.
	Update(comment *models.BookComment) error

	// Delete удаляет комментарий по его ID.
	Delete(id int) error

	// GetByID возвращает комментарий по его ID.
	GetByID(id int) (*models.BookComment, error)

	// GetByBookID возвращает список комментариев к книге с поддержкой пагинации и фильтрации по статусу.
	GetByBookID(bookID int, limit, offset int, statuses []string) ([]models.BookComment, error)

	// GetCommentsByBookID возвращает активные комментарии для заданной книги (без пагинации).
	GetCommentsByBookID(bookID int) ([]models.BookComment, error)

	// GetByUserID возвращает список комментариев, оставленных пользователем, с пагинацией.
	GetByUserID(userID int, limit, offset int) ([]models.BookComment, error)

	// GetLast возвращает последние активные комментарии.
	GetLast(limit int) ([]models.BookComment, error)

	// SetStatus обновляет статус комментария (например, active, deleted, pending).
	SetStatus(id int, status string) error

	// CountByBook возвращает количество активных комментариев у книги.
	CountByBook(bookID int) (int, error)
}

type commentRepo struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepo{db: db}
}

func (r *commentRepo) Create(comment *models.BookComment) error {
	query := `
		INSERT INTO comments (book_id, user_id, text, created_at, updated_at, status)
		VALUES ($1, $2, $3, NOW(), NOW(), $4)
		RETURNING id`
	return r.db.QueryRow(query, comment.BookID, comment.UserID, comment.Text, comment.Status).Scan(&comment.ID)
}

func (r *commentRepo) Update(comment *models.BookComment) error {
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

func (r *commentRepo) GetByID(id int) (*models.BookComment, error) {
	var c models.BookComment
	query := `SELECT id, book_id, user_id, text, created_at, updated_at, status FROM comments WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&c.ID, &c.BookID, &c.UserID, &c.Text, &c.CreatedAt, &c.UpdatedAt, &c.Status)
	return &c, err
}

func (r *commentRepo) GetByBookID(bookID int, limit, offset int, statuses []string) ([]models.BookComment, error) {
	var comments []models.BookComment

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
		var c models.BookComment
		if err := rows.Scan(&c.ID, &c.BookID, &c.UserID, &c.Text, &c.CreatedAt, &c.UpdatedAt, &c.Status); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, nil
}

func (r *commentRepo) GetCommentsByBookID(bookID int) ([]models.BookComment, error) {
	query := `
		SELECT id, book_id, user_id, text, created_at, updated_at, status
		FROM comments
		WHERE book_id = $1 AND status = ANY($2)
		ORDER BY created_at DESC
	`
	statuses := []string{models.CommentStatusActive} // по умолчанию, можно вынести выше
	rows, err := r.db.Query(query, bookID, pq.Array(statuses))
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var comments []models.BookComment
	for rows.Next() {
		var c models.BookComment
		if err := rows.Scan(&c.ID, &c.BookID, &c.UserID, &c.Text, &c.CreatedAt, &c.UpdatedAt, &c.Status); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, nil
}

func (r *commentRepo) GetByUserID(userID int, limit, offset int) ([]models.BookComment, error) {
	var comments []models.BookComment

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
		var c models.BookComment
		if err := rows.Scan(&c.ID, &c.BookID, &c.UserID, &c.Text, &c.CreatedAt, &c.UpdatedAt, &c.Status); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, nil
}

func (r *commentRepo) GetLast(limit int) ([]models.BookComment, error) {
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

	var comments []models.BookComment
	for rows.Next() {
		var c models.BookComment
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
