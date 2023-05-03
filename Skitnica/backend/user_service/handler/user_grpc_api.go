package handler

import (
	"context"

	pb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/user_service/generated"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/user_service/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	user, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	userPb := mapUser(user)
	response := &pb.GetResponse{
		User: userPb,
	}
	return response, nil
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	users, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}
	for _, user := range users {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler *UserHandler) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := mapUserPb(request.User)
	err := handler.service.Insert(*user)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		User: mapUser(user),
	}, nil
}

func (handler *UserHandler) EditUser(ctx context.Context, request *pb.EditUserRequest) (*pb.EditUserResponse, error) {
	user := mapUserPb(request.User)
	err := handler.service.Edit(*user)
	if err != nil {
		return nil, err
	}
	return &pb.EditUserResponse{
		User: mapUser(user),
	}, nil
}

func (handler *UserHandler) DeleteUser(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	errr := handler.service.Delete(objectId)
	if errr != nil {
		return nil, err
	}
	return &pb.DeleteResponse{}, nil
}
