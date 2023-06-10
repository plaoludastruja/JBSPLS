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
func (handler *ReservationHandler) GetAllRes(ctx context.Context, request *pb.GetAllReq) (*pb.GetAllResponse, error) {
	accomodations, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Reservations: []*pb.Reservation{},
	}
	for _, reservation := range accomodations {
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

func (handler *ReservationHandler) GetAllPending(ctx context.Context, request *pb.GetAllPendingRequest) (*pb.GetAllPendingResponse, error) {
	fmt.Println(request.HostUsername)
	hostUsername := request.HostUsername
	allPendingReservations, err := handler.service.GetAllPending(hostUsername)

	if err != nil {
		return nil, err
	}
	response := &pb.GetAllPendingResponse{
		ReservationDtos: []*pb.ReservationDto{},
	}
	cancelationNum := int32(0)
	for _, reservation := range allPendingReservations {
		cancelationNum, _ = handler.service.GetCanceledForUser(reservation.Username)
		reservationDto := &domain.ReservationDto{
			Id:              reservation.Id,
			AccomodationId:  reservation.AccomodationId,
			Username:        reservation.Username,
			StartDate:       reservation.StartDate,
			EndDate:         reservation.EndDate,
			GuestNumber:     reservation.GuestNumber,
			CancellationNum: cancelationNum,
		}
		current := mapReservationDto(reservationDto)
		response.ReservationDtos = append(response.ReservationDtos, current)
	}
	return response, nil
}

func (handler *ReservationHandler) ApproveReservation(ctx context.Context, request *pb.ApproveReservationRequest) (*pb.ApproveReservationResponse, error) {
	reservation := mapReservationDtoPb(request.ReservationDto)

	handler.service.ApproveReservation(*reservation)

	return &pb.ApproveReservationResponse{
		ReservationDto: mapReservationDto(reservation),
	}, nil
}

func (handler *ReservationHandler) RejectReservation(ctx context.Context, request *pb.RejectReservationRequest) (*pb.RejectReservationResponse, error) {
	reservation := mapReservationDtoPb(request.ReservationDto)

	handler.service.RejectReservation(*reservation)

	return &pb.RejectReservationResponse{
		ReservationDto: mapReservationDto(reservation),
	}, nil
}

func (handler *ReservationHandler) GetForGuest(ctx context.Context, request *pb.GetGuestRequest) (*pb.GetGuestResponse, error) {
	usernamesR := handler.service.GetForGuest(request.Username)

	response := &pb.GetGuestResponse{
		Usernames: usernamesR,
	}
	return response, nil

}

func (handler *ReservationHandler) IsHostBestHost(ctx context.Context, request *pb.IsHostBestHostRequest) (*pb.IsHostBestHostResposne, error) {
	isBestHost := handler.service.IsBestHostCheck(request.HostUsername)
	response := &pb.IsHostBestHostResposne{
		IsBestHost: isBestHost,
	}
	return response, nil
}
