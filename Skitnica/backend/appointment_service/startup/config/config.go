package config

import "os"

type Config struct {
	Port   string
	DBHost string
	DBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:   os.Getenv("APPOINTMENT_SERVICE_PORT"),
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
	}
}
