package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	// Driver Current database driver
	Driver = "sqlite3"
	// SqliteDbAddress  Address of the database
	SqliteDbAddress = "C:\\Users\\jolash\\Dev\\Database\\Quotes.db"
)

// ExecDB Execute an SQL statement
func ExecDB(sqlStatement string, args ...interface{}) (sql.Result, error) {
	db := DBConnect()
	defer db.Close()

	result, err := db.Exec(sqlStatement, args...)
	return result, err
}

// QueryDB Query the database for something
func QueryDB(sqlStatement string) (*sql.Rows, error) {
	db := DBConnect()
	defer db.Close()

	rows, err := db.Query(sqlStatement)
	return rows, err
}

// DBConnect Connect to the database
func DBConnect() *sql.DB {
	db, err := sql.Open(Driver, SqliteDbAddress)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
