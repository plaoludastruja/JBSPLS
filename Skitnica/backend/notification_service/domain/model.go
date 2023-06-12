package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Notification struct {
	Id      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Rceiver string             `json:"receiver" bson:"receiver" binding:"required"`
	Sender  string             `json:"sender" bson:"sender" binding:"sender"`
	Subject string             `json:"subject" bson:"subject" binding:"required"`
	Message string             `json:"message" bson:"message" binding:"required"`
	IsRead  bool               `json:"isRead" bson:"isRead" binding:"required"`
}
