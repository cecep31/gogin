package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connectdb sadasd
func Connectdb() *gorm.DB {

	dsn := "pilput:pilput31@tcp(47.254.215.121:3306)/gogin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&Book{})
	return db

}
