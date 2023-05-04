package startup

import (
	"context"
	"fmt"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	userPb "github.com/plaoludastruja/JBSPLS/Skitnica/backend/common/proto/user_service/generated"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/user_service/handler"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/user_service/repository"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/user_service/service"
	"github.com/plaoludastruja/JBSPLS/Skitnica/backend/user_service/startup/config"
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
	userStore := repository.NewUserRepo(mongoClient)
	userService := service.NewUserService(userStore)
	userHandler := handler.NewUserHandler(userService)
	server.startGrpcServer(userHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	uri := fmt.Sprintf("mongodb://%s:%s/", server.config.UserDBHost, server.config.UserDBPort)
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func (server *Server) startGrpcServer(userHandler *handler.UserHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	userPb.RegisterUserServiceServer(grpcServer, userHandler)
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
