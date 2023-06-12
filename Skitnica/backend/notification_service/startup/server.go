package startup

import (
	"context"
	"fmt"
	"log"
	"net"

	notificationPb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/notification_service/generated"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/notification_service/handler"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/notification_service/repository"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/notification_service/service"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/notification_service/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
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
	notificationStore := repository.NewNotificationRepo(mongoClient)
	notificationService := service.NewNotificationService(notificationStore)
	notificationHandler := handler.NewNotificationHandler(notificationService)
	server.startGrpcServer(notificationHandler)
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

func (server *Server) startGrpcServer(notificationHandler *handler.NotificationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	notificationPb.RegisterNotificationServiceServer(grpcServer, notificationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
