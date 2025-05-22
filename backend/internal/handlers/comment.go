package handlers

import "github.com/gin-gonic/gin"

func GetCommentsForBook(c *gin.Context) {
	c.JSON(200, gin.H{"message": "comments for book " + c.Param("id")})
}

func CreateComment(c *gin.Context) {
	c.JSON(201, gin.H{"message": "comment created"})
}
