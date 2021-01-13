package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() *sql.DB {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3300)/go_db")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return db
}
