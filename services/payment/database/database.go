package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(dataSourceName string) (*sql.DB, error) {
	// Try to connect to the database
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("DB connected")

	// Check if the database and table exist
	if err := checkDatabases(db); err != nil {
		return nil, err
	}

	return db, nil
}

func checkDatabases(db *sql.DB) error {
	// Check if the database exists
	_, err := db.Exec("SELECT 1 FROM pg_database WHERE datname='auth'")
	if err != nil {
		return err
	}

	// Check if the table exists
	_, err = db.Exec("SELECT 1 FROM pg_tables WHERE tablename='stocks'")
	if err != nil {
		return err
	}

	return nil
}
