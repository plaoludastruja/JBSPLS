package service

import (
	"fmt"

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

func (service *ReservationService) Search(startDay string, startMonth string, startYear string, endDay string, endMonth string, endYear string) ([]*domain.Reservation, error) {
	return service.store.Search(startDay, startMonth, startYear, endDay, endMonth, endYear)
}
func (service *ReservationService) Check(dateRange domain.DateRange) ([]*domain.Reservation, error) {
	return service.store.Check(&dateRange)
}

func (service *ReservationService) GetAllPending(hostUsername string) ([]*domain.Reservation, error) {
	return service.store.GetAllPending(hostUsername)
}

func (service *ReservationService) GetCanceledForUser(username string) (int32, error) {
	list, _ := service.store.GetCanceledForUser(username)
	num := 0
	for _, reservation := range list {
		fmt.Println(reservation.Username)
		num = num + 1
	}
	return int32(num), nil
}

func (service *ReservationService) ApproveReservation(reservationDto domain.ReservationDto) {
	service.store.ApproveReservation(&reservationDto)
}

func (service *ReservationService) RejectReservation(reservationDto domain.ReservationDto) {
	service.store.RejectReservation(&reservationDto)
}
