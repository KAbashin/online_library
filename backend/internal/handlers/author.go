package handlers

import "github.com/gin-gonic/gin"

func GetAuthorByID(c *gin.Context) {
	c.JSON(200, gin.H{"message": "author details for " + c.Param("id")})
}

func SearchAuthors(c *gin.Context) {
	c.JSON(200, gin.H{"message": "author search", "query": c.Request.URL.Query()})
}
