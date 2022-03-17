package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB = initDb()

func initDb() *sqlx.DB {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3336)/homework")
	if err != nil {
		fmt.Println("open mysql error", err)
		panic("open mysql error")
	}
	return database
}
