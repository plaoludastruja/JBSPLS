package main

import (
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/api_gateway/startup"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/api_gateway/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
