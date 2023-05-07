package main

import (
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/appointment_service/startup"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/appointment_service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
