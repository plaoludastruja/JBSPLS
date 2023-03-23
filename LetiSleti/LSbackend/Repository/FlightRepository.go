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
	filter := bson.M{}
	if searchCriteria.StartPlace != "" {
		if searchCriteria.EndPlace != "" {
			filter = bson.M{"startPlace": searchCriteria.StartPlace, "endPlace": searchCriteria.EndPlace}
		} else {
			filter = bson.M{"startPlace": searchCriteria.StartPlace}
		}
	} else {
		if searchCriteria.EndPlace != "" {
			filter = bson.M{"endPlace": searchCriteria.EndPlace}
		}
	}
	cursor, err := flightsCollection.Find(context.TODO(), filter)
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

func DeleteFlight(flight Models.Flight) bool {

	_, err := flightsCollection.DeleteOne(context.TODO(), bson.M{"id": flight.Id})
	if err != nil {
		fmt.Println("Could not delete flight")
		fmt.Println(err)
		log.Panic("Could not delete flight", err.Error())
		return false
	}
	return true
}
