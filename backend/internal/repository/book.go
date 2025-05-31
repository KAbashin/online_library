package repository

import (
	"database/sql"
	"fmt"
	"online_library/backend/internal/models"
	"strings"
)

type BookRepository interface {
	CreateBook(book *models.Book) (int, error)
	UpdateBook(book *models.Book) error
	DeleteBook(id int) error
	GetBookByID(id int, allowedStatuses []string) (*models.Book, error)
	GetBooksByStatuses(statuses []string, offset, limit int) ([]models.Book, error)
	GetBooksByAuthor(authorID int, statuses []string, limit, offset int) ([]models.Book, error)
	GetBooksByTag(tagID int, statuses []string, limit, offset int) ([]models.Book, error)
	SetBookAuthors(bookID int, authorIDs []int) error
	AddBookAuthor(bookID, authorID int) error
	RemoveBookAuthor(bookID, authorID int) error
	SetBookTags(bookID int, tagIDs []int) error
	AddBookTag(bookID, tagID int) error
	RemoveBookTag(bookID, tagID int) error
	SearchBooks(query string, allowedStatuses []string, limit, offset int) ([]*models.Book, error)
	GetDuplicateBooks(title string) ([]*models.Book, error)
	GetUserBooks(userID int) ([]*models.Book, error)
	GetUserFavoriteBooks(userID int, statuses []string) ([]*models.Book, error)
	AddBookToFavorites(userID, bookID int) error
	RemoveBookFromFavorites(userID, bookID int) error
	UpdateBookStatus(bookID int, status string) error
	GetBookMeta(bookID int) (*models.Book, error) // Только базовые данные: id, created_by
}

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) CreateBook(book *models.Book) (int, error) {
	query := `
		INSERT INTO books (title, description, publish_year, pages, language, publisher, type, rating, cover_url, status, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,NOW())
		RETURNING id
	`
	var id int
	err := r.db.QueryRow(query,
		book.Title, book.Description, book.PublishYear, book.Pages,
		book.Language, book.Publisher, book.Type, book.Rating,
		book.CoverURL, book.Status,
	).Scan(&id)
	return id, err
}

func (r *bookRepository) UpdateBook(book *models.Book) error {
	query := `
		UPDATE books
		SET title=$1, description=$2, publish_year=$3, pages=$4, language=$5,
		    publisher=$6, type=$7, rating=$8, cover_url=$9, status=$10
		WHERE id=$11
	`
	_, err := r.db.Exec(query,
		book.Title, book.Description, book.PublishYear, book.Pages,
		book.Language, book.Publisher, book.Type, book.Rating,
		book.CoverURL, book.Status, book.ID,
	)
	return err
}

func (r *bookRepository) DeleteBook(id int) error {
	_, err := r.db.Exec("DELETE FROM books WHERE id=$1", id)
	return err
}

