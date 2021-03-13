package main

import (
	"github.com/cecep31/gogin/config"
	"github.com/cecep31/gogin/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// connect databases
	config.Connectdb()

	r.GET("/book", controllers.FindBooks)
	r.POST("/book", controllers.CreateBook)
	r.GET("/book/:id", controllers.FindBook)
	r.PATCH("/book/:id", controllers.UpdateBook)
	r.DELETE("/book/:id", controllers.DeleteBook)
	r.Run()

}
