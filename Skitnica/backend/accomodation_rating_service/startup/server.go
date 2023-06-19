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

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
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
	neo4jSession := server.initNeo()
	accomodationRatingStore := repository.NewAccomodationRatingRepo(mongoClient, neo4jSession)
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

func (server *Server) initNeo() neo4j.Session {
	uri := "neo4j+s://59ab3a7b.databases.neo4j.io"
	auth := neo4j.BasicAuth("neo4j", "SyxR9cQpsGOfa2u5-Ol-Ygrw04UC3pQ-X9Js93EKqeI", "")
	driver, err := neo4j.NewDriver(uri, auth)
	if err != nil {
		panic(err)
	}
	session := driver.NewSession(neo4j.SessionConfig{DatabaseName: "neo4j"})
	return session
}

/*func (server *Server) initNeo() neo4j.Session {
	uri := "neo4j+s://8fe831a7.databases.neo4j.io"                                      // Promenite URI na NEO4J_URI vrednost
	auth := neo4j.BasicAuth("neo4j", "S_wvIcfiFAiCWu1Zqa9Bp8ShNW8OAqMQ7jvMF4m4HjE", "") // Promenite korisničko ime i lozinku na NEO4J_USERNAME i NEO4J_PASSWORD vrednosti
	driver, err := neo4j.NewDriver(uri, auth)
	if err != nil {
		panic(err)
	}
	session := driver.NewSession(neo4j.SessionConfig{DatabaseName: "neo4j"})
	return session
}*/
