package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAppointmentRepo interface {
	Get(id primitive.ObjectID) (*Appointment, error)
	GetAll() ([]*Appointment, error)
	Insert(appointment *Appointment) error
	Edit(appointment *Appointment) error
	DeleteAll()
	Delete(id primitive.ObjectID) error
}
