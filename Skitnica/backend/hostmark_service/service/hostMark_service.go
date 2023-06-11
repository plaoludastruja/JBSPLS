package service

import (
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/hostmark_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HostMarkService struct {
	store domain.IHostMarkRepo
}

func NewHostMarkService(store domain.IHostMarkRepo) *HostMarkService {
	return &HostMarkService{
		store: store,
	}
}

func (service *HostMarkService) Get(id primitive.ObjectID) (*domain.HostMark, error) {
	return service.store.Get(id)
}

func (service *HostMarkService) GetAll() ([]*domain.HostMark, error) {
	return service.store.GetAll()
}

func (service *HostMarkService) Insert(hostMark domain.HostMark) error {
	return service.store.Insert(&hostMark)
}

func (service *HostMarkService) Delete(id primitive.ObjectID) error {
	return service.store.Delete(id)
}

func (service *HostMarkService) Edit(hostMark domain.HostMark) error {
	return service.store.Edit(&hostMark)
}

func (service *HostMarkService) GetByHostAndUser(username string, hostUsername string) ([]*domain.HostMark, error) {
	return service.store.GetByUsername(username, hostUsername)
}

func (service *HostMarkService) IsBestHostCheck(hostUsername string) bool {
	marks, _ := service.store.GetAllByHostUsername(hostUsername)
	numOfMarks := len(marks)
	if numOfMarks == 0 {
		numOfMarks = 1
	}

	marksSum := 0
	for _, mark := range marks {
		marksSum += int(mark.Grade)
	}

	marksAverage := float64(marksSum) / float64(numOfMarks)
	if marksAverage < 4 {
		return false
	}

	return true
}
