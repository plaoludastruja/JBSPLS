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
