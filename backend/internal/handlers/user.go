package handlers

import "github.com/gin-gonic/gin"

func GetUsers(c *gin.Context) {
	// заглушка
	c.JSON(200, gin.H{"users": []string{"Alice", "Bob"}})
}

func CreateUser(c *gin.Context) {
	// заглушка
	c.JSON(201, gin.H{"message": "User created"})
}
