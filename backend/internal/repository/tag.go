package repository

import (
	"database/sql"
	"online_library/backend/internal/models"
)

// TagRepository определяет методы для управления тегами книг и связью между книгами и тегами.
type TagRepository interface {
	// GetAllTags возвращает все доступные теги, отсортированные по имени.
	GetAllTags() ([]models.Tag, error)

	// GetTagByID возвращает тег по его уникальному ID.
	GetTagByID(id int) (models.Tag, error)

	// CreateTag сохраняет новый тег в базе и возвращает его ID.
	CreateTag(tag *models.Tag) error

	// UpdateTag обновляет имя и цвет существующего тега.
	UpdateTag(tag *models.Tag) error

	// DeleteTag удаляет тег по его ID.
	DeleteTag(id int) error

	// GetTagsByBookID возвращает список тегов, связанных с конкретной книгой.
	GetTagsByBookID(bookID int) ([]models.Tag, error)

	// GetBookTags возвращает ассоциации тегов и их веса для указанной книги.
	GetBookTags(bookID int) ([]models.BookTag, error)

	// AssignTagToBook связывает тег с книгой. Если связь уже существует — обновляет её вес.
	AssignTagToBook(bookTag *models.BookTag) error

	// RemoveTagFromBook удаляет связь между книгой и тегом.
	RemoveTagFromBook(bookID, tagID int) error
}

type tagRepo struct {
	db *sql.DB
}

func NewTagRepository(db *sql.DB) TagRepository {
	return &tagRepo{db: db}
}

func (r *tagRepo) GetAllTags() ([]models.Tag, error) {
	rows, err := r.db.Query(`SELECT id, name, color FROM tags ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var tags []models.Tag
	for rows.Next() {
		var tag models.Tag
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.Color); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func (r *tagRepo) GetTagByID(id int) (models.Tag, error) {
	var tag models.Tag
	err := r.db.QueryRow(`SELECT id, name, color FROM tags WHERE id = $1`, id).
		Scan(&tag.ID, &tag.Name, &tag.Color)
	return tag, err
}

func (r *tagRepo) CreateTag(tag *models.Tag) error {
	return r.db.QueryRow(
		`INSERT INTO tags (name, color) VALUES ($1, $2) RETURNING id`,
		tag.Name, tag.Color,
	).Scan(&tag.ID)
}

func (r *tagRepo) UpdateTag(tag *models.Tag) error {
	_, err := r.db.Exec(
		`UPDATE tags SET name = $1, color = $2 WHERE id = $3`,
		tag.Name, tag.Color, tag.ID,
	)
	return err
}

func (r *tagRepo) DeleteTag(id int) error {
	_, err := r.db.Exec(`DELETE FROM tags WHERE id = $1`, id)
	return err
}

func (r *tagRepo) GetTagsByBookID(bookID int) ([]models.Tag, error) {
	rows, err := r.db.Query(`
		SELECT t.id, t.name, t.color
		FROM tags t
		JOIN book_tags bt ON bt.tag_id = t.id
		WHERE bt.book_id = $1
	`, bookID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var tags []models.Tag
	for rows.Next() {
		var tag models.Tag
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.Color); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func (r *tagRepo) GetBookTags(bookID int) ([]models.BookTag, error) {
	query := `SELECT book_id, tag_id, weight FROM book_tags WHERE book_id = $1`
	rows, err := r.db.Query(query, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookTags []models.BookTag
	for rows.Next() {
		var bt models.BookTag
		if err := rows.Scan(&bt.BookID, &bt.TagID, &bt.Weight); err != nil {
			return nil, err
		}
		bookTags = append(bookTags, bt)
	}
	return bookTags, rows.Err()
}

func (r *tagRepo) AssignTagToBook(bt *models.BookTag) error {
	_, err := r.db.Exec(
		`INSERT INTO book_tags (book_id, tag_id, weight) VALUES ($1, $2, $3)
		 ON CONFLICT (book_id, tag_id) DO UPDATE SET weight = EXCLUDED.weight`,
		bt.BookID, bt.TagID, bt.Weight,
	)
	return err
}

func (r *tagRepo) RemoveTagFromBook(bookID, tagID int) error {
	_, err := r.db.Exec(`DELETE FROM book_tags WHERE book_id = $1 AND tag_id = $2`, bookID, tagID)
	return err
}
