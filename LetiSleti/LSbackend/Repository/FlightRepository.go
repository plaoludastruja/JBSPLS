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

	//date := time.Date(searchCriteria.Date.Year(), searchCriteria.Date.Month(), searchCriteria.Date.Year(), searchCriteria.Date.Day(), searchCriteria.Date.Hour(), searchCriteria.Date.Minute(), searchCriteria.Date.Second(), time.UTC)
	filter := bson.M{}
	if searchCriteria.StartPlace != "" && searchCriteria.EndPlace != "" && !searchCriteria.Date.IsZero() {
		filter = bson.M{"startPlace": searchCriteria.StartPlace, "endPlace": searchCriteria.EndPlace, "start": bson.M{"$gte": searchCriteria.Date}, "remaining": bson.M{"$gte": searchCriteria.NumberOfPlaces}}
	} else if searchCriteria.StartPlace != "" && searchCriteria.EndPlace != "" {
		filter = bson.M{"startPlace": searchCriteria.StartPlace, "endPlace": searchCriteria.EndPlace, "remaining": bson.M{"$gte": searchCriteria.NumberOfPlaces}}
	} else if searchCriteria.StartPlace != "" && !searchCriteria.Date.IsZero() {
		filter = bson.M{"startPlace": searchCriteria.StartPlace, "start": bson.M{"$gte": searchCriteria.Date}, "remaining": bson.M{"$gte": searchCriteria.NumberOfPlaces}}
	} else if searchCriteria.EndPlace != "" && !searchCriteria.Date.IsZero() {
		filter = bson.M{"endPlace": searchCriteria.EndPlace, "start": bson.M{"$gte": searchCriteria.Date}, "remaining": bson.M{"$gte": searchCriteria.NumberOfPlaces}}
	} else if searchCriteria.StartPlace != "" {
		filter = bson.M{"startPlace": searchCriteria.StartPlace, "remaining": bson.M{"$gte": searchCriteria.NumberOfPlaces}}
	} else if !searchCriteria.Date.IsZero() {
		filter = bson.M{"start": bson.M{"$gte": searchCriteria.Date}, "remaining": bson.M{"$gte": searchCriteria.NumberOfPlaces}}
	} else if searchCriteria.EndPlace != "" {
		filter = bson.M{"endPlace": searchCriteria.EndPlace, "remaining": bson.M{"$gte": searchCriteria.NumberOfPlaces}}
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

	fmt.Println("SearchResult:")
	fmt.Println(&results)

	return results
}

func DeleteFlight(flightId string) int64 {

	objectId, err := primitive.ObjectIDFromHex(flightId)
	if err != nil {
		log.Println("Invalid id")
	}
	res, err := flightsCollection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: objectId}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	return res.DeletedCount
}

func ChangePlacesLeft(id string) error {

	var flight Models.Flight
	objID, _ := primitive.ObjectIDFromHex(id)
	err := flightsCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: objID}}).Decode(&flight)
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"remaining": flight.Remaining - 1,
	}}
	flightsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
