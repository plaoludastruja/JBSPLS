package Repository

import (
	"context"
	"fmt"
	"log"

	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models/DTO"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateFlight(flight Models.Flight) bool {
	if _, err := flightsCollection.InsertOne(context.TODO(), flight); err != nil {
		log.Panic("Could not save document to database", err.Error())
	}
	return true
}

func GetAllFlights() []Models.Flight {
	results := []Models.Flight{}
	cursor, _ := flightsCollection.Find(context.TODO(), bson.D{})
	cursor.All(context.TODO(), &results)
	return results
}

func SearchFlights(searchCriteria DTO.SearchDTO) []Models.Flight {
	results := []Models.Flight{}
	cursor, err := flightsCollection.Find(context.TODO(), bson.M{"startPlace": searchCriteria.StartPlace, "endPlace": searchCriteria.EndPlace})
	if err != nil {
		log.Panic("Could not find document in database", err.Error())
		return nil
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Panic("Could not find document in database", err.Error())
		return nil
	}
	fmt.Println("search criteria:")
	fmt.Println(searchCriteria.StartPlace)
	fmt.Println("SearchResult:")
	fmt.Println(&results)
	return results
}
