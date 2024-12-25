package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// InitDB initializes and verifies a connection to a PostgresSQL database.
func InitDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to establish database connection: %w", err)
	}

	return db, nil
}
