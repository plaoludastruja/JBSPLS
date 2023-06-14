package service

import (
	"fmt"

	nats "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/saga/messaging/nats"
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

func (service *NotificationService) Insert(notification domain.Notification) error {
	sub, _ := nats.NewNATSPublisher("nats", "4222", "ruser", "T0pS3cr3t", notification.Receiver)
	fmt.Println("natsPublish : ", notification.Message)
	sub.Publish(notification.Message)
	return service.store.Insert(&notification)
}
