package conndb // import "conndb"

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	db *sql.DB
)

// ConnectToDb - hahaha
func ConnectToDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:13306)/test")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("db is connected")
	}
	return db, err
}