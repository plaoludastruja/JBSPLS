package Models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Flight struct {
	Id                bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Start             time.Time     `json:"startTime" bson:"start_time" binding:"required"`
	StartPlace        string        `json:"startPlace" bson:"start_place" binding:"required"`
	End               time.Time     `json:"endTime" bson:"end_time" binding:"required"`
	EndPlace          string        `json:"endPlace" bson:"end_place" binding:"required"`
	MaxNumberOfPlaces int           `json:"maxNumberOfPlaces" bson:"max_number_of_places" binding:"required"`
	PricePerPlace     int           `json:"pricePErPlace" bson:"price_per_place" binding:"required"`
}
