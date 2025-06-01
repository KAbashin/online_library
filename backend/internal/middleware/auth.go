package middleware

import (
	"fmt"
	"net/http"
	"online_library/backend/internal/pkg/auth"
	"online_library/backend/internal/pkg/roles"
	"online_library/backend/internal/service"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	userService service.UserService
}

func NewAuthMiddleware(userService service.UserService) *AuthMiddleware {
	return &AuthMiddleware{userService: userService}
}

func (m *AuthMiddleware) AuthRequired() gin.HandlerFunc {
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

		user, err := m.userService.GetByID(c, claims.UserID)
		if err != nil || user == nil || !user.Is_active {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
			return
		}

		if user.TokenVersion != claims.TokenVersion {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "session expired"})
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, ok := c.Get("role")
		role, _ := roleVal.(string)
		if !ok || (role != roles.RoleAdmin && role != roles.RoleSuperAdmin) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "admin access required"})
			return
		}
		c.Next()
	}
}

func SuperAdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != roles.RoleSuperAdmin {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "superadmin access required"})
			return
		}
		c.Next()
	}
}

func OwnerOrAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDFromToken, ok := c.Get("userID")
		roleVal, _ := c.Get("role")
		role, _ := roleVal.(string)

		paramID := c.Param("userID") // string
		if !ok || (fmt.Sprintf("%v", userIDFromToken) != paramID && !roles.IsAdmin(role)) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "access denied"})
			return
		}
		c.Next()
	}
}
