package config

import "os"

type Config struct {
	Port                     string
	DBHost                   string
	DBPort                   string
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	DeleteUserCommandSubject string
	DeleteUserReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                     os.Getenv("ACCOMODATION_SERVICE_PORT"),
		DBHost:                   os.Getenv("DB_HOST"),
		DBPort:                   os.Getenv("DB_PORT"),
		NatsHost:                 os.Getenv("NATS_HOST"),
		NatsPort:                 os.Getenv("NATS_PORT"),
		NatsUser:                 os.Getenv("NATS_USER"),
		NatsPass:                 os.Getenv("NATS_PASS"),
		DeleteUserCommandSubject: os.Getenv("DELETE_USER_COMMAND_SUBJECT"),
		DeleteUserReplySubject:   os.Getenv("DELETE_USER_REPLY_SUBJECT"),
	}
}
