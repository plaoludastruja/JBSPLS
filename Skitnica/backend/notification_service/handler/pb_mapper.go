package handler

import (
	pb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/notification_service/generated"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/notification_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapNotification(notification *domain.Notification) *pb.Notification {
	notificationPb := &pb.Notification{
		Id:       notification.Id.Hex(),
		Receiver: notification.Receiver,
		Sender:   notification.Sender,
		Subject:  notification.Subject,
		Message:  notification.Message,
		IsRead:   notification.IsRead,
	}
	return notificationPb
}

func mapNotificationPb(notificationPb *pb.Notification) *domain.Notification {
	notificationPbId, _ := primitive.ObjectIDFromHex(notificationPb.Id)
	notification := &domain.Notification{
		Id:       notificationPbId,
		Receiver: notificationPb.Receiver,
		Sender:   notificationPb.Sender,
		Subject:  notificationPb.Subject,
		Message:  notificationPb.Message,
		IsRead:   notificationPb.IsRead,
	}
	return notification
}

func mapNotificationFilter(notification *domain.NotificationFilter) *pb.NotificationFilter {
	notificationPb := &pb.NotificationFilter{
		Id:          notification.Id.Hex(),
		Username:    notification.Username,
		Reservation: notification.Reservation,
		Rating:      notification.Rating,
		Super:       notification.Super,
	}
	return notificationPb
}

func mapNotificationFilterPb(notificationPb *pb.NotificationFilter) *domain.NotificationFilter {
	notificationPbId, _ := primitive.ObjectIDFromHex(notificationPb.Id)
	notification := &domain.NotificationFilter{
		Id:          notificationPbId,
		Username:    notificationPb.Username,
		Reservation: notificationPb.Reservation,
		Rating:      notificationPb.Rating,
		Super:       notificationPb.Super,
	}
	return notification
}
