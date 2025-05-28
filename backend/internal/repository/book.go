package repository

import (
	"database/sql"
	"fmt"
	"online_library/backend/internal/models"
	"strings"
)

type BookRepository interface {
	GetBooksByTags(tagIDs []int, limit, offset int, sort string) ([]models.Book, error)
}

type bookRepo struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepo{db: db}
}

func (r *BookRepository) GetBooksByTags(tagIDs []int, limit, offset int, sort string) ([]models.Book, error) {
	// --- сортировка ---
	var orderBy string
	switch sort {
	case "title_desc":
		orderBy = "ORDER BY b.title DESC"
	default: // "title_asc"
		orderBy = "ORDER BY b.title ASC"
	}

	// --- подготовка SQL ---
	placeholders := make([]string, len(tagIDs))
	args := make([]interface{}, 0, len(tagIDs)+3)

	for i, id := range tagIDs {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args = append(args, id)
	}
	tagCountIdx := len(args) + 1
	limitIdx := tagCountIdx + 1
	offsetIdx := limitIdx + 1

	query := fmt.Sprintf(`
		SELECT b.*
		FROM books b
		JOIN book_tags bt ON b.id = bt.book_id
		WHERE bt.tag_id IN (%s)
		GROUP BY b.id
		HAVING COUNT(DISTINCT bt.tag_id) = $%d
		%s
		LIMIT $%d OFFSET $%d
	`, strings.Join(placeholders, ","), tagCountIdx, orderBy, limitIdx, offsetIdx)

	args = append(args, len(tagIDs), limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		// заполнение полей книги
		if err := rows.Scan(&b.ID, &b.Title /* и т.д. */); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}
