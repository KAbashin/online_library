package handlers

/*
import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"online_library/backend/internal/models"
	"online_library/backend/internal/repository"
)


type BookHandler struct {
	Repo *repository.Queries
	DB   *sql.DB
}

func NewBookHandler(db *sql.DB) *BookHandler {
	return &BookHandler{
		Repo: repository.New(db),
		DB:   db,
	}
}

func GetBookByID(c *gin.Context) {
	c.JSON(200, gin.H{"message": "book details for " + c.Param("id")})
}

func SearchBooks(c *gin.Context) {
	c.JSON(200, gin.H{"message": "search books", "query": c.Request.URL.Query()})
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var req models.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	book, err := h.Repo.CreateBook(c, repository.CreateBookParams{
		Title:       req.Title,
		Description: req.Description,
		PublishYear: req.PublishYear,
		Pages:       req.Pages,
		Language:    req.Language,
		Publisher:   req.Publisher,
		Type:        req.Type,
		CoverUrl:    req.CoverURL,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании книги"})
		return
	}

	c.JSON(http.StatusCreated, book)
}
*/
