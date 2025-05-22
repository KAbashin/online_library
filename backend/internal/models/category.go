package models

import "context"

type Category struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	ParentID *int       `json:"parent_id"` // nil для корня
	Children []Category `json:"children,omitempty"`
}

// Получение дерева (рекурсивно)
func (r *CategoryRepository) GetTree(ctx context.Context) ([]Category, error) {
	var categories []Category
	// Запрос к БД: SELECT id, name, parent_id FROM categories
	// ... рекурсивно собираем дерево (можно через WITH RECURSIVE в SQL или кодом)

	/*
		sql = `WITH RECURSIVE category_tree AS (
		    SELECT id, name, parent_id
		    FROM categories
		    WHERE parent_id IS NULL  -- Корневые категории
		    UNION ALL
		    SELECT c.id, c.name, c.parent_id
		    FROM categories c
		    JOIN category_tree ct ON c.parent_id = ct.id
		)
		SELECT * FROM category_tree;`

	*/
	return categories, nil
}
