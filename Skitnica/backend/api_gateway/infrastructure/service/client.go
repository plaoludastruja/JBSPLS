package service

import (
	"log"

	accomodationRatingGW "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/accomodation_rating_service/generated"
	accomodationGW "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/accomodation_service/generated"
	appointmentGW "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/appointment_service/generated"
	hostmarkGW "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/hostmark_service/generated"
	reservationGW "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/reservation_service/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAccomodationClient(address string) accomodationGW.AccomodationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return accomodationGW.NewAccomodationServiceClient(conn)
}

func NewAppointmentClient(address string) appointmentGW.AppointmentServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return appointmentGW.NewAppointmentServiceClient(conn)
}

func NewReservationClient(address string) reservationGW.ReservationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return reservationGW.NewReservationServiceClient(conn)
}

func NewHostMarkClient(address string) hostmarkGW.HostMarkServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return hostmarkGW.NewHostMarkServiceClient(conn)
}

func NewAccomodationRatingClient(address string) accomodationRatingGW.AccomodationRatingServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return accomodationRatingGW.NewAccomodationRatingServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
