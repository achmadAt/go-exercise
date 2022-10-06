package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*Create mysql connection*/
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/todos")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}
