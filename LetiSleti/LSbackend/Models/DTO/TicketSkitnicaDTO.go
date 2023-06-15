package DTO

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketSkitnicaDTO struct {
	ApiKey     string             `json:"apiKey" bson:"apiKey"`
	Start      string             `json:"startTime" bson:"start_time"`
	StartPlace string             `json:"startPlace" bson:"start_place"`
	End        string             `json:"endTime" bson:"end_time"`
	EndPlace   string             `json:"endPlace" bson:"end_place"`
	Price      int                `json:"price" bson:"price"`
	Count      int                `json:"count" bson:"count"`
	FlightId   primitive.ObjectID `json:"flightId" bson:"flightId"`
}
