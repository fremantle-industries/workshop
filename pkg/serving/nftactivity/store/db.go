package store

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const MaxLimit = 50

var db *sql.DB

func Open(dsn string) (*sql.DB, error) {
	odb, err := sql.Open("pgx", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open database: %v\n", err)
		return &sql.DB{}, err
	}
	db = odb

	return db, nil
}

func IngestDSN() string {
	dsn := os.Getenv("INGEST_DSN")
	return dsn
}

func QueryRow(query string, args ...any) *sql.Row {
	return db.QueryRow(query, args...)
}

func Query(query string, args ...any) (*sql.Rows, error) {
	return db.Query(query, args...)
}
