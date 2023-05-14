package handler

import (
	"fmt"
	"time"

	pb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/reservation_service/generated"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/reservation_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapReservation(reservation *domain.Reservation) *pb.Reservation {
	reservationPb := &pb.Reservation{
		Id:             reservation.Id.Hex(),
		AccomodationId: reservation.AccomodationId,
		Username:       reservation.Username,
		StartDate:      reservation.StartDate.String(),
		EndDate:        reservation.EndDate.String(),
		GuestNumber:    reservation.GuestNumber,
		Status:         reservation.Status,
		HostUsername:   reservation.HostUsername,
	}
	return reservationPb
}

func mapReservationPb(reservationPb *pb.Reservation) *domain.Reservation {
	reservationPbId, _ := primitive.ObjectIDFromHex(reservationPb.Id)
	const layout = "2006-01-02"
	fmt.Println("From mapper:")
	fmt.Println(reservationPb.StartDate)
	startDate, _ := time.Parse(layout, reservationPb.StartDate)
	fmt.Println(startDate)
	endDate, _ := time.Parse(layout, reservationPb.EndDate)
	reservation := &domain.Reservation{
		Id:             reservationPbId,
		AccomodationId: reservationPb.AccomodationId,
		Username:       reservationPb.Username,
		StartDate:      startDate,
		EndDate:        endDate,
		GuestNumber:    reservationPb.GuestNumber,
		Status:         reservationPb.Status,
		HostUsername:   reservationPb.HostUsername,
	}
	return reservation
}

func mapDateRange(dateRange *domain.DateRange) *pb.DateRange {
	dateRangePb := &pb.DateRange{
		StartDate:      dateRange.StartDate.String(),
		EndDate:        dateRange.EndDate.String(),
		AccomodationId: dateRange.AccomodationId,
	}
	return dateRangePb
}

func mapDateRangePb(dateRangePb *pb.DateRange) *domain.DateRange {
	const layout = "2006-01-02"
	startDate, _ := time.Parse(layout, dateRangePb.StartDate)
	endDate, _ := time.Parse(layout, dateRangePb.EndDate)
	dateRange := &domain.DateRange{
		StartDate:      startDate,
		EndDate:        endDate,
		AccomodationId: dateRangePb.AccomodationId,
	}
	return dateRange
}

func mapReservationDto(reservationDto *domain.ReservationDto) *pb.ReservationDto {
	reservationDtoPb := &pb.ReservationDto{
		Id:              reservationDto.Id.Hex(),
		AccomodationId:  reservationDto.AccomodationId,
		Username:        reservationDto.Username,
		StartDate:       reservationDto.StartDate.String(),
		EndDate:         reservationDto.EndDate.String(),
		GuestNumber:     reservationDto.GuestNumber,
		CancellationNum: reservationDto.CancellationNum,
	}
	return reservationDtoPb
}

func mapReservationDtoPb(reservationDtoPb *pb.ReservationDto) *domain.ReservationDto {
	reservationDtoPbId, _ := primitive.ObjectIDFromHex(reservationDtoPb.Id)
	const layout = "2006-01-02"
	fmt.Println("From mapper:")
	fmt.Println(reservationDtoPb.StartDate)
	startDate, _ := time.Parse(layout, reservationDtoPb.StartDate)
	fmt.Println(startDate)
	endDate, _ := time.Parse(layout, reservationDtoPb.EndDate)
	reservationDto := &domain.ReservationDto{
		Id:              reservationDtoPbId,
		AccomodationId:  reservationDtoPb.AccomodationId,
		Username:        reservationDtoPb.Username,
		StartDate:       startDate,
		EndDate:         endDate,
		GuestNumber:     reservationDtoPb.GuestNumber,
		CancellationNum: reservationDtoPb.CancellationNum,
	}
	return reservationDto
}
