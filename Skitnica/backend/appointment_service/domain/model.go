package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Appointment struct {
	Id             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	AccomodationId string             `json:"accomodationdId" bson:"accomodationId" binding:"required"`
	Start          time.Time          `json:"start" bson:"start" binding:"required"`
	End            time.Time          `json:"end" bson:"end" binding:"required"`
	PriceType      string             `json:"priceType" bson:"priceType" binding:"required"`
	Price          int32              `json:"price" bson:"price" binding:"required"`
	Status         string             `json:"status" bson:"status" binding:"required"`
}
