package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func dbs() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/gogin")
	if err != nil {
		panic(err.Error())
	}
	return db
}
