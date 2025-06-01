package middleware

import (
	"github.com/gin-gonic/gin"
)

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
