package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reservation struct {
	Id             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	AccomodationId string             `json:"accomodationId" bson:"accomodationId,omitempty"`
	Username       string             `json:"username" bson:"username,omitempty"`
	StartDate      time.Time          `json:"startDate" bson:"startDate" binding:"required"`
	EndDate        time.Time          `json:"endDate" bson:"endDate" binding:"required"`
	GuestNumber    int32              `json:"guestNumber" bson:"guestNumber" binding:"required"`
	//Status options: APPROVED, PENDING, REJECTED, CANCELED
	Status       string `json:"status" bson:"status" binding:"required"`
	HostUsername string `json:"hostUsername" bson:"hostUsername,omitempty"`
}

type DateRange struct {
	StartDate time.Time `json:"startDate" bson:"startDate" binding:"required"`
	EndDate   time.Time `json:"endDate" bson:"endDate" binding:"required"`
}

type ReservationDto struct {
	Id              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	AccomodationId  string             `json:"accomodationId" bson:"accomodationId,omitempty"`
	Username        string             `json:"username" bson:"username,omitempty"`
	StartDate       time.Time          `json:"startDate" bson:"startDate" binding:"required"`
	EndDate         time.Time          `json:"endDate" bson:"endDate" binding:"required"`
	GuestNumber     int32              `json:"guestNumber" bson:"guestNumber" binding:"required"`
	CancellationNum int32              `json:"cancellationNum" bson:"cancellationNum" binding:"required"`
}
