package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// SystemConfig : System configuration model
type SystemConfig struct {
	hostURL  string
	hostPort string
}

// DatabaseConfig : Database configuration model
type DatabaseConfig struct {
	URL      string
	port     string
	name     string
	user     string
	password string
}

// InitDatabaseConfiguration : Initialize the database configuration
func InitDatabaseConfiguration() *DatabaseConfig {
	viper.SetConfigName("config") // name of config file (without extension)
	// viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	configuration := &DatabaseConfig{
		viper.GetString("database.url"),
		viper.GetString("database.port"),
		viper.GetString("database.name"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
	}

	return configuration
}
