package service

import "github.com/plaoludastruja/JBSPLS/Skitnica/backend/notification_service/domain"

type NotificationService struct {
	store domain.INotificationRepo
}

func NewNotificationService(store domain.INotificationRepo) *NotificationService {
	return &NotificationService{
		store: store,
	}
}
