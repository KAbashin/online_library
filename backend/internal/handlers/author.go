package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online_library/backend/internal/dto"
	"online_library/backend/internal/models"
	"online_library/backend/internal/service"
	"strconv"
)

type AuthorHandler struct {
	service service.AuthorServiceInterface
}

func NewAuthorHandler(service service.AuthorServiceInterface) *AuthorHandler {
	return &AuthorHandler{service: service}
}

func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var aDto dto.AuthorDTO
	if err := c.ShouldBindJSON(&aDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	author := &models.Author{
		NameRU:   aDto.NameRU,
		NameEN:   aDto.NameEN,
		Bio:      aDto.Bio,
		PhotoURL: aDto.PhotoURL,
	}

	if err := h.service.CreateAuthor(author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	aDto.ID = author.ID
	c.JSON(http.StatusCreated, aDto)
}

func (h *AuthorHandler) UpdateAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author ID"})
		return
	}

	var aDto dto.AuthorDTO
	if err := c.ShouldBindJSON(&aDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	author := &models.Author{
		ID:       id,
		NameRU:   aDto.NameRU,
		NameEN:   aDto.NameEN,
		Bio:      aDto.Bio,
		PhotoURL: aDto.PhotoURL,
	}

	if err := h.service.UpdateAuthor(author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aDto.ID = id
	c.JSON(http.StatusOK, aDto)
}

func (h *AuthorHandler) DeleteAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author ID"})
		return
	}

	if err := h.service.DeleteAuthor(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete author"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *AuthorHandler) GetAuthorByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author ID"})
		return
	}

	author, err := h.service.GetAuthorByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "author not found"})
		return
	}

	c.JSON(http.StatusOK, author)
}

func (h *AuthorHandler) ListAuthors(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	query := c.Query("query")

	if query != "" {
		authors, count, err := h.service.SearchAuthors(query, limit, offset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to search authors"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": authors, "count": count})
		return
	}

	authors, err := h.service.GetAllAuthors(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch authors"})
		return
	}

	c.JSON(http.StatusOK, authors)
}
