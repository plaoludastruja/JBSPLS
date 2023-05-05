package config

import "os"

type Config struct {
	Port               string
	AccomodationDBHost string
	AccomodationDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:               os.Getenv("ACCOMODATION_SERVICE_PORT"),
		AccomodationDBHost: os.Getenv("ACCOMODATION_DB_HOST"),
		AccomodationDBPort: os.Getenv("ACCOMODATION_DB_PORT"),
	}
}
