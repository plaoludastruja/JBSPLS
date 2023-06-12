package service

import (
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/notification_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationService struct {
	store domain.INotificationRepo
}

func NewNotificationService(store domain.INotificationRepo) *NotificationService {
	return &NotificationService{
		store: store,
	}
}

func (service *NotificationService) Get(id primitive.ObjectID) (*domain.Notification, error) {
	return service.store.Get(id)
}

func (service *NotificationService) GetByReceiver(username string) ([]*domain.Notification, error) {
	return service.store.GetByReceiver(username)
}

func (service *NotificationService) GetBySender(username string) ([]*domain.Notification, error) {
	return service.store.GetBySender(username)
}

func (service *NotificationService) GetAll() ([]*domain.Notification, error) {
	return service.store.GetAll()
}

func (service *NotificationService) Insert(user domain.Notification) error {
	return service.store.Insert(&user)
}
