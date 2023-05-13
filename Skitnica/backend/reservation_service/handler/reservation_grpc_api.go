package handler

import (
	"context"
	"fmt"

	pb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/reservation_service/generated"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/reservation_service/domain"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/reservation_service/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationHandler struct {
	pb.UnimplementedReservationServiceServer
	service *service.ReservationService
}

func NewReservationHandler(service *service.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		service: service,
	}
}

func (handler *ReservationHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	reservation, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	reservationPb := mapReservation(reservation)
	response := &pb.GetResponse{
		Reservation: reservationPb,
	}
	return response, nil
}

func (handler *ReservationHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	fmt.Println(request.Username)
	username := request.Username
	allReservations, err := handler.service.GetAll()
	reservations := make([]*domain.Reservation, 0)
	for _, res := range allReservations {
		if res.Username == username {
			reservations = append(reservations, res)
		}

	}
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Reservations: []*pb.Reservation{},
	}
	for _, reservation := range reservations {
		current := mapReservation(reservation)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}

func (handler *ReservationHandler) CreateReservation(ctx context.Context, request *pb.CreateReservationRequest) (*pb.CreateReservationResponse, error) {
	reservation := mapReservationPb(request.Reservation)
	fmt.Println(reservation)
	err := handler.service.Insert(*reservation)
	if err != nil {
		return nil, err
	}
	return &pb.CreateReservationResponse{
		Reservation: mapReservation(reservation),
	}, nil
}

func (handler *ReservationHandler) EditReservation(ctx context.Context, request *pb.EditReservationRequest) (*pb.EditReservationResponse, error) {
	reservation := mapReservationPb(request.Reservation)
	fmt.Println(request.Reservation)
	err := handler.service.Edit(*reservation)
	if err != nil {
		return nil, err
	}
	return &pb.EditReservationResponse{
		Reservation: mapReservation(reservation),
	}, nil
}

func (handler *ReservationHandler) DeleteReservation(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
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

func (handler *ReservationHandler) Search(ctx context.Context, request *pb.SearchRequest) (*pb.SearchResponse, error) {
	startDay := request.StartDay
	startMonth := request.StartMonth
	startYear := request.StartYear
	endDay := request.EndDay
	endMonth := request.EndMonth
	endYear := request.EndYear
	reservations, err := handler.service.Search(startDay, startMonth, startYear, endDay, endMonth, endYear)
	if err != nil {
		return nil, err
	}
	response := &pb.SearchResponse{
		Reservations: []*pb.Reservation{},
	}
	for _, reservation := range reservations {
		current := mapReservation(reservation)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}

func (handler *ReservationHandler) Check(ctx context.Context, request *pb.DateRangeRequest) (*pb.DateRangeResponse, error) {
	dateRange := mapDateRangePb(request.DateRange)
	reservations, err := handler.service.Check(*dateRange)
	if err != nil {
		return nil, err
	}
	response := &pb.DateRangeResponse{
		Reservations: []*pb.Reservation{},
	}
	for _, reservation := range reservations {
		current := mapReservation(reservation)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}
