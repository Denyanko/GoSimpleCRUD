package controllers

import (
	"github.com/Denyanko/GoSimpleCRUD/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdatePostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreatePost(c *gin.Context) {
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "input is invalid. please fill all the fields"})
		return
	}

	post := models.Post{Title: input.Title, Content: input.Content}
	models.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

func FindPosts(c *gin.Context) {
	var posts []models.Post
	models.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func FindPost(c *gin.Context) {
	var post models.Post

	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "data is not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "data is not found"})
		return
	}

	var input UpdatePostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid data input. please fill all the fields"})
		return
	}

	UpdatedPost := models.Post{Title: input.Title, Content: input.Content}

	models.DB.Model(&post).Updates(&UpdatedPost)
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func DeletePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "data is not found"})
		return
	}

	models.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"data": "success"})
}
