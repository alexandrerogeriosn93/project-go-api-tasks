package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func initDB(filepath string) *sql.DB {
	db, error := sql.Open("sqlite", filepath)

	if error != nil {
		log.Fatal(error)
	}

	if db == nil {
		log.Fatal("Database is nil")
	}

	createTable(db)
	return db
}

func createTable(db *sql.DB) {
	sqlTable := `CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
	);`

	_, error := db.Exec(sqlTable)

	if error != nil {
		log.Fatalf("Failed to create table: %v", error)
	}
}
