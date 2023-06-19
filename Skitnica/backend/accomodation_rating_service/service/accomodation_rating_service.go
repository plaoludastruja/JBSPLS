package service

import (
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_rating_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccomodationRatingService struct {
	store domain.IAccomodationRatingRepo
}

func NewAccomodationRatingService(store domain.IAccomodationRatingRepo) *AccomodationRatingService {
	return &AccomodationRatingService{
		store: store,
	}
}

func (service *AccomodationRatingService) Get(id primitive.ObjectID) (*domain.AccomodationRating, error) {
	return service.store.Get(id)
}

func (service *AccomodationRatingService) GetAll() ([]*domain.AccomodationRating, error) {
	return service.store.GetAll()
}

func (service *AccomodationRatingService) Insert(accomodation domain.AccomodationRating) error {
	return service.store.Insert(&accomodation)
}

func (service *AccomodationRatingService) Delete(id primitive.ObjectID) error {
	return service.store.Delete(id)
}

func (service *AccomodationRatingService) Edit(accomodation domain.AccomodationRating) error {
	return service.store.Edit(&accomodation)
}

func (service *AccomodationRatingService) GetRecommended(email string) ([]string, error) {
	return service.store.GetRecommended(email)
}
