package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "asik")
	})
	r.GET("/user", showuser)

	r.Run()

}

func showaja() {

}
