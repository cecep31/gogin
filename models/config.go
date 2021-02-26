package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connectdb() *gorm.DB {
	dsn := "pilput:pilput31@tcp(127.0.0.1:3306)/gogin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&Book{})
	return db

}
