package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func showuser(c *gin.Context) {
	results, err := dbs().Query("SELECT id, nama FROM users where id = 1")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var user User
	for results.Next() {
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.id, &user.nama)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Println(user.nama)
	}
	dbs()
	// dbs().Create(&User{nama: "cecep", email: "asd@100", alamat: "rawang"})
	// dbs().First(&user, 1).Scan(&user)
	defer dbs().Close()
	c.JSON(200, gin.H{
		"message": user.nama,
	})
}
