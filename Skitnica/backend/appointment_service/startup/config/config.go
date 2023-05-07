package config

import "os"

type Config struct {
	Port              string
	AppointmentDBHost string
	AppointmentDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:              os.Getenv("APPOINTMENT_SERVICE_PORT"),
		AppointmentDBHost: os.Getenv("APPOINTMENT_DB_HOST"),
		AppointmentDBPort: os.Getenv("APPOINTMENT_DB_PORT"),
	}
}
