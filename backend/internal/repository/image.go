package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"online_library/backend/internal/models"
)

type ImageRepository interface {
	// GetAllImagesByCategory возвращает все изображения книг,
	// которые принадлежат указанной категории и её подкатегориям.
	GetAllImagesByCategory(categoryID int) ([]models.BookImage, error)

	// GetImageByID возвращает изображение по его ID.
	// Если изображение не найдено, возвращает заглушку с URL "/static/images/placeholder.png".
	GetImageByID(id int) (models.BookImage, error)

	// CreateImage добавляет новое изображение книги и заполняет поле ID у модели.
	CreateImage(image *models.BookImage) error

	// UpdateImage обновляет данные существующего изображения.
	UpdateImage(image *models.BookImage) error

	// DeleteImage удаляет изображение по ID.
	DeleteImage(id int) error

	// GetImagesByBookID возвращает все изображения для книги с указанным ID,
	// отсортированные по полю order_index.
	GetImagesByBookID(bookID int) ([]models.BookImage, error)

	// GetCoverImageByBookID возвращает главное изображение книги (обычно с order_index = 1).
	// Если изображение не найдено, возвращается заглушка.
	GetCoverImageByBookID(bookID int) (models.BookImage, error)

	// ReorderImages обновляет порядок изображений книги согласно переданному массиву ID.
	// newOrder — это слайс image.ID в новом порядке. Порядок присваивается начиная с 1.
	ReorderImages(bookID int, newOrder []int) error

	// DeleteImagesByBookID удаляет все изображения, связанные с заданной книгой.
	DeleteImagesByBookID(bookID int) error
}

type imageRepo struct {
	db *sql.DB
}

func NewImageRepository(db *sql.DB) ImageRepository {
	return &imageRepo{db: db}
}

func (i *imageRepo) GetAllImagesByCategory(categoryID int) ([]models.BookImage, error) {
	query := `
		WITH RECURSIVE category_tree AS (
			SELECT id FROM categories WHERE id = $1
			UNION ALL
			SELECT c.id FROM categories c
			INNER JOIN category_tree ct ON c.parent_id = ct.id
		)
		SELECT bi.id, bi.book_id, bi.url, bi.order_index
		FROM book_categories bc
		JOIN book_images bi ON bc.book_id = bi.book_id
		WHERE bc.category_id IN (SELECT id FROM category_tree)
		ORDER BY bi.book_id, bi.order_index;
	`

	rows, err := i.db.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var images []models.BookImage
	for rows.Next() {
		var img models.BookImage
		if err := rows.Scan(&img.ID, &img.BookID, &img.URL, &img.OrderIndex); err != nil {
			return nil, err
		}
		images = append(images, img)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return images, nil
}

func (i *imageRepo) GetImageByID(id int) (models.BookImage, error) {
	query := `SELECT id, book_id, url, order_index FROM book_images WHERE id = $1`
	var img models.BookImage
	err := i.db.QueryRow(query, id).Scan(&img.ID, &img.BookID, &img.URL, &img.OrderIndex)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.BookImage{
				URL: "/static/images/placeholder.png",
			}, nil
		}
		return img, err
	}
	return img, nil
}

func (i *imageRepo) CreateImage(image *models.BookImage) error {
	query := `INSERT INTO book_images (book_id, url, order_index) VALUES ($1, $2, $3) RETURNING id`
	row := i.db.QueryRow(query, image.BookID, image.URL, image.OrderIndex)

	err := row.Scan(&image.ID)
	if err != nil {
		return fmt.Errorf("failed to insert book image: %w", err)
	}

	return nil
}

func (i *imageRepo) UpdateImage(image *models.BookImage) error {
	query := `UPDATE book_images SET book_id = $1, url = $2, order_index = $3 WHERE id = $4`
	_, err := i.db.Exec(query, image.BookID, image.URL, image.OrderIndex, image.ID)
	return err
}

func (i *imageRepo) DeleteImage(id int) error {
	query := `DELETE FROM book_images WHERE id = $1`
	_, err := i.db.Exec(query, id)
	return err
}

func (i *imageRepo) GetImagesByBookID(bookID int) ([]models.BookImage, error) {
	query := `SELECT id, book_id, url, order_index FROM book_images WHERE book_id = $1 ORDER BY order_index`
	rows, err := i.db.Query(query, bookID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var images []models.BookImage
	for rows.Next() {
		var img models.BookImage
		if err := rows.Scan(&img.ID, &img.BookID, &img.URL, &img.OrderIndex); err != nil {
			return nil, err
		}
		images = append(images, img)
	}
	return images, rows.Err()
}

func (i *imageRepo) GetCoverImageByBookID(bookID int) (models.BookImage, error) {
	query := `SELECT id, book_id, url, order_index FROM book_images WHERE book_id = $1 AND order_index = 1 LIMIT 1`
	var img models.BookImage
	err := i.db.QueryRow(query, bookID).Scan(&img.ID, &img.BookID, &img.URL, &img.OrderIndex)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.BookImage{
				URL: "/static/images/placeholder.png",
			}, nil
		}
		return img, err
	}
	return img, nil
}

func (i *imageRepo) ReorderImages(bookID int, newOrder []int) error {
	tx, err := i.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	for index, imageID := range newOrder {
		query := `UPDATE book_images SET order_index = $1 WHERE id = $2 AND book_id = $3`
		_, err := tx.Exec(query, index+1, imageID, bookID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (i *imageRepo) DeleteImagesByBookID(bookID int) error {
	query := `DELETE FROM book_images WHERE book_id = $1`
	_, err := i.db.Exec(query, bookID)
	return err
}
