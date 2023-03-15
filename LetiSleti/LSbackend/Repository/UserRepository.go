package Repository

import (
	"context"
	"fmt"
	"log"

	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"go.mongodb.org/mongo-driver/bson"
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
	fmt.Println(results)
	return results
}
