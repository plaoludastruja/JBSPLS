package Models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Flight struct {
	Id                bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Start             time.Time     `json:"start" bson:"start" binding:"required"`
	StartPlace        string        `json:"startPlace" bson:"startPlace" binding:"required"`
	End               time.Time     `json:"end"  bson:"end" binding:"required"`
	EndPlace          string        `json:"endPlace" bson:"endPlace" binding:"required"`
	MaxNumberOfPlaces int           `json:"maxNumberOfPlaces" bson:"maxNumberOfPlaces" binding:"required"`
	PricePerPlace     int           `json:"pricePerPlace" bson:"pricePerPlace" binding:"required"`
}