func (r *bookRepository) GetBookByID(id int, allowedStatuses []string) (*models.Book, error) {
	if len(allowedStatuses) == 0 {
		return nil, fmt.Errorf("no statuses allowed")
	}

	placeholders := make([]string, len(allowedStatuses))
	args := make([]interface{}, len(allowedStatuses)+1)
	args[0] = id

	for i, status := range allowedStatuses {
		placeholders[i] = fmt.Sprintf("$%d", i+2)
		args[i+1] = status
	}

	query := "SELECT id, title, description, publish_year, pages, language, " +
		"publisher, type, rating, cover_url, status, created_at " +
		"FROM books " +
		"WHERE id = $1 AND status IN (" + strings.Join(placeholders, ", ") + ")"

	row := r.db.QueryRow(query, args...)
	var b models.Book
	err := row.Scan(
		&b.ID, &b.Title, &b.Description, &b.PublishYear, &b.Pages,
		&b.Language, &b.Publisher, &b.Type, &b.Rating,
		&b.CoverURL, &b.Status, &b.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *bookRepository) GetBooksByStatuses(statuses []string, offset, limit int) ([]models.Book, error) {
	if len(statuses) == 0 {
		return nil, fmt.Errorf("no statuses provided")
	}

	placeholders := make([]string, len(statuses))
	args := make([]interface{}, len(statuses)+2)

	for i, status := range statuses {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = status
	}
	args[len(statuses)] = limit
	args[len(statuses)+1] = offset

	query := "SELECT id, title, description, publish_year, pages, language, " +
		"publisher, type, rating, cover_url, status, created_at " +
		"FROM books " +
		"WHERE status IN (" + strings.Join(placeholders, ", ") + ") " +
		"ORDER BY created_at DESC " +
		"LIMIT $" + fmt.Sprint(len(statuses)+1) + " OFFSET $" + fmt.Sprint(len(statuses)+2)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var books []models.Book
	for rows.Next() {
		var b models.Book
		err := rows.Scan(
			&b.ID, &b.Title, &b.Description, &b.PublishYear, &b.Pages,
			&b.Language, &b.Publisher, &b.Type, &b.Rating,
			&b.CoverURL, &b.Status, &b.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (r *bookRepository) GetBooksByAuthor(authorID int, statuses []string, limit, offset int) ([]models.Book, error) {
	if len(statuses) == 0 {
		return nil, fmt.Errorf("no statuses provided")
	}

	placeholders := make([]string, len(statuses))
	args := make([]interface{}, len(statuses)+3)
	args[0] = authorID

	for i, status := range statuses {
		placeholders[i] = fmt.Sprintf("$%d", i+2)
		args[i+1] = status
	}
	args[len(statuses)+1] = limit
	args[len(statuses)+2] = offset

	query := "SELECT b.id, b.title, b.description, b.publish_year, b.pages, b.language, " +
		"b.publisher, b.type, b.rating, b.cover_url, b.status, b.created_at " +
		"FROM books b " +
		"JOIN book_authors ba ON b.id = ba.book_id " +
		"WHERE ba.author_id = $1 AND b.status IN (" + strings.Join(placeholders, ", ") + ") " +
		"ORDER BY b.created_at DESC " +
		"LIMIT $" + fmt.Sprint(len(statuses)+2) + " OFFSET $" + fmt.Sprint(len(statuses)+3)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var books []models.Book
	for rows.Next() {
		var b models.Book
		err := rows.Scan(
			&b.ID, &b.Title, &b.Description, &b.PublishYear, &b.Pages,
			&b.Language, &b.Publisher, &b.Type, &b.Rating,
			&b.CoverURL, &b.Status, &b.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (r *bookRepository) GetBooksByTag(tagID int, statuses []string, limit, offset int) ([]models.Book, error) {
	if len(statuses) == 0 {
		return nil, fmt.Errorf("no statuses provided")
	}

	placeholders := make([]string, len(statuses))
	args := make([]interface{}, len(statuses)+3)
	args[0] = tagID

	for i, status := range statuses {
		placeholders[i] = fmt.Sprintf("$%d", i+2)
		args[i+1] = status
	}
	args[len(statuses)+1] = limit
	args[len(statuses)+2] = offset

	query := "SELECT b.id, b.title, b.description, b.publish_year, b.pages, b.language, " +
		"b.publisher, b.type, b.rating, b.cover_url, b.status, b.created_at " +
		"FROM books b " +
		"JOIN book_tags bt ON b.id = bt.book_id " +
		"WHERE bt.tag_id = $1 AND b.status IN (" + strings.Join(placeholders, ", ") + ") " +
		"ORDER BY b.created_at DESC" +
		"LIMIT $" + fmt.Sprint(len(statuses)+2) + " OFFSET $" + fmt.Sprint(len(statuses)+3)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var books []models.Book
	for rows.Next() {
		var b models.Book
		err := rows.Scan(
			&b.ID, &b.Title, &b.Description, &b.PublishYear, &b.Pages,
			&b.Language, &b.Publisher, &b.Type, &b.Rating,
			&b.CoverURL, &b.Status, &b.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (r *bookRepository) SetBookAuthors(bookID int, authorIDs []int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {

		}
	}(tx)

	_, err = tx.Exec("DELETE FROM book_authors WHERE book_id = $1", bookID)
	if err != nil {
		return err
	}

	for _, authorID := range authorIDs {
		_, err := tx.Exec("INSERT INTO book_authors (book_id, author_id) VALUES ($1, $2)", bookID, authorID)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (r *bookRepository) AddBookAuthor(bookID, authorID int) error {
	_, err := r.db.Exec("INSERT INTO book_authors (book_id, author_id) VALUES ($1, $2) ON CONFLICT DO NOTHING", bookID, authorID)
	return err
}

func (r *bookRepository) RemoveBookAuthor(bookID, authorID int) error {
	_, err := r.db.Exec("DELETE FROM book_authors WHERE book_id = $1 AND author_id = $2", bookID, authorID)
	return err
}

func (r *bookRepository) SetBookTags(bookID int, tagIDs []int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {

		}
	}(tx)

	_, err = tx.Exec("DELETE FROM book_tags WHERE book_id = $1", bookID)
	if err != nil {
		return err
	}

	for _, tagID := range tagIDs {
		_, err := tx.Exec("INSERT INTO book_tags (book_id, tag_id, weight) VALUES ($1, $2, 1)", bookID, tagID)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (r *bookRepository) AddBookTag(bookID, tagID int) error {
	_, err := r.db.Exec("INSERT INTO book_tags (book_id, tag_id, weight) VALUES ($1, $2, 1) ON CONFLICT DO NOTHING", bookID, tagID)
	return err
}

func (r *bookRepository) RemoveBookTag(bookID, tagID int) error {
	_, err := r.db.Exec("DELETE FROM book_tags WHERE book_id = $1 AND tag_id = $2", bookID, tagID)
	return err
}

func (r *bookRepository) UpdateBookStatus(bookID int, status string) error {
	_, err := r.db.Exec("UPDATE books SET status = $1 WHERE id = $2", status, bookID)
	return err
}

func (r *bookRepository) SearchBooks(query string, allowedStatuses []string, limit, offset int) ([]*models.Book, error) {
	if len(allowedStatuses) == 0 {
		return nil, fmt.Errorf("no allowed statuses")
	}

	placeholders := make([]string, len(allowedStatuses))
	args := make([]interface{}, 0)
	args = append(args, "%"+query+"%")
	for i, status := range allowedStatuses {
		placeholders[i] = fmt.Sprintf("$%d", i+2)
		args = append(args, status)
	}
	args = append(args, limit, offset)

	q := "SELECT id, title, description, publish_year, pages, language, publisher, type, rating, cover_url, status, created_at " +
		"FROM books " +
		"WHERE (title ILIKE $1 OR description ILIKE $1) " +
		"AND status IN (" + strings.Join(placeholders, ", ") + ") " +
		"ORDER BY created_at DESC " +
		"LIMIT $" + fmt.Sprint(len(args)-1) + " OFFSET $" + fmt.Sprint(len(args))

	rows, err := r.db.Query(q, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var books []*models.Book
	for rows.Next() {
		var b models.Book
		err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.PublishYear, &b.Pages,
			&b.Language, &b.Publisher, &b.Type, &b.Rating,
			&b.CoverURL, &b.Status, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &b)
	}
	return books, nil
}

func (r *bookRepository) GetDuplicateBooks(title string) ([]*models.Book, error) {
	rows, err := r.db.Query(`
		SELECT id, title, description, publish_year, pages, language,
		       publisher, type, rating, cover_url, status, created_at
		FROM books
		WHERE LOWER(title) LIKE LOWER($1)`, "%"+title+"%")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var books []*models.Book
	for rows.Next() {
		var b models.Book
		err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.PublishYear, &b.Pages,
			&b.Language, &b.Publisher, &b.Type, &b.Rating,
			&b.CoverURL, &b.Status, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &b)
	}
	return books, nil
}

func (r *bookRepository) GetUserBooks(userID int) ([]*models.Book, error) {
	rows, err := r.db.Query(`
		SELECT id, title, description, publish_year, pages, language,
		       publisher, type, rating, cover_url, status, created_at
		FROM books
		WHERE created_by = $1
		ORDER BY created_at DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var books []*models.Book
	for rows.Next() {
		var b models.Book
		err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.PublishYear, &b.Pages,
			&b.Language, &b.Publisher, &b.Type, &b.Rating,
			&b.CoverURL, &b.Status, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &b)
	}
	return books, nil
}

func (r *bookRepository) GetUserFavoriteBooks(userID int, statuses []string) ([]*models.Book, error) {
	if len(statuses) == 0 {
		return []*models.Book{}, nil // Возвращаем пустой список, если доступных статусов нет
	}

	placeholders := make([]string, len(statuses))
	args := make([]interface{}, len(statuses)+1)
	args[0] = userID

	for i, status := range statuses {
		placeholders[i] = fmt.Sprintf("$%d", i+2)
		args[i+1] = status
	}

	query :=
		"SELECT b.id, b.title, b.description, b.publish_year, b.pages, b.language, " +
			"b.publisher, b.type, b.rating, b.cover_url, b.status, b.created_at " +
			"FROM books b " +
			"JOIN book_favorites f ON b.id = f.book_id " +
			"WHERE f.user_id = $1 AND b.status IN (" + strings.Join(placeholders, ", ") + ") " +
			"ORDER BY f.created_at DESC"

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var books []*models.Book
	for rows.Next() {
		var b models.Book
		err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.PublishYear, &b.Pages,
			&b.Language, &b.Publisher, &b.Type, &b.Rating,
			&b.CoverURL, &b.Status, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &b)
	}
	return books, nil
}

func (r *bookRepository) AddBookToFavorites(userID, bookID int) error {
	_, err := r.db.Exec(`
		INSERT INTO book_favorites (user_id, book_id, created_at)
		VALUES ($1, $2, NOW())
		ON CONFLICT DO NOTHING
	`, userID, bookID)
	return err
}

func (r *bookRepository) RemoveBookFromFavorites(userID, bookID int) error {
	_, err := r.db.Exec(`
		DELETE FROM book_favorites
		WHERE user_id = $1 AND book_id = $2
	`, userID, bookID)
	return err
}

func (r *bookRepository) GetBookMeta(bookID int) (*models.Book, error) {
	query := `SELECT id, created_by FROM books WHERE id = $1`
	row := r.db.QueryRow(query, bookID)
	var book models.Book
	err := row.Scan(&book.ID, &book.CreatedBy)
	if err != nil {
		return nil, err
	}
	return &book, nil
}
