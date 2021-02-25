package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func dbs() *sql.DB {
	db, err := sql.Open("mysql", "pilput:pilput31@tcp(127.0.0.1:3306)/gogin?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}
