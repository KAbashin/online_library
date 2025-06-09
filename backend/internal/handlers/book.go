package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online_library/backend/internal/middleware"
	"online_library/backend/internal/models"
	"online_library/backend/internal/pkg/roles"
	"online_library/backend/internal/service"
	"strconv"
)

type BookHandler struct {
	bookService service.BookService
}

type TagListRequest struct {
	TagIDs []int `json:"tag_ids"`
}

type AuthorListRequest struct {
	AuthorIDs []int `json:"author_ids"`
}

type StatusUpdateRequest struct {
	Status string `json:"status"`
}

func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	userID, userRole, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	bookID, err := h.bookService.CreateBook(&book, userRole, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": bookID})
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	userID, userRole, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if err := h.bookService.UpdateBook(&book, userID, userRole); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	userID, userRole, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	if err := h.bookService.DeleteBook(bookID, userID, userRole); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	userID, userRole, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	book, err := h.bookService.GetBookByID(bookID, userID, userRole)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) GetBooksByStatuses(c *gin.Context) {
	_, userRole, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	books, err := h.bookService.GetBooksByStatuses(userRole, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) SearchBooks(c *gin.Context) {
	_, userRole, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	query := c.Query("q")
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	books, err := h.bookService.SearchBooks(query, userRole, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetBooksByAuthor(c *gin.Context) {
	_, userRole, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	authorID, err := strconv.Atoi(c.Param("author_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author ID"})
		return
	}

	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	books, err := h.bookService.GetBooksByAuthor(authorID, userRole, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetBooksByTag(c *gin.Context) {
	_, userRole, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	tagID, err := strconv.Atoi(c.Param("tag_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tag ID"})
		return
	}

	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	books, err := h.bookService.GetBooksByTag(tagID, userRole, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetUserBooks(c *gin.Context) {
	userID, _, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	books, err := h.bookService.GetUserBooks(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetUserFavoriteBooks(c *gin.Context) {
	userID, userRole, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	books, err := h.bookService.GetUserFavoriteBooks(userID, userRole)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) AddBookToFavorites(c *gin.Context) {
	userID, _, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	if err := h.bookService.AddBookToFavorites(userID, bookID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *BookHandler) RemoveBookFromFavorites(c *gin.Context) {
	userID, _, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	if err := h.bookService.RemoveBookFromFavorites(userID, bookID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *BookHandler) SetBookTags(c *gin.Context) {
	userID, userRole, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	var req TagListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if err := h.bookService.SetBookTags(bookID, req.TagIDs, userID, userRole); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *BookHandler) AddBookTag(c *gin.Context) {
	userID, userRole, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	bookID, _ := strconv.Atoi(c.Param("book_id"))
	tagID, _ := strconv.Atoi(c.Param("tag_id"))

	if err := h.bookService.AddBookTag(bookID, tagID, userID, userRole); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *BookHandler) RemoveBookTag(c *gin.Context) {
	userID, userRole, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	bookID, _ := strconv.Atoi(c.Param("book_id"))
	tagID, _ := strconv.Atoi(c.Param("tag_id"))

	if err := h.bookService.RemoveBookTag(bookID, tagID, userID, userRole); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *BookHandler) SetBookAuthors(c *gin.Context) {
	userID, userRole, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	bookID, _ := strconv.Atoi(c.Param("book_id"))

	var req AuthorListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if err := h.bookService.SetBookAuthors(bookID, req.AuthorIDs, userID, userRole); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *BookHandler) AddBookAuthor(c *gin.Context) {
	userID, userRole, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	bookID, _ := strconv.Atoi(c.Param("book_id"))
	authorID, _ := strconv.Atoi(c.Param("author_id"))

	if err := h.bookService.AddBookAuthor(bookID, authorID, userID, userRole); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *BookHandler) RemoveBookAuthor(c *gin.Context) {
	userID, userRole, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	bookID, _ := strconv.Atoi(c.Param("book_id"))
	authorID, _ := strconv.Atoi(c.Param("author_id"))

	if err := h.bookService.RemoveBookAuthor(bookID, authorID, userID, userRole); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *BookHandler) UpdateBookStatus(c *gin.Context) {
	_, userRole, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	bookID, _ := strconv.Atoi(c.Param("book_id"))

	var req StatusUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if err := h.bookService.UpdateBookStatus(bookID, req.Status, userRole); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *BookHandler) GetDuplicateBooks(c *gin.Context) {
	_, role, ok := middleware.ExtractUser(c)
	if !ok || !roles.IsAdmin(role) {
		c.JSON(http.StatusForbidden, gin.H{"error": "admin access required"})
		return
	}

	title := c.Param("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing title"})
		return
	}

	books, err := h.bookService.GetDuplicateBooks(title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find duplicates"})
		return
	}

	c.JSON(http.StatusOK, books)
}
