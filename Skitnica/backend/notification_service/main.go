package main

import (
	"fmt"

	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/notification_service/startup"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/notification_service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
	fmt.Println("a ...any")
}
