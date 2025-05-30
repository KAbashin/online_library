package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online_library/backend/internal/middleware"
	"online_library/backend/internal/models"
	"online_library/backend/internal/service"
	"strconv"
	"time"
)

const (
	DefaultLimit = "10"
	MaxLimit     = 100
)

type CommentHandler struct {
	service service.CommentService
}

func NewCommentHandler(service service.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	userID, _, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.UserID = userID
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = comment.CreatedAt
	comment.Status = models.CommentStatusActive // по умолчанию

	if err := h.service.Create(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

func (h *CommentHandler) UpdateComment(c *gin.Context) {
	userID, role, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid comment id"})
		return
	}

	input.ID = id

	if err := h.service.Update(&input, userID, role); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	userID, role, ok := middleware.ExtractUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	existing, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "comment not found"})
		return
	}

	isOwner := existing.UserID == userID
	isAdmin := role == models.RoleAdmin || role == models.RoleSuperAdmin
	if !isOwner && !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	if err := h.service.Delete(id, userID, role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *CommentHandler) GetCommentsByBook(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", DefaultLimit))

	if limit > MaxLimit {
		limit = MaxLimit
	}

	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	statuses := []string{models.CommentStatusActive}

	comments, err := h.service.GetByBookID(bookID, limit, offset, statuses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) GetCommentsByUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", DefaultLimit))

	if limit > MaxLimit {
		limit = MaxLimit
	}

	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	comments, err := h.service.GetByUserID(userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) GetLastComments(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", DefaultLimit))

	if limit > MaxLimit {
		limit = MaxLimit
	}

	comments, err := h.service.GetLast(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) SetStatus(c *gin.Context) {
	_, role, ok := middleware.ExtractUser(c)
	if !ok || (role != models.RoleAdmin && role != models.RoleSuperAdmin) {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var payload struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	if err := h.service.SetStatus(id, payload.Status, role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
