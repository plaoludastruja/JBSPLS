package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAccomodationRepo interface {
	Get(id primitive.ObjectID) (*Accomodation, error)
	GetAll() ([]*Accomodation, error)
	Insert(accomodation *Accomodation) error
	Edit(accomodation *Accomodation) error
	DeleteAll()
	Delete(id primitive.ObjectID) error
	Search(location string, guestNumber int32) ([]*Accomodation, error)
	GetByHostUsernameList(hostUsername string) ([]*Accomodation, error)
}
