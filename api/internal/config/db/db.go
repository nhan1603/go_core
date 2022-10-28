package db

import (
	"database/sql"

	"github.com/friendsofgo/errors"
	_ "github.com/lib/pq" // driver postgres
)

// ConnectDB returns a database connection
func ConnectDB(driver string, dbURL string) (*sql.DB, error) {
	if dbURL == "" {
		return nil, errors.New("Invalid dbURL")
	}

	conn, err := sql.Open(driver, dbURL)
	if err != nil {
		return nil, err
	}

	// If no error is returned, we can assume a successful connection
	if err := conn.Ping(); err != nil {
		return nil, err
	}
	return conn, nil
}
