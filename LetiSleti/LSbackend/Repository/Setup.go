package Repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DATABASE_URI = "mongodb://db_host:27017"

var usersCollection *mongo.Collection
var flightsCollection *mongo.Collection
var ticketsCollection *mongo.Collection

var client *mongo.Client

func Setup() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DATABASE_URI))

	if err != nil {
		log.Panic(err)
	}

	usersCollection = client.Database("LetiSleti").Collection("users")
	flightsCollection = client.Database("LetiSleti").Collection("flights")
	ticketsCollection = client.Database("LetiSleti").Collection("tickets")
}

func Disconnect() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Panic(err)
	}
}
