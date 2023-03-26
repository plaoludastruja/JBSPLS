package Repository

import (
	"context"
	"log"

	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateTicket(ticket Models.Ticket) bool {
	if _, err := ticketsCollection.InsertOne(context.TODO(), ticket); err != nil {
		log.Panic("Could not save document to database", err.Error())
	}
	return true
}

func GetAllTickets(email string) []Models.Ticket {
	results := []Models.Ticket{}
	cursor, _ := ticketsCollection.Find(context.TODO(), bson.M{"email": email})
	cursor.All(context.TODO(), &results)
	return results
}
