package Models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	FirstName string        `json:"firstName" bson:"first_name" binding:"required"`
	LastName  string        `json:"lastName" bson:"last_name" binding:"required"`
	Email     string        `json:"email" bson:"email" binding:"required"`
	Password  string        `json:"password" bson:"password" binding:"required"`
}
