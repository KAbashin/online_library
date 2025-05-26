package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online_library/backend/internal/service"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{Service: s}
}

type registerRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Bio      string `json:"bio"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	token, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if req.Bio == "" {
		req.Bio = "Пока нет описания"
	}

	if exists, _ := h.Service.UserService.CheckEmailExists(req.Email); exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
		return
	}

	err := h.Service.Register(req.Email, req.Name, req.Password, req.Bio)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "registered"})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
