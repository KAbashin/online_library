package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func (h *BookHandler) SearchBooks(c *gin.Context) {
	// --- tags ---
	var tagIDs []int
	if tagsParam := c.Query("tags"); tagsParam != "" {
		for _, s := range strings.Split(tagsParam, ",") {
			if id, err := strconv.Atoi(strings.TrimSpace(s)); err == nil {
				tagIDs = append(tagIDs, id)
			}
		}
	}

	// --- pagination ---
	limit := 20 // default
	offset := 0
	if l := c.Query("limit"); l != "" {
		if val, err := strconv.Atoi(l); err == nil {
			limit = val
		}
	}
	if o := c.Query("offset"); o != "" {
		if val, err := strconv.Atoi(o); err == nil {
			offset = val
		}
	}

	// --- sort ---
	sort := c.DefaultQuery("sort", "title_asc")

	// --- search ---
	books, err := h.bookService.SearchBooks(tagIDs, limit, offset, sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "search failed"})
		return
	}
	c.JSON(http.StatusOK, books)
}
