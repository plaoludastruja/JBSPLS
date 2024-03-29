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
	accomodationRatingGw "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/accomodation_rating_service/generated"
	accomodationGw "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/accomodation_service/generated"
	appointmentGw "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/appointment_service/generated"
	hostmarkGw "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/hostmark_service/generated"
	notificationGw "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/notification_service/generated"
	reservationGw "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/reservation_service/generated"
	userGw "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/user_service/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type Server struct {
	config *config.Config
	mux    *runtime.ServeMux
}

func NewServer(config *config.Config) *Server {
	server := &Server{
		config: config,
		mux: runtime.NewServeMux(
			runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
				header := request.Header.Get("Authorization")
				// send all the headers received from the client
				md := metadata.Pairs("auth", header)
				return md
			}),
		),
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

	hostmarkEndpoint := fmt.Sprintf("%s:%s", server.config.HostMarkHost, server.config.HostMarkPort)
	errHostMark := hostmarkGw.RegisterHostMarkServiceHandlerFromEndpoint(context.TODO(), server.mux, hostmarkEndpoint, opts)
	if errAppointment != nil {
		panic(errHostMark)
	}

	accomodationRatingEndpoint := fmt.Sprintf("%s:%s", server.config.AccomodationRatingHost, server.config.AccomodationRatingPort)
	errAccomodationRating := accomodationRatingGw.RegisterAccomodationRatingServiceHandlerFromEndpoint(context.TODO(), server.mux, accomodationRatingEndpoint, opts)
	if errAccomodationRating != nil {
		panic(errAccomodationRating)
	}

	notificationEndpoint := fmt.Sprintf("%s:%s", server.config.NotificationkHost, server.config.NotificationPort)
	errNotificationnRating := notificationGw.RegisterNotificationServiceHandlerFromEndpoint(context.TODO(), server.mux, notificationEndpoint, opts)
	if errNotificationnRating != nil {
		panic(errNotificationnRating)
	}

}

func (server *Server) initCustomHandlers() {
	//userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	accomodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccomodationHost, server.config.AccomodationPort)
	appointmentEndpoint := fmt.Sprintf("%s:%s", server.config.AppointmentHost, server.config.AppointmentPort)
	reservationEndpoint := fmt.Sprintf("%s:%s", server.config.ReservationHost, server.config.ReservationPort)
	hostmarkEndpoint := fmt.Sprintf("%s:%s", server.config.HostMarkHost, server.config.HostMarkPort)

	searchHandler := handler.NewSearchHandler(accomodationEndpoint, appointmentEndpoint, reservationEndpoint)
	searchHandler.Init(server.mux)

	bestHostHandler := handler.NewBestHosthHandler(reservationEndpoint, hostmarkEndpoint)
	bestHostHandler.Init(server.mux)

	filterHandler := handler.NewFilterHandler(accomodationEndpoint, appointmentEndpoint, reservationEndpoint, hostmarkEndpoint)
	filterHandler.Init(server.mux)
}

func (server *Server) Start() {
	corsHandler := cors.CORSMiddleware(server.mux)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), corsHandler))
}
