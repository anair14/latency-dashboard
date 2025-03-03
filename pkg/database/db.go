package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the database and creates tables if necessary
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./database/data.db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	createTables()
}

func createTables() {
	usersTable := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );`

	endpointsTable := `CREATE TABLE IF NOT EXISTS endpoints (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL,
        url TEXT NOT NULL,
        threshold INTEGER NOT NULL DEFAULT 500,
        FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
    );`

	if _, err := DB.Exec(usersTable); err != nil {
		log.Fatal("Error creating users table:", err)
	}

	if _, err := DB.Exec(endpointsTable); err != nil {
		log.Fatal("Error creating endpoints table:", err)
	}
}

