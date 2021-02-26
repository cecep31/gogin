package main

import (
	"github.com/cecep31/gogin/controllers"
	"github.com/cecep31/gogin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.Connectdb()
	r.GET("/books", controllers.Findbooks)
	r.Run()

}
