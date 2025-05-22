package repository

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
