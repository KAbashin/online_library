package handlers

import "github.com/gin-gonic/gin"

func GetRootCategories(c *gin.Context) {
	c.JSON(200, gin.H{"message": "list of root categories"})
}

func GetCategoryChildren(c *gin.Context) {
	c.JSON(200, gin.H{"message": "children of category " + c.Param("id")})
}

func GetBooksByCategory(c *gin.Context) {
	c.JSON(200, gin.H{"message": "books in category " + c.Param("id")})
}
