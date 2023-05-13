package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type SearchResult struct {
	Id                  primitive.ObjectID
	AccomodationId      string
	Name                string
	Location            string
	Facilities          string
	MinNumberOfGuests   int32
	MaxNumberOfGuests   int32
	TotalPrice          int32
	IsAutomaticApproved string
}
