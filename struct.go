package main

import "time"

type (
	// User
	User struct {
		ID        int       `json:"id"`
		Nama      string    `json:"nama"`
		Email     string    `json:"email"`
		Createdat time.Time `json:"created_at"`
		Updatedat time.Time `json:"updated_at"`
	}
)
type (
	// Mahasiswa
	Mahasiswa struct {
		ID        int       `json:"id"`
		NIM       int       `json:"nim"`
		Name      string    `name:"name"`
		Semester  int       `json:"semester"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
