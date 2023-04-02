package Models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type Ticket struct {
	Id         bson.ObjectId      `json:"id" bson:"_id,omitempty"`
	Email      string             `json:"email" bson:"email" binding:"required"`
	FirstName  string             `json:"firstName" bson:"first_name" binding:"required"`
	LastName   string             `json:"lastName" bson:"last_name" binding:"required"`
	Start      string             `json:"start" bson:"start" binding:"required"`
	StartPlace string             `json:"startPlace" bson:"startPlace" binding:"required"`
	End        string             `json:"end"  bson:"end" binding:"required"`
	EndPlace   string             `json:"endPlace" bson:"endPlace" binding:"required"`
	Price      int                `json:"price" bson:"price" binding:"required"`
	Count      int                `json:"count" bson:"count" binding:"required"`
	FlightId   primitive.ObjectID `json:"flightId" bson:"flightId,omitempty"`
}
