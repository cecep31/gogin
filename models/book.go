package models

import _ "gorm.io/gorm"

//Book model
type Book struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
