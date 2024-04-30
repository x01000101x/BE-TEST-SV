package controllers

import (
	"fmt"
	"go-crud/initializers"
	"go-crud/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ArticlesCreate(c *gin.Context) {

	var body struct {
		Title    string `json:"Title" binding:"required,min=20"`
		Content  string `json:"Content" binding:"required,min=200"`
		Category string `json:"Category" binding:"required,min=3"`
		Status   string `json:"Status" binding:"required,oneof=publish draft trash"`
	}

	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	articles := models.Articles{
		Title:    body.Title,
		Content:  body.Content,
		Category: body.Category,
		Status:   body.Status,
	}

	result := initializers.DB.Create(&articles)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Err result"})
		return
	}

	c.JSON(200, gin.H{
		"message": articles,
	})
}

func ArticlesId(c *gin.Context) {

	id := c.Param("id")

	var article models.Articles
	if err := initializers.DB.First(&article, id).Error; err != nil {
		fmt.Println("Error adalah = ", err)
		c.JSON(404, gin.H{"error": "Article not found"})
		return
	}

	c.JSON(200, gin.H{
		"message": article,
	})
}

func ArticlesShow(c *gin.Context) {

	limitStr := c.Param("limit")
	offsetStr := c.Param("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid offset parameter"})
		return
	}

	var articles []models.Articles
	err = initializers.DB.Offset(offset).Limit(limit).Find(&articles).Error
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": "Failed to retrieve articles"})
		return
	}

	c.JSON(200, gin.H{
		"message": articles,
	})
}

func ArticlesUpdate(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		Title    string `json:"Title" binding:"required,min=20"`
		Content  string `json:"Content" binding:"required,min=200"`
		Category string `json:"Category" binding:"required,min=3"`
		Status   string `json:"Status" binding:"required,oneof=publish draft trash"`
	}

	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var articles models.Articles

	err := initializers.DB.First(&articles, id).Error
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": "Failed to retrieve articles"})
		return
	}

	initializers.DB.Model(&articles).Updates(models.Articles{
		Title:    body.Title,
		Content:  body.Content,
		Category: body.Category,
		Status:   body.Status,
	})

	c.JSON(200, gin.H{
		"message": articles,
	})
}

func ArticlesDelete(c *gin.Context) {

	id := c.Param("id")

	if err := initializers.DB.Delete(&models.Articles{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete article"})
		return
	}
	c.JSON(200, gin.H{
		"message": "Sukses menghapus",
	})
}
