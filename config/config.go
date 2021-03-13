package config

import (
	"fmt"
	"log"
	"os"

	"github.com/cecep31/gogin/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connectdb sadasd
func Connectdb() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	user := os.Getenv("db_user")
	pass := os.Getenv("db_pass")
	dbname := os.Getenv("db_name")
	host := os.Getenv("db_host")
	port := os.Getenv("db_port")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&models.Book{})

	return db

}
