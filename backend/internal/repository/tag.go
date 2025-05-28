package repository

import (
	"database/sql"
	"online_library/backend/internal/models"
)

type TagRepository interface {
	GetAllTags() ([]models.Tag, error)
	GetTagByID(id int) (models.Tag, error)

	CreateTag(tag *models.Tag) error
	UpdateTag(tag *models.Tag) error
	DeleteTag(id int) error

	GetTagsByBookID(bookID int) ([]models.Tag, error)
	AssignTagToBook(bookTag *models.BookTag) error
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
