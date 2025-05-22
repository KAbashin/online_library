package handlers

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	c.JSON(200, gin.H{"token": "jwt_token_here"})
}

func Register(c *gin.Context) {
	c.JSON(201, gin.H{"message": "user registered"})
}
