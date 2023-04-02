package Models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName string             `json:"firstName" bson:"first_name" binding:"required"`
	LastName  string             `json:"lastName" bson:"last_name" binding:"required"`
	Email     string             `json:"email" bson:"email" binding:"required"`
	Password  string             `json:"password" bson:"password" binding:"required"`
	Role      string             `json:"role" bson:"role" binding:"required"`
}
