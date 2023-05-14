package handler

import (
	"context"

	pb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/user_service/generated"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/user_service/service"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/user_service/token"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

func (handler *UserHandler) GetByUsername(ctx context.Context, request *pb.GetRequestUsername) (*pb.GetResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Internal, "failed to get metadata from context")
	}
	myHeaderValues := md.Get("Authorization")
	if len(myHeaderValues) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Authorization is required")
	}
	myHeaderValue := myHeaderValues[0]
	username, err := token.ExtractTokenUsername(myHeaderValue)
	if err != nil {
		return nil, err
	}

	user, err := handler.service.GetByUsername(username)
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

func (handler *UserHandler) LoginUser(ctx context.Context, request *pb.LoginDTO) (*pb.UserToken, error) {

	/*user, err := handler.service.GetByUsername(request.LoginDTO.Username)
	if err != nil {
		return nil, err
	}
	generatedToken, err := token.GenerateToken(user.Username, user.Role)*/
	generatedToken, err := handler.service.Login(request.LoginDTO.Username, request.LoginDTO.Password)
	if err != nil {
		return nil, err
	}
	response := &pb.UserToken{
		Token: generatedToken,
	}
	return response, nil
}
