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
	notificationFilter, _ := service.GetNotificationFilterByUsername(username)
	notifications, err := service.store.GetByReceiver(username)
	if err != nil {
		return notifications, err
	}
	var mySlice []*domain.Notification
	for _, notification := range notifications {

		if (notification.Subject == "reservation" && notificationFilter.Reservation) || (notification.Subject == "rating" && notificationFilter.Rating) || (notification.Subject == "super" && notificationFilter.Super) {
			mySlice = append(mySlice, notification)
		}

	}

	return mySlice, err
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

func (service *NotificationService) ReadAllByUsername(username string) error {
	notifications, _ := service.GetByReceiver(username)
	for _, notification := range notifications {
		err := service.store.Edit(notification)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *NotificationService) InsertNotificationFilters(notificationFilter domain.NotificationFilter) error {
	return service.store.InsertNotificationFilters(&notificationFilter)
}

func (service *NotificationService) EditNotificationFilter(notificationFilter domain.NotificationFilter) error {
	return service.store.EditNotificationFilter(&notificationFilter)
}

func (service *NotificationService) GetNotificationFilterByUsername(username string) (*domain.NotificationFilter, error) {
	return service.store.GetNotificationFilterByUsername(username)
}
