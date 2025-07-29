package sqliteconfig

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Init(debug bool) *sql.DB {
	dbPath := "./data/flight.db"
	if debug {
		dbPath = "." + dbPath
	}

	DB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	DB.SetMaxOpenConns(10)                 // Max number of open connections (reads)
	DB.SetMaxIdleConns(5)                  // Max idle connections in pool
	DB.SetConnMaxLifetime(5 * time.Minute) // Max lifetime of a connection

	if err := DB.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}

	return DB
}
