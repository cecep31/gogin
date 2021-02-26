package controllers

import (
	"net/http"

	"github.com/cecep31/gogin/models"
	"github.com/gin-gonic/gin"
)

func Findbooks(c *gin.Context) {
	var books []models.Book
	models.Connectdb().Find(&books)

	c.JSON(http.StatusOK, gin.H{
		"buku": books,
	})

}
