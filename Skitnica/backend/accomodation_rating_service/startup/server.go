package startup

import (
	"context"
	"fmt"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_rating_service/handler"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_rating_service/repository"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_rating_service/service"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_rating_service/startup/config"
	accomodationRatingPb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/accomodation_rating_service/generated"
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
	accomodationRatingStore := repository.NewAccomodationRatingRepo(mongoClient)
	accomodationRatingService := service.NewAccomodationRatingService(accomodationRatingStore)
	accomodationRatingHandler := handler.NewAccomodationRatingHandler(accomodationRatingService)
	server.startGrpcServer(accomodationRatingHandler)
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

func (server *Server) startGrpcServer(accomodationRatingHandler *handler.AccomodationRatingHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	accomodationRatingPb.RegisterAccomodationRatingServiceServer(grpcServer, accomodationRatingHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
