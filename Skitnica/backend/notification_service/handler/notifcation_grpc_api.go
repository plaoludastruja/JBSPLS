package handler

import (
	"context"

	pb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/notification_service/generated"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/notification_service/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationHandler struct {
	pb.UnimplementedNotificationServiceServer
	service *service.NotificationService
}

func NewNotificationHandler(service *service.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		service: service,
	}
}

func (handler *NotificationHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	notification, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	notificationPb := mapNotification(notification)
	response := &pb.GetResponse{
		Notification: notificationPb,
	}
	return response, nil
}

func (handler *NotificationHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	notifications, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Notifications: []*pb.Notification{},
	}
	for _, notification := range notifications {
		current := mapNotification(notification)
		response.Notifications = append(response.Notifications, current)
	}
	return response, nil
}

func (handler *NotificationHandler) GetByReceiver(ctx context.Context, request *pb.GetRequestReceiver) (*pb.GetAllResponse, error) {
	notifications, err := handler.service.GetByReceiver(request.Receiver)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Notifications: []*pb.Notification{},
	}
	for _, notification := range notifications {
		current := mapNotification(notification)
		response.Notifications = append(response.Notifications, current)
	}
	return response, nil
}

func (handler *NotificationHandler) GetBySender(ctx context.Context, request *pb.GetRequestSender) (*pb.GetAllResponse, error) {
	notifications, err := handler.service.GetBySender(request.Sender)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Notifications: []*pb.Notification{},
	}
	for _, notification := range notifications {
		current := mapNotification(notification)
		response.Notifications = append(response.Notifications, current)
	}
	return response, nil
}

func (handler *NotificationHandler) CreateNotification(ctx context.Context, request *pb.CreateNotificationRequest) (*pb.CreateNotificationResponse, error) {
	notification := mapNotificationPb(request.Notification)
	err := handler.service.Insert(*notification)
	if err != nil {
		return nil, err
	}
	return &pb.CreateNotificationResponse{
		Notification: mapNotification(notification),
	}, nil
}

func (handler *NotificationHandler) ReadAllByUsername(ctx context.Context, request *pb.ReadAllByUsernameRequest) (*pb.ReadAllByUsernameResponse, error) {

	err := handler.service.ReadAllByUsername(request.Username)
	if err != nil {
		return nil, err
	}
	return &pb.ReadAllByUsernameResponse{}, nil
}

func (handler *NotificationHandler) GetNotificationFilterByUsername(ctx context.Context, request *pb.ReadAllByUsernameRequest) (*pb.NotificationFilterResponse, error) {
	notificationFilter, err := handler.service.GetNotificationFilterByUsername(request.Username)
	if err != nil {
		return nil, err
	}
	notificationFilterPb := mapNotificationFilter(notificationFilter)
	response := &pb.NotificationFilterResponse{
		NotificationFilter: notificationFilterPb,
	}
	return response, nil
}

func (handler *NotificationHandler) CreateNotificationFilter(ctx context.Context, request *pb.NotificationFilterRequest) (*pb.NotificationFilterResponse, error) {
	notification := mapNotificationFilterPb(request.NotificationFilter)
	err := handler.service.InsertNotificationFilters(*notification)
	if err != nil {
		return nil, err
	}
	return &pb.NotificationFilterResponse{
		NotificationFilter: mapNotificationFilter(notification),
	}, nil
}

func (handler *NotificationHandler) EditNotificationFilter(ctx context.Context, request *pb.NotificationFilterRequest) (*pb.NotificationFilterResponse, error) {
	notification := mapNotificationFilterPb(request.NotificationFilter)
	err := handler.service.EditNotificationFilter(*notification)
	if err != nil {
		return nil, err
	}
	return &pb.NotificationFilterResponse{
		NotificationFilter: mapNotificationFilter(notification),
	}, nil
}
