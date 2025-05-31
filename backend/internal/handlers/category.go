package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online_library/backend/internal/models"
	"online_library/backend/internal/service"
	"strconv"
)

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(s service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s}
}

func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	tree, err := h.service.GetCategoryTree()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения категорий"})
		return
	}
	c.JSON(http.StatusOK, tree)
}

func (h *CategoryHandler) GetRootCategories(c *gin.Context) {
	root, err := h.service.GetCategoryRoot()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить корневые категории"})
		return
	}
	c.JSON(http.StatusOK, root)
}

func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	id, errC := strconv.Atoi(c.Param("id"))
	if errC != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category id"})
		return
	}
	cat, err := h.service.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Категория не найдена" + c.Param("id")})
		return
	}
	c.JSON(http.StatusOK, cat)
}

func (h *CategoryHandler) GetCategoryChildren(c *gin.Context) {
	id, errC := strconv.Atoi(c.Param("id"))
	if errC != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category id"})
		return
	}
	children, err := h.service.GetCategoryChildren(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения подкатегорий" + c.Param("id")})
		return
	}
	c.JSON(http.StatusOK, children)
}

func (h *CategoryHandler) GetBooksInCategory(c *gin.Context) {
	id, errC := strconv.Atoi(c.Param("id"))
	if errC != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category id"})
		return
	}
	books, err := h.service.GetBooksByCategoryIDRecursive(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения книг" + c.Param("id")})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ввод"})
		return
	}

	id, err := h.service.CreateCategory(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать категорию"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ввод"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}
	input.ID = id

	if err := h.service.UpdateCategory(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Категория обновлена"})
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}
	if err := h.service.DeleteCategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Категория удалена"})
}
