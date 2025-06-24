package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online_library/backend/internal/dto"
	"online_library/backend/internal/models"
	"online_library/backend/internal/service"
	"strconv"
	"strings"
)

type TagHandler struct {
	tagService service.TagService
}

func NewTagHandler(tagService service.TagService) *TagHandler {
	return &TagHandler{tagService: tagService}
}

func (h *TagHandler) SearchTags(c *gin.Context) {
	query := strings.TrimSpace(c.Query("query"))

	tags, err := h.tagService.GetAllTags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tags"})
		return
	}

	var filtered []models.Tag
	if query != "" {
		queryLower := strings.ToLower(query)
		for _, tag := range tags {
			if strings.Contains(strings.ToLower(tag.Name), queryLower) {
				filtered = append(filtered, tag)
			}
		}
	} else {
		filtered = tags
	}

	c.JSON(http.StatusOK, dto.ConvertTags(filtered))
}

func (h *TagHandler) GetTagByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	tag, err := h.tagService.GetTagByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tag not found"})
		return
	}

	c.JSON(http.StatusOK, dto.TagDTO{
		ID:    tag.ID,
		Name:  tag.Name,
		Color: tag.Color,
	})
}

func (h *TagHandler) CreateTag(c *gin.Context) {
	var input dto.TagDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	tag := input.ToModel()
	if err := h.tagService.CreateTag(tag); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create tag"})
		return
	}

	input.ID = tag.ID
	c.JSON(http.StatusCreated, input)
}

func (h *TagHandler) UpdateTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input dto.TagDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	input.ID = id
	tag := input.ToModel()

	if err := h.tagService.UpdateTag(tag); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update tag"})
		return
	}

	c.JSON(http.StatusOK, input)
}

func (h *TagHandler) DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.tagService.DeleteTag(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete tag"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *TagHandler) GetTagsByBookID(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("bookID"))

	tags, err := h.tagService.GetTagsByBookID(bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tags"})
		return
	}

	c.JSON(http.StatusOK, dto.ConvertTags(tags))
}

func (h *TagHandler) AssignTagToBook(c *gin.Context) {
	var bt models.BookTag
	if err := c.ShouldBindJSON(&bt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	if err := h.tagService.AssignTagToBook(&bt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to assign tag"})
		return
	}
	c.Status(http.StatusOK)
}

func (h *TagHandler) RemoveTagFromBook(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Query("book_id"))
	tagID, _ := strconv.Atoi(c.Query("tag_id"))
	if err := h.tagService.RemoveTagFromBook(bookID, tagID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remove tag"})
		return
	}
	c.Status(http.StatusOK)
}
