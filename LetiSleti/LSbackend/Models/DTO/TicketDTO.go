package DTO

import "time"

type TicketDTO struct {
	Email      string    `json:"email" bson:"email" binding:"required"`
	FirstName  string    `json:"firstName" bson:"first_name" binding:"required"`
	LastName   string    `json:"lastName" bson:"last_name" binding:"required"`
	Start      time.Time `json:"startTime" bson:"start_time" binding:"required"`
	StartPlace string    `json:"startPlace" bson:"start_place" binding:"required"`
	End        time.Time `json:"endTime" bson:"end_time" binding:"required"`
	EndPlace   string    `json:"endPlace" bson:"end_place" binding:"required"`
	Price      int       `json:"price" bson:"price" binding:"required"`
	Count      int       `json:"count" bson:"count" binding:"required"`
}
