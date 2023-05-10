package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config represents the main configuration elements needed for the app
type Config struct {
	Server   Server
	Database Database
}

// Server refers to configuration needed for the server
type Server struct {
	Port string `envconfig:"SERVER_PORT"`
	Host string `envconfig:"SERVER_HOST"`
}

// Database refers to configuration needed for the DB
type Database struct {
	Name     string `envconfig:"DB_NAME"`
	Username string `envconfig:"DB_USERNAME"`
	Password string `envconfig:"DB_PASSWORD"`
	Host     string `envconfig:"DB_HOST"`
	Port     string `envconfig:"DB_PORT"`
}

// URL returns the string representation of the Database fields.
func (d *Database) URL() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", d.Username, d.Password, d.Host, d.Port, d.Name)
}

// Load builds the config from a given .env file
func Load(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	var c Config
	err = envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
