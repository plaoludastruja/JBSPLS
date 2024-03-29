package startup

import (
	"context"
	"fmt"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/appointment_service/handler"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/appointment_service/repository"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/appointment_service/service"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/appointment_service/startup/config"
	appointmentPb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/appointment_service/generated"
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
	appointmentStore := repository.NewAppointmentRepo(mongoClient)
	appointmentService := service.NewAppointmentService(appointmentStore)
	appointmentHandler := handler.NewAppointmentHandler(appointmentService)
	server.startGrpcServer(appointmentHandler)
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

func (server *Server) startGrpcServer(appointmentHandler *handler.AppointmentHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	appointmentPb.RegisterAppointmentServiceServer(grpcServer, appointmentHandler)
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
