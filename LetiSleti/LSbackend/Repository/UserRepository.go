package Repository

import (
	"context"
	"log"

	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user Models.User) bool {
	if _, err := usersCollection.InsertOne(context.TODO(), user); err != nil {
		log.Panic("Could not save document to database", err.Error())
	}
	return true
}

func GetAllUsers() []Models.User {
	results := []Models.User{}
	cursor, _ := usersCollection.Find(context.TODO(), bson.D{})
	cursor.All(context.TODO(), &results)
	return results
}

func GetUserByEmail(email string) (Models.User, error) {
	var result Models.User
	err := usersCollection.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, err
		}
		log.Fatal(err)
	}
	return result, err
}
