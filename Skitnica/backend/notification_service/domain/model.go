package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Notification struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Receiver string             `json:"receiver" bson:"receiver" binding:"required"`
	Sender   string             `json:"sender" bson:"sender" binding:"required"`
	Subject  string             `json:"subject" bson:"subject" binding:"required"`
	Message  string             `json:"message" bson:"message" binding:"required"`
	IsRead   string             `json:"isRead" bson:"isRead" binding:"required"`
}

type NotificationFilter struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username    string             `json:"username" bson:"username" binding:"required"`
	Reservation bool               `json:"reservation" bson:"reservation" binding:"required"`
	Rating      bool               `json:"rating" bson:"rating" binding:"required"`
	Super       bool               `json:"super" bson:"super" binding:"required"`
}
