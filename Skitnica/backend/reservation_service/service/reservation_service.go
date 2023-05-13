package service

import (
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/reservation_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationService struct {
	store domain.IReservationRepo
}

func NewReservationService(store domain.IReservationRepo) *ReservationService {
	return &ReservationService{
		store: store,
	}
}

func (service *ReservationService) Get(id primitive.ObjectID) (*domain.Reservation, error) {
	return service.store.Get(id)
}

func (service *ReservationService) GetAll() ([]*domain.Reservation, error) {
	return service.store.GetAll()
}

func (service *ReservationService) Insert(reservation domain.Reservation) error {
	return service.store.Insert(&reservation)
}

func (service *ReservationService) Delete(id primitive.ObjectID) error {
	return service.store.Delete(id)
}

func (service *ReservationService) Edit(reservation domain.Reservation) error {
	return service.store.Edit(&reservation)
}

func (service *ReservationService) Check(dateRange domain.DateRange) ([]*domain.Reservation, error) {
	return service.store.Check(&dateRange)
}
