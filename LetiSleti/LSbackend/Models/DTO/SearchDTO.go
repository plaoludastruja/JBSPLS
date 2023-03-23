package DTO

import "time"

type SearchDTO struct {
	StartPlace string    `json:"startPlace" bson:"start_place"`
	EndPlace   string    `json:"endPlace" bson:"end_place"`
	Date       time.Time `json:"date" bson:"date"`
}
