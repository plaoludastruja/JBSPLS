package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type IReservationRepo interface {
	Get(id primitive.ObjectID) (*Reservation, error)
	GetAll() ([]*Reservation, error)
	Insert(reservation *Reservation) error
	Edit(reservation *Reservation) error
	DeleteAll()
	Delete(id primitive.ObjectID) error
	Search(startDay string, startMonth string, startYear string, endDay string, endMonth string, endYear string) ([]*Reservation, error)
	Check(dateRange *DateRange) ([]*Reservation, error)
	GetAllPending(hostUsername string) ([]*Reservation, error)
	ApproveReservation(reservationDto *ReservationDto) error
	RejectReservation(reservationDto *ReservationDto) error
	GetCanceledForUser(username string) ([]*Reservation, error)
	RejectOverlapsed(reservationDto *ReservationDto) error
}
