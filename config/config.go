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
	Port     string
	Name     string
	User     string
	Password string
	Options  []string `mapstructure:"options"`
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

	var config DatabaseConfig
	err = viper.UnmarshalKey("database", &config)
	if err != nil {
		panic("Unable to unmarshal configuration")
	}

	return &config
}
