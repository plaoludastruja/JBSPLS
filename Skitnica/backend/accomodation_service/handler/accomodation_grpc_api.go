package handler

import (
	"context"
	"fmt"

	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_service/domain"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_service/service"
	pb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/accomodation_service/generated"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AccomodationHandler struct {
	pb.UnimplementedAccomodationServiceServer
	service *service.AccomodationService
}

func NewAccomodationHandler(service *service.AccomodationService) *AccomodationHandler {
	return &AccomodationHandler{
		service: service,
	}
}

func (handler *AccomodationHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	accomodation, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	accomodationPb := mapAccomodation(accomodation)
	response := &pb.GetResponse{
		Accomodation: accomodationPb,
	}
	return response, nil
}

func (handler *AccomodationHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	accomodations, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Accomodations: []*pb.Accomodation{},
	}
	for _, accomodation := range accomodations {
		current := mapAccomodation(accomodation)
		response.Accomodations = append(response.Accomodations, current)
	}
	return response, nil
}

func (handler *AccomodationHandler) CreateAccomodation(ctx context.Context, request *pb.CreateAccomodationRequest) (*pb.CreateAccomodationResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Internal, "failed to get metadata from context")
	}

	// Get the value of a specific header
	myHeaderValues := md.Get("Authorization")
	fmt.Println(len(myHeaderValues), "-------------------------------------")
	if len(myHeaderValues) == 0 {
		return nil, status.Error(codes.InvalidArgument, "my-header is required")
	}
	fmt.Println(myHeaderValues, "-------------------------------------")
	myHeaderValue := myHeaderValues[0]

	fmt.Println(myHeaderValue, "-------------------------------------")
	accomodation := mapAccomodationPb(request.Accomodation)
	err := handler.service.Insert(*accomodation)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccomodationResponse{
		Accomodation: mapAccomodation(accomodation),
	}, nil
}

func (handler *AccomodationHandler) EditAccomodation(ctx context.Context, request *pb.EditAccomodationRequest) (*pb.EditAccomodationResponse, error) {
	accomodation := mapAccomodationPb(request.Accomodation)
	err := handler.service.Edit(*accomodation)
	if err != nil {
		return nil, err
	}
	return &pb.EditAccomodationResponse{
		Accomodation: mapAccomodation(accomodation),
	}, nil
}

func (handler *AccomodationHandler) DeleteAccomodation(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
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

func (handler *AccomodationHandler) Search(ctx context.Context, request *pb.SearchRequest) (*pb.SearchResponse, error) {
	location := request.Location
	guestNumber := request.GuestNumber
	accomodations, err := handler.service.Search(location, guestNumber)
	if err != nil {
		return nil, err
	}
	response := &pb.SearchResponse{
		Accomodations: []*pb.Accomodation{},
	}
	for _, accomodation := range accomodations {
		current := mapAccomodation(accomodation)
		response.Accomodations = append(response.Accomodations, current)
	}
	return response, nil
}

func (handler *AccomodationHandler) GetByHostUsernameList(ctx context.Context, request *pb.GetByHostUsernameRequest) (*pb.GetByHostUsernameResponse, error) {
	/*accomodations, err := handler.service.GetByHostUsername(request.HostUsername)
	if err != nil {
		return nil, err
	}
	response := &pb.GetByHostUsernameResponse{
		Accomodations: []*pb.Accomodation{},
	}
	for _, accomodation := range accomodations {
		current := mapAccomodation(accomodation)
		response.Accomodations = append(response.Accomodations, current)
	}
	return response, nil*/
	username := request.HostUsername
	allAccomodations, err := handler.service.GetAll()
	accomodations := make([]*domain.Accomodation, 0)
	for _, res := range allAccomodations {
		if res.HostUsername == username {
			accomodations = append(accomodations, res)
		}

	}
	if err != nil {
		return nil, err
	}
	response := &pb.GetByHostUsernameResponse{
		Accomodations: []*pb.Accomodation{},
	}
	for _, accomodation := range accomodations {
		current := mapAccomodation(accomodation)
		response.Accomodations = append(response.Accomodations, current)
	}
	return response, nil
}
