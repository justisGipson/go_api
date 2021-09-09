package database

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v4/stdlib" // pgx driver for PostgreSQL
)

// PostgreSQL Conn func
func PostgreSQLConnection() (*sqlx.DB, error) {
	// connection settings
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	// define db conn for pgsql
	db, err := sqlx.Connect("pgx", os.Getenv("DB_SERVER_URL"))
	if err != nil {
		return nil, fmt.Errorf("error, no connection to database: %e", err)
	}

	db.SetMaxOpenConns(maxConn)                           // default is 0/unlimited connections
	db.SetMaxIdleConns(maxIdleConn)                       // default is 2
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn)) // 0, connections are reused forever

	// attempt db ping
	if err := db.Ping(); err != nil {
		defer db.Close() // close db if conn throws error
		return nil, fmt.Errorf("connection error, cannot ping database: %e", err)
	}
	return db, nil
}
