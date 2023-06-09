package handler

import (
	"context"
	"fmt"

	pb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/hostmark_service/generated"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/hostmark_service/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HostMarkHandler struct {
	pb.UnimplementedHostMarkServiceServer
	service *service.HostMarkService
}

func NewHostMarkHandler(service *service.HostMarkService) *HostMarkHandler {
	return &HostMarkHandler{
		service: service,
	}
}

func (handler *HostMarkHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	hostMark, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	hostMarkPb := mapHostMark(hostMark)
	response := &pb.GetResponse{
		Hostmark: hostMarkPb,
	}
	return response, nil
}

func (handler *HostMarkHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	hostMarks, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		HostMark: []*pb.HostMark{},
	}
	for _, hostMark := range hostMarks {
		current := mapHostMark(hostMark)
		response.HostMark = append(response.HostMark, current)
	}
	return response, nil
}

func (handler *HostMarkHandler) CreateHostMark(ctx context.Context, request *pb.CreateHostMarkRequest) (*pb.CreateHostMarkResponse, error) {
	hostMark := mapHostMarkPb(request.HostMark)
	fmt.Println(hostMark)
	err := handler.service.Insert(*hostMark)
	if err != nil {
		return nil, err
	}
	return &pb.CreateHostMarkResponse{
		HostMark: mapHostMark(hostMark),
	}, nil
}

func (handler *HostMarkHandler) EditHostMark(ctx context.Context, request *pb.EditHostMarkRequest) (*pb.EditHostMarkResponse, error) {
	hostMark := mapHostMarkPb(request.HostMark)
	fmt.Println(request.HostMark)
	err := handler.service.Edit(*hostMark)
	if err != nil {
		return nil, err
	}
	return &pb.EditHostMarkResponse{
		HostMark: mapHostMark(hostMark),
	}, nil
}

func (handler *HostMarkHandler) DeleteHostMark(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
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

func (handler *HostMarkHandler) GetByUsername(ctx context.Context, request *pb.GetByUsernameAndHostRequest) (*pb.GetResponse, error) {
	username := request.Username
	hostUsername := request.HostUsername
	hostMark, err := handler.service.GetByUsername(username, hostUsername)
	if err != nil {
		return nil, err
	}
	hostMarkPb := mapHostMark(hostMark)
	response := &pb.GetResponse{
		Hostmark: hostMarkPb,
	}
	return response, nil
}
