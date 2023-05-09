package service

import (
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/appointment_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AppointmentService struct {
	store domain.IAppointmentRepo
}

func NewAppointmentService(store domain.IAppointmentRepo) *AppointmentService {
	return &AppointmentService{
		store: store,
	}
}

func (service *AppointmentService) Get(id primitive.ObjectID) (*domain.Appointment, error) {
	return service.store.Get(id)
}

func (service *AppointmentService) GetAll() ([]*domain.Appointment, error) {
	return service.store.GetAll()
}

func (service *AppointmentService) Insert(appointment domain.Appointment) error {
	return service.store.Insert(&appointment)
}

func (service *AppointmentService) Delete(id primitive.ObjectID) error {
	return service.store.Delete(id)
}

func (service *AppointmentService) Edit(appointment domain.Appointment) error {
	return service.store.Edit(&appointment)
}
