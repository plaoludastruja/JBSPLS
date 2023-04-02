package DTO

import "gopkg.in/mgo.v2/bson"

type DeleteDTO struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
}
