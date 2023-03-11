package Repository

import (
	"context"
	"log"

	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
)

func CreateUser(user Models.User) bool {
	if _, err := usersCollection.InsertOne(context.TODO(), user); err != nil {
		log.Panic("Could not save document to database", err.Error())
	}
	return true
}
