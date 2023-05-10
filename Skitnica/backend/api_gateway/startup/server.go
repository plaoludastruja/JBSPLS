package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	cors "github.com/plaoludastruja/JBSPLS/Skitnica/backend/api_gateway/Helper/Cors"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/api_gateway/infrastructure/handler"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/api_gateway/startup/config"
	accomodationGw "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/accomodation_service/generated"
	appointmentGw "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/appointment_service/generated"
	reservationGw "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/reservation_service/generated"
	userGw "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/user_service/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config *config.Config
	mux    *runtime.ServeMux
}

func NewServer(config *config.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	server.initCustomHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	errUser := userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if errUser != nil {
		panic(errUser)
	}

	accomodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccomodationHost, server.config.AccomodationPort)
	errAccomodation := accomodationGw.RegisterAccomodationServiceHandlerFromEndpoint(context.TODO(), server.mux, accomodationEndpoint, opts)
	if errAccomodation != nil {
		panic(errAccomodation)
	}

	reservationEndpoint := fmt.Sprintf("%s:%s", server.config.ReservationHost, server.config.ReservationPort)
	errReservation := reservationGw.RegisterReservationServiceHandlerFromEndpoint(context.TODO(), server.mux, reservationEndpoint, opts)
	if errReservation != nil {
		panic(errAccomodation)
	}

	appointmentEndpoint := fmt.Sprintf("%s:%s", server.config.AppointmentHost, server.config.AppointmentPort)
	errAppointment := appointmentGw.RegisterAppointmentServiceHandlerFromEndpoint(context.TODO(), server.mux, appointmentEndpoint, opts)
	if errAppointment != nil {
		panic(errAppointment)
	}

}

func (server *Server) initCustomHandlers() {
	accomodationEmdpoint := fmt.Sprintf("%s:%s", server.config.AccomodationHost, server.config.AccomodationPort)
	appointmentEmdpoint := fmt.Sprintf("%s:%s", server.config.AppointmentHost, server.config.AppointmentPort)
	reservationEmdpoint := fmt.Sprintf("%s:%s", server.config.ReservationHost, server.config.ReservationPort)
	searchHandler := handler.NewSearchHandler(accomodationEmdpoint, appointmentEmdpoint, reservationEmdpoint)
	searchHandler.Init(server.mux)
}

func (server *Server) Start() {
	corsHandler := cors.CORSMiddleware(server.mux)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), corsHandler))
}
