package middleware

import (
	"github.com/gin-gonic/gin"
	"online_library/backend/internal/models"
)

func IsAdmin(role string) bool {
	return role == models.RoleAdmin || role == models.RoleSuperAdmin
}

func IsSuperAdmin(role string) bool {
	return role == models.RoleSuperAdmin
}

func ExtractUser(c *gin.Context) (int, string, bool) {
	userIDRaw, ok1 := c.Get("userID")
	roleRaw, ok2 := c.Get("role")

	userID, okID := userIDRaw.(int)
	role, okRole := roleRaw.(string)

	if !ok1 || !ok2 || !okID || !okRole {
		return 0, "", false
	}
	return userID, role, true
}
