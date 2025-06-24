package repository

import (
	"database/sql"
	"online_library/backend/internal/models"
	"online_library/backend/internal/pkg/translit"
)

// AuthorRepository определяет интерфейс для операций с авторами книг.
type AuthorRepository interface {
	// CreateAuthor добавляет нового автора в базу данных.
	// Если name_en не указан, он будет сгенерирован из name_ru с помощью транслитерации.
	CreateAuthor(author *models.Author) error

	// UpdateAuthor обновляет информацию об авторе в базе данных по его ID.
	UpdateAuthor(author *models.Author) error

	// DeleteAuthor удаляет автора по ID.
	DeleteAuthor(id int) error

	// GetAuthorByID возвращает информацию об авторе по его ID.
	GetAuthorByID(id int) (*models.Author, error)

	// SearchAuthorByName ищет авторов по имени (русскому или английскому).
	// Использует ILIKE для частичного совпадения. Поддерживает пагинацию.
	SearchAuthorByName(query string, limit, offset int) ([]*models.Author, error)

	// GetAllAuthors возвращает список всех авторов с пагинацией.
	GetAllAuthors(offset, limit int) ([]models.Author, error)

	// CountAuthors возвращает количество авторов, имя которых соответствует строке поиска.
	CountAuthors(query string) (int, error)

	// AuthorExists проверяет, существует ли автор с заданным именем (на русском или английском).
	// Исключает из поиска автора с указанным ID (используется при обновлении).
	AuthorExists(nameRu, nameEn string, excludeID int) (bool, error)
}

type authorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) AuthorRepository {
	return &authorRepository{db: db}
}

func (r *authorRepository) GetAuthorByID(id int) (*models.Author, error) {
	var a models.Author
	err := r.db.QueryRow(`SELECT id, name_ru, name_en, bio, photo_url FROM authors WHERE id = $1`, id).
		Scan(&a.ID, &a.NameRU, &a.NameEN, &a.Bio, &a.PhotoURL)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *authorRepository) GetAllAuthors(offset, limit int) ([]models.Author, error) {
	rows, err := r.db.Query(`
		SELECT id, name_ru, name_en, bio, photo_url
		FROM authors
		ORDER BY name_ru ASC
		LIMIT $1 OFFSET $2
	`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var authors []models.Author
	for rows.Next() {
		var a models.Author
		err := rows.Scan(&a.ID, &a.NameRU, &a.NameEN, &a.Bio, &a.PhotoURL)
		if err != nil {
			return nil, err
		}
		authors = append(authors, a)
	}
	return authors, nil
}

func (r *authorRepository) AuthorExists(nameRu, nameEn string, excludeID int) (bool, error) {
	var id int
	err := r.db.QueryRow(`
		SELECT id FROM authors
		WHERE (name_ru = $1 OR name_en = $2)
		  AND id != $3
		LIMIT 1
	`, nameRu, nameEn, excludeID).Scan(&id)

	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *authorRepository) CreateAuthor(author *models.Author) error {
	if author.NameEN == "" {
		author.NameEN = translit.ToLatin(author.NameRU)
	}

	query := `
		INSERT INTO authors (name_ru, name_en, bio, photo_url)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.QueryRow(query,
		author.NameRU,
		author.NameEN,
		author.Bio,
		author.PhotoURL,
	).Scan(&author.ID)

	return err
}

func (r *authorRepository) UpdateAuthor(author *models.Author) error {
	query := `
		UPDATE authors
		SET name_ru = $1, name_en = $2, bio = $3, photo_url = $4
		WHERE id = $5
	`
	_, err := r.db.Exec(query,
		author.NameRU,
		author.NameEN,
		author.Bio,
		author.PhotoURL,
		author.ID,
	)
	return err
}

func (r *authorRepository) DeleteAuthor(id int) error {
	_, err := r.db.Exec(`DELETE FROM authors WHERE id = $1`, id)
	return err
}

func (r *authorRepository) SearchAuthorByName(query string, limit, offset int) ([]*models.Author, error) {
	rows, err := r.db.Query(`
		SELECT id, name_ru, name_en, bio, photo_url
		FROM authors
		WHERE name_ru ILIKE '%' || $1 || '%' OR name_en ILIKE '%' || $1 || '%'
		ORDER BY name_ru ASC
		LIMIT $2 OFFSET $3
	`, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var authors []*models.Author
	for rows.Next() {
		var a models.Author
		err := rows.Scan(&a.ID, &a.NameRU, &a.NameEN, &a.Bio, &a.PhotoURL)
		if err != nil {
			return nil, err
		}
		authors = append(authors, &a)
	}
	return authors, nil
}

func (r *authorRepository) CountAuthors(query string) (int, error) {
	var count int
	err := r.db.QueryRow(`
		SELECT COUNT(*)
		FROM authors
		WHERE name_ru ILIKE '%' || $1 || '%' OR name_en ILIKE '%' || $1 || '%'
	`, query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
