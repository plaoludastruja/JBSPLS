package handler

import (
	pb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/notification_service/generated"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/notification_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapNotification(user *domain.Notification) *pb.Notification {
	userPb := &pb.Notification{
		Id:       user.Id.Hex(),
		Receiver: user.Receiver,
		Sender:   user.Sender,
		Subject:  user.Subject,
		Message:  user.Message,
		IsRead:   user.IsRead,
	}
	return userPb
}

func mapNotificationPb(userPb *pb.Notification) *domain.Notification {
	userPbId, _ := primitive.ObjectIDFromHex(userPb.Id)
	user := &domain.Notification{
		Id:       userPbId,
		Receiver: userPb.Receiver,
		Sender:   userPb.Sender,
		Subject:  userPb.Subject,
		Message:  userPb.Message,
		IsRead:   userPb.IsRead,
	}
	return user
}
