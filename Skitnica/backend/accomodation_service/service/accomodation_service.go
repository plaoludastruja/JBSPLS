package service

import (
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccomodationService struct {
	store domain.IAccomodationRepo
}

func NewAccomodationService(store domain.IAccomodationRepo) *AccomodationService {
	return &AccomodationService{
		store: store,
	}
}

func (service *AccomodationService) Get(id primitive.ObjectID) (*domain.Accomodation, error) {
	return service.store.Get(id)
}

func (service *AccomodationService) GetAll() ([]*domain.Accomodation, error) {
	return service.store.GetAll()
}

func (service *AccomodationService) Insert(accomodation domain.Accomodation) error {
	return service.store.Insert(&accomodation)
}

func (service *AccomodationService) Delete(id primitive.ObjectID) error {
	return service.store.Delete(id)
}

func (service *AccomodationService) Edit(accomodation domain.Accomodation) error {
	return service.store.Edit(&accomodation)
}

func (service *AccomodationService) Search(location string, guestNumber int32) ([]*domain.Accomodation, error) {
	return service.store.Search(location, guestNumber)
}
