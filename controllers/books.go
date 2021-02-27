package controllers

import (
	"net/http"

	"github.com/cecep31/gogin/models"
	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.Connectdb().Find(&books)

	c.JSON(http.StatusOK, gin.H{
		"buku": books,
	})

}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
	book := models.Book{Title: input.Title, Author: input.Author}
	models.Connectdb().Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
func FindBook(c *gin.Context) {
	var book models.Book
	idbook := c.Param("id")
	if err := models.Connectdb().Where("id = ?", idbook).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.Connectdb().Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.Connectdb().Model(&book).Updates(models.Book{Title: input.Title, Author: input.Author})

	c.JSON(http.StatusOK, gin.H{"data": book})
}
func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.Connectdb().Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.Connectdb().Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
