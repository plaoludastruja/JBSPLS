package startup

import (
	"context"
	"fmt"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_service/handler"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_service/repository"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_service/service"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/accomodation_service/startup/config"
	accomodationPb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/accomodation_service/generated"
	saga "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/saga/messaging"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/saga/messaging/nats"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "accomodation_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	accomodationStore := repository.NewAccomodationRepo(mongoClient)
	accomodationService := service.NewAccomodationService(accomodationStore)

	commandSubscriber := server.initSubscriber(server.config.DeleteUserCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.DeleteUserReplySubject)
	server.initDeleteUserHandler(accomodationService, replyPublisher, commandSubscriber)

	accomodationHandler := handler.NewAccomodationHandler(accomodationService)
	server.startGrpcServer(accomodationHandler)
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

func (server *Server) startGrpcServer(accomodationHandler *handler.AccomodationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	accomodationPb.RegisterAccomodationServiceServer(grpcServer, accomodationHandler)
	// router := mux.NewRouter().StrictSlash(true)

	// router.Methods("OPTIONS").HandlerFunc(
	// 	func(w http.ResponseWriter, r *http.Request) {
	// 		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	// 		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	// 		w.Header().Set("Access-Control-Allow-Credentials", "true")
	// 		w.WriteHeader(http.StatusNoContent)
	// 	})
	// fmt.Println("server running ")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initDeleteUserHandler(service *service.AccomodationService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := handler.NewDeleteUserCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}
