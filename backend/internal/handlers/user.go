package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online_library/backend/internal/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// админ / суперадмин
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// админ / суперадмин
func (h *UserHandler) СreateUser(c *gin.Context) {
	// заглушка
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

// владелец или админ/ суперадмин
func (h *UserHandler) UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

// только админы
func (h *UserHandler) SoftDeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

// Только для суперадмина
func (h *UserHandler) HardDeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
