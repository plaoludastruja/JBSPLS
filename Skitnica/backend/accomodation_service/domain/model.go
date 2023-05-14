package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Accomodation struct {
	Id                  primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name                string             `json:"name" bson:"name" binding:"required"`
	Location            string             `json:"location" bson:"location" binding:"required"`
	Facilities          string             `json:"facilities" bson:"facilities" binding:"required"`
	MinNumberOfGuests   int32              `json:"minNumberOfGuests" bson:"minNumberOfGuests" binding:"required"`
	MaxNumberOfGuests   int32              `json:"maxNumberOfGuests" bson:"maxNumberOfGuests" binding:"required"`
	IsAutomaticApproved bool               `json:"isAutomaticApproved" bson:"isAutomaticApproved" binding:"required"`
	HostUsername        string             `json:"hostUsername" bson:"hostUsername" binding:"required"`
}
