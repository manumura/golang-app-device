package config

import (
	"database/sql"
	"fmt"
	// Not needed for import.
	_ "github.com/lib/pq"
)

var Database *sql.DB

func init() {
	var err error
	Database, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/mtx-web?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = Database.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database connection successful.")
}
