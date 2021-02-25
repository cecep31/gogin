package main

import (
	"github.com/gin-gonic/gin"
)

func showuser(c *gin.Context) {
	var users []User
	results, err := dbs().Query("SELECT ID,nama FROM users")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var user User
	for results.Next() {
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.Nama)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		users = append(users, user)

	}

	dbs()
	// dbs().Create(&User{nama: "cecep", email: "asd@100", alamat: "rawang"})
	// dbs().First(&user, 1).Scan(&user)
	defer dbs().Close()
	c.JSON(200, users)
}
