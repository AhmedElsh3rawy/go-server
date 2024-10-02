package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	// postgresql://user:password@host:port/DB_name
	connStr := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", connStr+"?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
