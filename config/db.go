package config

import (
	"database/sql"
	"log"
	// Driver not needed for import.
	_ "github.com/lib/pq"
)

// TODO : DI
// Database : The pointer to the sql.DB
var Database *sql.DB

func init() {

	config := InitDatabaseConfiguration()

	connectionString := "postgres://"
	connectionString += config.user
	connectionString += ":"
	connectionString += config.password
	connectionString += "@"
	connectionString += config.URL
	connectionString += ":"
	connectionString += config.port
	connectionString += "/"
	connectionString += config.name
	connectionString += "?sslmode=disable"
	log.Println(connectionString)

	var err error
	Database, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	if err = Database.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database connection successful.")
}
