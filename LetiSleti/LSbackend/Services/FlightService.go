package Services

import (
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Repository"
)

func AddFlight(flight Models.Flight) bool {
	return Repository.CreateFlight(flight)
}
