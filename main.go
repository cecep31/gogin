package main

import (
	"github.com/cecep31/gogin/controllers"
	"github.com/cecep31/gogin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// connect databases
	models.Connectdb()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.Run()

}
