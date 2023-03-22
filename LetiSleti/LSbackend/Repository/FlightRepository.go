package Repository

import (
	"context"
	"fmt"
	"log"

	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models/DTO"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func DeleteFlight(id string) bool {
	objID, er := primitive.ObjectIDFromHex(id)
	if er != nil {
		fmt.Println("Could not convert id")
		fmt.Println(er)
		log.Panic("Could not convert id", er.Error())
		return false
	}
	_, err := flightsCollection.DeleteOne(context.TODO(), bson.M{"id": objID})
	if err != nil {
		fmt.Println("Could not delete flight")
		fmt.Println(err)
		log.Panic("Could not delete flight", err.Error())
		return false
	}
	return true
}
