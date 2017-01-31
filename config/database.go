package database

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() sql.DB {
	db, err := sql.Open("mysql", "root:@/golang")
	checkErr(err)
	return *db
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Panic WOY")
		panic(err)
	}
}
