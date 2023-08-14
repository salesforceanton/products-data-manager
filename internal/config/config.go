package config

import (
	"errors"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB   DatabaseConfig
	Port int
}

type DatabaseConfig struct {
	Uri      string
	Password string
	Username string
	DBName   string `envconfig:"DB_DATABASE_NAME"`
}

const (
	DB_CONFIG_ENV_PREFIX     = "db"
	SERVER_CONFIG_ENV_PREFIX = "server"
)

// Recieve configuration values from env variables
func InitConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("Error with config initialization")
	}

	var cfg Config
	if err := envconfig.Process(SERVER_CONFIG_ENV_PREFIX, &cfg); err != nil {
		return nil, errors.New("Error with config initialization")
	}

	if err := envconfig.Process(DB_CONFIG_ENV_PREFIX, &cfg.DB); err != nil {
		return nil, errors.New("Error with config initialization")
	}

	return &cfg, nil
}
