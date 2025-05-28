package middleware

import (
	"net/http"
	"online_library/backend/internal/models"
	"online_library/backend/internal/pkg/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
			return
		}
		tokenStr := strings.TrimPrefix(header, "Bearer ")

		claims, err := auth.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get("role")
		if !ok || (role != models.RoleAdmin && role != models.RoleSuperAdmin) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "admin access required"})
			return
		}
		c.Next()
	}
}

func SuperAdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != models.RoleSuperAdmin {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "superadmin access required"})
			return
		}
		c.Next()
	}
}
