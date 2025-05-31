package repository

import (
	"database/sql"
	"online_library/backend/internal/models"
)

type CategoryRepository interface {
	GetAllCategories() ([]models.Category, error)
	GetRootCategories() ([]*models.Category, error)
	GetCategoryByID(id int) (*models.Category, error)
	GetCategoryChildren(parentID int) ([]*models.Category, error)
	GetBooksByCategoryIDRecursive(categoryID int) ([]*models.Book, error)
	CreateCategory(category *models.Category) (int, error)
	UpdateCategory(category *models.Category) error
	DeleteCategory(id int) error
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetAllCategories() ([]models.Category, error) {
	rows, err := r.db.Query(`SELECT id, name, parent_id, slug, description FROM categories`)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		if err := rows.Scan(&cat.ID, &cat.Name, &cat.ParentID, &cat.Slug, &cat.Description); err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}

	return categories, nil
}

func (r *categoryRepository) GetRootCategories() ([]*models.Category, error) {
	rows, err := r.db.Query(`SELECT id, name, parent_id, slug, description FROM categories WHERE parent_id IS NULL`)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)

	var result []*models.Category
	for rows.Next() {
		var cat models.Category
		if err := rows.Scan(&cat.ID, &cat.Name, &cat.ParentID, &cat.Slug, &cat.Description); err != nil {
			return nil, err
		}
		result = append(result, &cat)
	}
	return result, nil
}

func (r *categoryRepository) GetCategoryByID(id int) (*models.Category, error) {
	var category models.Category
	err := r.db.QueryRow(`
		SELECT id, name, parent_id, slug, description 
		FROM categories 
		WHERE id = $1`, id).
		Scan(&category.ID, &category.Name, &category.ParentID, &category.Slug, &category.Description)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) GetCategoryChildren(parentID int) ([]*models.Category, error) {
	rows, err := r.db.Query(`
		SELECT id, name, parent_id, slug, description 
		FROM categories 
		WHERE parent_id = $1`, parentID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var categories []*models.Category
	for rows.Next() {
		var cat models.Category
		if err := rows.Scan(&cat.ID, &cat.Name, &cat.ParentID, &cat.Slug, &cat.Description); err != nil {
			return nil, err
		}
		categories = append(categories, &cat)
	}
	return categories, nil
}

func (r *categoryRepository) GetBooksByCategoryIDRecursive(categoryID int) ([]*models.Book, error) {
	query := `
	WITH RECURSIVE subcategories AS (
		SELECT id FROM categories WHERE id = $1
		UNION ALL
		SELECT c.id
		FROM categories c
		INNER JOIN subcategories sc ON sc.id = c.parent_id
	)
	SELECT b.id, b.title, b.rating, b.cover_url
	FROM books b
	JOIN book_categories bc ON b.id = bc.book_id
	WHERE bc.category_id IN (SELECT id FROM subcategories);`

	rows, err := r.db.Query(query, categoryID)
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
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Rating, &book.CoverURL)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}

func (r *categoryRepository) CreateCategory(category *models.Category) (int, error) {
	var id int
	err := r.db.QueryRow(`
		INSERT INTO categories (name, parent_id, slug, description) 
		VALUES ($1, $2, $3, $4) RETURNING id`,
		category.Name, category.ParentID, category.Slug, category.Description,
	).Scan(&id)
	return id, err
}

func (r *categoryRepository) UpdateCategory(category *models.Category) error {
	_, err := r.db.Exec(`
		UPDATE categories SET name = $1, parent_id = $2, slug = $3, description = $4 
		WHERE id = $5`,
		category.Name, category.ParentID, category.Slug, category.Description, category.ID,
	)
	return err
}

func (r *categoryRepository) DeleteCategory(id int) error {
	_, err := r.db.Exec(`DELETE FROM categories WHERE id = $1`, id)
	return err
}
