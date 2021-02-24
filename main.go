package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "asik")
	})
	r.GET("/api/user", showuser)
	r.GET("/api/product", showuser)
	r.POST("/api/product", showuser)
	r.DELETE("/api/product", showuser)

	r.Run()

}

func showaja() {

}
