package config

import "os"

type Config struct {
	Port             string
	UserHost         string
	UserPort         string
	AccomodationHost string
	AccomodationPort string
	AppointmentHost  string
	AppointmentPort  string
}

func NewConfig() *Config {
	return &Config{
		Port:             os.Getenv("GATEWAY_PORT"),
		UserHost:         os.Getenv("USER_SERVICE_HOST"),
		UserPort:         os.Getenv("USER_SERVICE_PORT"),
		AccomodationHost: os.Getenv("ACCOMODATION_SERVICE_HOST"),
		AccomodationPort: os.Getenv("ACCOMODATION_SERVICE_PORT"),
		AppointmentHost:  os.Getenv("APPOINTMENT_SERVICE_HOST"),
		AppointmentPort:  os.Getenv("APPOINTMENT_SERVICE_PORT"),
	}
}
