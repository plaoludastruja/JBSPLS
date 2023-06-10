package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAccomodationRatingRepo interface {
	Get(id primitive.ObjectID) (*AccomodationRating, error)
	GetAll() ([]*AccomodationRating, error)
	Insert(accomodation *AccomodationRating) error
	Edit(accomodation *AccomodationRating) error
	DeleteAll()
	Delete(id primitive.ObjectID) error
}
