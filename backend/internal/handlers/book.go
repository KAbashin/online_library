package handlers

import "strconv"

/*
Пагинация книг в категории
*/
func (h *BookHandler) GetBooksByCategory(c echo.Context) error {
	categoryID := c.Param("id")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	limit := 20 // Книг на страницу

	books, total, err := h.service.GetBooksByCategory(categoryID, page, limit)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"books": books,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}
