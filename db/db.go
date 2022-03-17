package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

var DB = initDb()

func initDb() *sqlx.DB {
	database, err := sqlx.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql error", err)
		panic("open mysql error")
	}
	return database
}
