package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HostMark struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username     string             `json:"username" bson:"username,omitempty"`
	Grade        int32              `json:"grade" bson:"grade" binding:"required"`
	HostUsername string             `json:"hostUsername" bson:"hostUsername,omitempty"`
}
