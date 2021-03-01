package models

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"uniqIndex" json:"email"`
	Password string `gorm:"->;<;not nul" json:"-"`
	Token    string `gorm:"-" json:"token.omitemply"`
}
