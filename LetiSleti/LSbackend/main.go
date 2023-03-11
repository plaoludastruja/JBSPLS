package main

import (
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Repository"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Routes"
)

func main() {

	Repository.Setup()
	Routes.InitRoutes()
	Repository.Disconnect()

}
