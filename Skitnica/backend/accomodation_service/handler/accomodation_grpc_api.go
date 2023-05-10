package handler

import (
	"context"

	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_service/service"
	pb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/accomodation_service/generated"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
