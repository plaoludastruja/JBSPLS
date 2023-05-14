package startup

import (
	"context"
	"fmt"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	reservationPb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/reservation_service/generated"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/reservation_service/handler"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/reservation_service/repository"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/reservation_service/service"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/reservation_service/startup/config"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	reservationStore := repository.NewReservationRepo(mongoClient)
	reservationService := service.NewReservationService(reservationStore)
	reservationHandler := handler.NewReservationHandler(reservationService)
	server.startGrpcServer(reservationHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	uri := fmt.Sprintf("mongodb://%s:%s/", server.config.DBHost, server.config.DBPort)
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func (server *Server) startGrpcServer(reservationHandler *handler.ReservationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reservationPb.RegisterReservationServiceServer(grpcServer, reservationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
