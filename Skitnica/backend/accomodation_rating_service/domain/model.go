package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type AccomodationRating struct {
	Id             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email          string             `json:"email" bson:"email" binding:"required"`
	AccomodationId string             `json:"accomodationId" bson:"accomodationId" binding:"required"`
	Rating         int32              `json:"rating" bson:"rating" binding:"required"`
	Date           string             `json:"date" bson:"date" binding:"required"`
}
