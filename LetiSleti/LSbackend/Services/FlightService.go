package Services

import (
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models/DTO"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Repository"
)

func AddFlight(flight Models.Flight) bool {
	return Repository.CreateFlight(flight)
}

func GetAllFlights() []Models.Flight {
	return Repository.GetAllFlights()
}

func SearchFlights(searchCriteria DTO.SearchDTO) []Models.Flight {
	return Repository.SearchFlights(searchCriteria)
}

func DeleteFlight(flightId string) int64 {
	return Repository.DeleteFlight(flightId)
}
func ChangePlacesLeft(id string) {
	Repository.ChangePlacesLeft(id)
}
