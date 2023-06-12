package handler

import (
	"context"
	"fmt"

	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_rating_service/domain"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_rating_service/service"
	pb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/accomodation_rating_service/generated"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AccomodationRatingHandler struct {
	pb.UnimplementedAccomodationRatingServiceServer
	service *service.AccomodationRatingService
}

func NewAccomodationRatingHandler(service *service.AccomodationRatingService) *AccomodationRatingHandler {
	return &AccomodationRatingHandler{
		service: service,
	}
}

func (handler *AccomodationRatingHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	accomodationRating, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	accomodationRatingPb := mapAccomodationRating(accomodationRating)
	response := &pb.GetResponse{
		AccomodationRating: accomodationRatingPb,
	}
	return response, nil
}

func (handler *AccomodationRatingHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	accomodationRatings, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		AccomodationRatings: []*pb.AccomodationRating{},
	}
	for _, accomodationRating := range accomodationRatings {
		current := mapAccomodationRating(accomodationRating)
		response.AccomodationRatings = append(response.AccomodationRatings, current)
	}
	return response, nil
}

func (handler *AccomodationRatingHandler) GetAllByAccomodationId(ctx context.Context, request *pb.GetAllByAccomodationIdRequest) (*pb.GetAllResponse, error) {
	fmt.Println(request.AccomodationId)
	accomodationId := request.AccomodationId
	allAccomodationRatings, err := handler.service.GetAll()
	accomodationRatings := make([]*domain.AccomodationRating, 0)
	for _, res := range allAccomodationRatings {
		if res.AccomodationId == accomodationId {
			accomodationRatings = append(accomodationRatings, res)
		}

	}
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		AccomodationRatings: []*pb.AccomodationRating{},
	}
	for _, accomodationRating := range accomodationRatings {
		current := mapAccomodationRating(accomodationRating)
		response.AccomodationRatings = append(response.AccomodationRatings, current)
	}
	return response, nil
}

func (handler *AccomodationRatingHandler) CreateAccomodationRating(ctx context.Context, request *pb.CreateAccomodationRatingRequest) (*pb.CreateAccomodationRatingResponse, error) {
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
	accomodationRating := mapAccomodationRatingPb(request.AccomodationRating)
	err := handler.service.Insert(*accomodationRating)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccomodationRatingResponse{
		AccomodationRating: mapAccomodationRating(accomodationRating),
	}, nil
}

func (handler *AccomodationRatingHandler) EditAccomodationRating(ctx context.Context, request *pb.EditAccomodationRatingRequest) (*pb.EditAccomodationRatingResponse, error) {
	accomodationRating := mapAccomodationRatingPb(request.AccomodationRating)
	err := handler.service.Edit(*accomodationRating)
	if err != nil {
		return nil, err
	}
	return &pb.EditAccomodationRatingResponse{
		AccomodationRating: mapAccomodationRating(accomodationRating),
	}, nil
}

func (handler *AccomodationRatingHandler) DeleteAccomodationRating(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
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

func (handler *AccomodationRatingHandler) GetByAccomodationAndUser(ctx context.Context, request *pb.GetByAccomodationAndUserRequest) (*pb.GetAllResponse, error) {
	accomodationId := request.AccomodationId
	email := request.Email
	allAccomodationRatings, err := handler.service.GetAll()
	accomodationRatings := make([]*domain.AccomodationRating, 0)
	for _, res := range allAccomodationRatings {
		if res.AccomodationId == accomodationId && res.Email == email {
			accomodationRatings = append(accomodationRatings, res)
		}

	}
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		AccomodationRatings: []*pb.AccomodationRating{},
	}
	for _, accomodationRating := range accomodationRatings {
		current := mapAccomodationRating(accomodationRating)
		response.AccomodationRatings = append(response.AccomodationRatings, current)
	}

	return response, nil
}
