package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/mydatabase")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}