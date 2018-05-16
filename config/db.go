package config

import (
	"database/sql"
	"log"
	// Driver not needed for import.
	_ "github.com/lib/pq"
)

// Database : The pointer to the sql.DB
var Database *sql.DB

func init() {

	config := InitDatabaseConfiguration()

	connectionString := "postgres://"
	connectionString += config.User
	connectionString += ":"
	connectionString += config.Password
	connectionString += "@"
	connectionString += config.URL
	connectionString += ":"
	connectionString += config.Port
	connectionString += "/"
	connectionString += config.Name

	// options
	if len(config.Options) > 0 {
		connectionString += "?"
		for i, option := range config.Options {
			connectionString += option
			if i < len(config.Options)-1 {
				connectionString += "&"
			}
		}
	}
	// connectionString += "?sslmode=disable"

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
