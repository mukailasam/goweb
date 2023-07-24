package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func DatabaseConnection() *sql.DB {
	dbconn, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Panic(err)
	}

	return dbconn

}
